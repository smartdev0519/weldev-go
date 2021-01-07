// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package modget implements the module-aware ``go get'' command.
package modget

// The arguments to 'go get' are patterns with optional version queries, with
// the version queries defaulting to "upgrade".
//
// The patterns are normally interpreted as package patterns. However, if a
// pattern cannot match a package, it is instead interpreted as a *module*
// pattern. For version queries such as "upgrade" and "patch" that depend on the
// selected version of a module (or of the module containing a package),
// whether a pattern denotes a package or module may change as updates are
// applied (see the example in mod_get_patchmod.txt).
//
// There are a few other ambiguous cases to resolve, too. A package can exist in
// two different modules at the same version: for example, the package
// example.com/foo might be found in module example.com and also in module
// example.com/foo, and those modules may have independent v0.1.0 tags — so the
// input 'example.com/foo@v0.1.0' could syntactically refer to the variant of
// the package loaded from either module! (See mod_get_ambiguous_pkg.txt.)
// If the argument is ambiguous, the user can often disambiguate by specifying
// explicit versions for *all* of the potential module paths involved.

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"

	"cmd/go/internal/base"
	"cmd/go/internal/cfg"
	"cmd/go/internal/imports"
	"cmd/go/internal/load"
	"cmd/go/internal/modload"
	"cmd/go/internal/par"
	"cmd/go/internal/search"
	"cmd/go/internal/work"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
)

var CmdGet = &base.Command{
	// Note: -d -u are listed explicitly because they are the most common get flags.
	// Do not send CLs removing them because they're covered by [get flags].
	UsageLine: "go get [-d] [-t] [-u] [-v] [-insecure] [build flags] [packages]",
	Short:     "add dependencies to current module and install them",
	Long: `
Get resolves and adds dependencies to the current development module
and then builds and installs them.

The first step is to resolve which dependencies to add.

For each named package or package pattern, get must decide which version of
the corresponding module to use. By default, get looks up the latest tagged
release version, such as v0.4.5 or v1.2.3. If there are no tagged release
versions, get looks up the latest tagged pre-release version, such as
v0.0.1-pre1. If there are no tagged versions at all, get looks up the latest
known commit. If the module is not already required at a later version
(for example, a pre-release newer than the latest release), get will use
the version it looked up. Otherwise, get will use the currently
required version.

This default version selection can be overridden by adding an @version
suffix to the package argument, as in 'go get golang.org/x/text@v0.3.0'.
The version may be a prefix: @v1 denotes the latest available version starting
with v1. See 'go help modules' under the heading 'Module queries' for the
full query syntax.

For modules stored in source control repositories, the version suffix can
also be a commit hash, branch identifier, or other syntax known to the
source control system, as in 'go get golang.org/x/text@master'. Note that
branches with names that overlap with other module query syntax cannot be
selected explicitly. For example, the suffix @v2 means the latest version
starting with v2, not the branch named v2.

If a module under consideration is already a dependency of the current
development module, then get will update the required version.
Specifying a version earlier than the current required version is valid and
downgrades the dependency. The version suffix @none indicates that the
dependency should be removed entirely, downgrading or removing modules
depending on it as needed.

The version suffix @latest explicitly requests the latest minor release of
the module named by the given path. The suffix @upgrade is like @latest but
will not downgrade a module if it is already required at a revision or
pre-release version newer than the latest released version. The suffix
@patch requests the latest patch release: the latest released version
with the same major and minor version numbers as the currently required
version. Like @upgrade, @patch will not downgrade a module already required
at a newer version. If the path is not already required, @upgrade is
equivalent to @latest, and @patch is disallowed.

Although get defaults to using the latest version of the module containing
a named package, it does not use the latest version of that module's
dependencies. Instead it prefers to use the specific dependency versions
requested by that module. For example, if the latest A requires module
B v1.2.3, while B v1.2.4 and v1.3.1 are also available, then 'go get A'
will use the latest A but then use B v1.2.3, as requested by A. (If there
are competing requirements for a particular module, then 'go get' resolves
those requirements by taking the maximum requested version.)

The -t flag instructs get to consider modules needed to build tests of
packages specified on the command line.

The -u flag instructs get to update modules providing dependencies
of packages named on the command line to use newer minor or patch
releases when available. Continuing the previous example, 'go get -u A'
will use the latest A with B v1.3.1 (not B v1.2.3). If B requires module C,
but C does not provide any packages needed to build packages in A
(not including tests), then C will not be updated.

The -u=patch flag (not -u patch) also instructs get to update dependencies,
but changes the default to select patch releases.
Continuing the previous example,
'go get -u=patch A@latest' will use the latest A with B v1.2.4 (not B v1.2.3),
while 'go get -u=patch A' will use a patch release of A instead.

When the -t and -u flags are used together, get will update
test dependencies as well.

In general, adding a new dependency may require upgrading
existing dependencies to keep a working build, and 'go get' does
this automatically. Similarly, downgrading one dependency may
require downgrading other dependencies, and 'go get' does
this automatically as well.

The -insecure flag permits fetching from repositories and resolving
custom domains using insecure schemes such as HTTP, and also bypassess
module sum validation using the checksum database. Use with caution.
This flag is deprecated and will be removed in a future version of go.
To permit the use of insecure schemes, use the GOINSECURE environment
variable instead. To bypass module sum validation, use GOPRIVATE or
GONOSUMDB. See 'go help environment' for details.

The second step is to download (if needed), build, and install
the named packages.

The -d flag instructs get to skip this step, downloading source code
needed to build the named packages and their dependencies, but not
building or installing.

Building and installing packages with get is deprecated. In a future release,
the -d flag will be enabled by default, and 'go get' will be only be used to
adjust dependencies of the current module. To install a package using
dependencies from the current module, use 'go install'. To install a package
ignoring the current module, use 'go install' with an @version suffix like
"@latest" after each argument.

If an argument names a module but not a package (because there is no
Go source code in the module's root directory), then the install step
is skipped for that argument, instead of causing a build failure.
For example 'go get golang.org/x/perf' succeeds even though there
is no code corresponding to that import path.

Note that package patterns are allowed and are expanded after resolving
the module versions. For example, 'go get golang.org/x/perf/cmd/...'
adds the latest golang.org/x/perf and then installs the commands in that
latest version.

With no package arguments, 'go get' applies to Go package in the
current directory, if any. In particular, 'go get -u' and
'go get -u=patch' update all the dependencies of that package.
With no package arguments and also without -u, 'go get' is not much more
than 'go install', and 'go get -d' not much more than 'go list'.

For more about modules, see 'go help modules'.

For more about specifying packages, see 'go help packages'.

This text describes the behavior of get using modules to manage source
code and dependencies. If instead the go command is running in GOPATH
mode, the details of get's flags and effects change, as does 'go help get'.
See 'go help modules' and 'go help gopath-get'.

See also: go build, go install, go clean, go mod.
	`,
}

// Note that this help text is a stopgap to make the module-aware get help text
// available even in non-module settings. It should be deleted when the old get
// is deleted. It should NOT be considered to set a precedent of having hierarchical
// help names with dashes.
var HelpModuleGet = &base.Command{
	UsageLine: "module-get",
	Short:     "module-aware go get",
	Long: `
The 'go get' command changes behavior depending on whether the
go command is running in module-aware mode or legacy GOPATH mode.
This help text, accessible as 'go help module-get' even in legacy GOPATH mode,
describes 'go get' as it operates in module-aware mode.

Usage: ` + CmdGet.UsageLine + `
` + CmdGet.Long,
}

var HelpVCS = &base.Command{
	UsageLine: "vcs",
	Short:     "controlling version control with GOVCS",
	Long: `
The 'go get' command can run version control commands like git
to download imported code. This functionality is critical to the decentralized
Go package ecosystem, in which code can be imported from any server,
but it is also a potential security problem, if a malicious server finds a
way to cause the invoked version control command to run unintended code.

To balance the functionality and security concerns, the 'go get' command
by default will only use git and hg to download code from public servers.
But it will use any known version control system (bzr, fossil, git, hg, svn)
to download code from private servers, defined as those hosting packages
matching the GOPRIVATE variable (see 'go help private'). The rationale behind
allowing only Git and Mercurial is that these two systems have had the most
attention to issues of being run as clients of untrusted servers. In contrast,
Bazaar, Fossil, and Subversion have primarily been used in trusted,
authenticated environments and are not as well scrutinized as attack surfaces.

The version control command restrictions only apply when using direct version
control access to download code. When downloading modules from a proxy,
'go get' uses the proxy protocol instead, which is always permitted.
By default, the 'go get' command uses the Go module mirror (proxy.golang.org)
for public packages and only falls back to version control for private
packages or when the mirror refuses to serve a public package (typically for
legal reasons). Therefore, clients can still access public code served from
Bazaar, Fossil, or Subversion repositories by default, because those downloads
use the Go module mirror, which takes on the security risk of running the
version control commands, using a custom sandbox.

The GOVCS variable can be used to change the allowed version control systems
for specific packages (identified by a module or import path).
The GOVCS variable applies both when using modules and when using GOPATH.
When using modules, the patterns match against the module path.
When using GOPATH, the patterns match against the import path
corresponding to the root of the version control repository.

The general form of the GOVCS setting is a comma-separated list of
pattern:vcslist rules. The pattern is a glob pattern that must match
one or more leading elements of the module or import path. The vcslist
is a pipe-separated list of allowed version control commands, or "all"
to allow use of any known command, or "off" to allow nothing.
The earliest matching pattern in the list applies, even if later patterns
might also match.

For example, consider:

	GOVCS=github.com:git,evil.com:off,*:git|hg

With this setting, code with an module or import path beginning with
github.com/ can only use git; paths on evil.com cannot use any version
control command, and all other paths (* matches everything) can use
only git or hg.

The special patterns "public" and "private" match public and private
module or import paths. A path is private if it matches the GOPRIVATE
variable; otherwise it is public.

If no rules in the GOVCS variable match a particular module or import path,
the 'go get' command applies its default rule, which can now be summarized
in GOVCS notation as 'public:git|hg,private:all'.

To allow unfettered use of any version control system for any package, use:

	GOVCS=*:all

To disable all use of version control, use:

	GOVCS=*:off

The 'go env -w' command (see 'go help env') can be used to set the GOVCS
variable for future go command invocations.
`,
}

var (
	getD   = CmdGet.Flag.Bool("d", false, "")
	getF   = CmdGet.Flag.Bool("f", false, "")
	getFix = CmdGet.Flag.Bool("fix", false, "")
	getM   = CmdGet.Flag.Bool("m", false, "")
	getT   = CmdGet.Flag.Bool("t", false, "")
	getU   upgradeFlag
	// -insecure is cfg.Insecure
	// -v is cfg.BuildV
)

// upgradeFlag is a custom flag.Value for -u.
type upgradeFlag struct {
	rawVersion string
	version    string
}

func (*upgradeFlag) IsBoolFlag() bool { return true } // allow -u

func (v *upgradeFlag) Set(s string) error {
	if s == "false" {
		v.version = ""
		v.rawVersion = ""
	} else if s == "true" {
		v.version = "upgrade"
		v.rawVersion = ""
	} else {
		v.version = s
		v.rawVersion = s
	}
	return nil
}

func (v *upgradeFlag) String() string { return "" }

func init() {
	work.AddBuildFlags(CmdGet, work.OmitModFlag)
	CmdGet.Run = runGet // break init loop
	CmdGet.Flag.BoolVar(&cfg.Insecure, "insecure", cfg.Insecure, "")
	CmdGet.Flag.Var(&getU, "u", "")
}

func runGet(ctx context.Context, cmd *base.Command, args []string) {
	switch getU.version {
	case "", "upgrade", "patch":
		// ok
	default:
		base.Fatalf("go get: unknown upgrade flag -u=%s", getU.rawVersion)
	}
	if *getF {
		fmt.Fprintf(os.Stderr, "go get: -f flag is a no-op when using modules\n")
	}
	if *getFix {
		fmt.Fprintf(os.Stderr, "go get: -fix flag is a no-op when using modules\n")
	}
	if *getM {
		base.Fatalf("go get: -m flag is no longer supported; consider -d to skip building packages")
	}
	if cfg.Insecure {
		fmt.Fprintf(os.Stderr, "go get: -insecure flag is deprecated; see 'go help get' for details\n")
	}
	load.ModResolveTests = *getT

	// Do not allow any updating of go.mod until we've applied
	// all the requested changes and checked that the result matches
	// what was requested.
	modload.DisallowWriteGoMod()

	// Allow looking up modules for import paths when outside of a module.
	// 'go get' is expected to do this, unlike other commands.
	modload.AllowMissingModuleImports()

	modload.LoadModFile(ctx) // Initializes modload.Target.

	queries := parseArgs(ctx, args)

	r := newResolver(ctx, queries)
	r.performLocalQueries(ctx)
	r.performPathQueries(ctx)

	for {
		r.performWildcardQueries(ctx)
		r.performPatternAllQueries(ctx)

		if changed := r.resolveCandidates(ctx, queries, nil); changed {
			// 'go get' arguments can be (and often are) package patterns rather than
			// (just) modules. A package can be provided by any module with a prefix
			// of its import path, and a wildcard can even match packages in modules
			// with totally different paths. Because of these effects, and because any
			// change to the selected version of a module can bring in entirely new
			// module paths as dependencies, we need to reissue queries whenever we
			// change the build list.
			//
			// The result of any version query for a given module — even "upgrade" or
			// "patch" — is always relative to the build list at the start of
			// the 'go get' command, not an intermediate state, and is therefore
			// dederministic and therefore cachable, and the constraints on the
			// selected version of each module can only narrow as we iterate.
			//
			// "all" is functionally very similar to a wildcard pattern. The set of
			// packages imported by the main module does not change, and the query
			// result for the module containing each such package also does not change
			// (it is always relative to the initial build list, before applying
			// queries). So the only way that the result of an "all" query can change
			// is if some matching package moves from one module in the build list
			// to another, which should not happen very often.
			continue
		}

		// When we load imports, we detect the following conditions:
		//
		// - missing transitive depencies that need to be resolved from outside the
		//   current build list (note that these may add new matches for existing
		//   pattern queries!)
		//
		// - transitive dependencies that didn't match any other query,
		//   but need to be upgraded due to the -u flag
		//
		// - ambiguous import errors.
		//   TODO(#27899): Try to resolve ambiguous import errors automatically.
		upgrades := r.findAndUpgradeImports(ctx, queries)
		if changed := r.resolveCandidates(ctx, nil, upgrades); changed {
			continue
		}

		r.findMissingWildcards(ctx)
		if changed := r.resolveCandidates(ctx, r.wildcardQueries, nil); changed {
			continue
		}

		break
	}

	r.checkWildcardVersions(ctx)

	var pkgPatterns []string
	for _, q := range queries {
		if q.matchesPackages {
			pkgPatterns = append(pkgPatterns, q.pattern)
		}
	}
	r.checkPackagesAndRetractions(ctx, pkgPatterns)

	// We've already downloaded modules (and identified direct and indirect
	// dependencies) by loading packages in findAndUpgradeImports.
	// So if -d is set, we're done after the module work.
	//
	// Otherwise, we need to build and install the packages matched by
	// command line arguments.
	// Note that 'go get -u' without arguments is equivalent to
	// 'go get -u .', so we'll typically build the package in the current
	// directory.
	if !*getD && len(pkgPatterns) > 0 {
		work.BuildInit()
		pkgs := load.PackagesAndErrors(ctx, pkgPatterns)
		load.CheckPackageErrors(pkgs)
		work.InstallPackages(ctx, pkgPatterns, pkgs)
		// TODO(#40276): After Go 1.16, print a deprecation notice when building
		// and installing main packages. 'go install pkg' or
		// 'go install pkg@version' should be used instead.
		// Give the specific argument to use if possible.
	}

	if !modload.HasModRoot() {
		return
	}

	// Everything succeeded. Update go.mod.
	oldReqs := reqsFromGoMod(modload.ModFile())

	modload.AllowWriteGoMod()
	modload.WriteGoMod()
	modload.DisallowWriteGoMod()

	newReqs := reqsFromGoMod(modload.ModFile())
	r.reportChanges(oldReqs, newReqs)
}

// parseArgs parses command-line arguments and reports errors.
//
// The command-line arguments are of the form path@version or simply path, with
// implicit @upgrade. path@none is "downgrade away".
func parseArgs(ctx context.Context, rawArgs []string) []*query {
	defer base.ExitIfErrors()

	var queries []*query
	for _, arg := range search.CleanPatterns(rawArgs) {
		q, err := newQuery(arg)
		if err != nil {
			base.Errorf("go get: %v", err)
			continue
		}

		// If there were no arguments, CleanPatterns returns ".". Set the raw
		// string back to "" for better errors.
		if len(rawArgs) == 0 {
			q.raw = ""
		}

		// Guard against 'go get x.go', a common mistake.
		// Note that package and module paths may end with '.go', so only print an error
		// if the argument has no version and either has no slash or refers to an existing file.
		if strings.HasSuffix(q.raw, ".go") && q.rawVersion == "" {
			if !strings.Contains(q.raw, "/") {
				base.Errorf("go get %s: arguments must be package or module paths", q.raw)
				continue
			}
			if fi, err := os.Stat(q.raw); err == nil && !fi.IsDir() {
				base.Errorf("go get: %s exists as a file, but 'go get' requires package arguments", q.raw)
				continue
			}
		}

		queries = append(queries, q)
	}

	return queries
}

type resolver struct {
	localQueries      []*query // queries for absolute or relative paths
	pathQueries       []*query // package path literal queries in original order
	wildcardQueries   []*query // path wildcard queries in original order
	patternAllQueries []*query // queries with the pattern "all"

	// Indexed "none" queries. These are also included in the slices above;
	// they are indexed here to speed up noneForPath.
	nonesByPath   map[string]*query // path-literal "@none" queries indexed by path
	wildcardNones []*query          // wildcard "@none" queries

	// resolvedVersion maps each module path to the version of that module that
	// must be selected in the final build list, along with the first query
	// that resolved the module to that version (the “reason”).
	resolvedVersion map[string]versionReason

	buildList                 []module.Version
	buildListResolvedVersions int               // len(resolvedVersion) when buildList was computed
	buildListVersion          map[string]string // index of buildList (module path → version)

	initialVersion map[string]string // index of the initial build list at the start of 'go get'

	missing []pathSet // candidates for missing transitive dependencies

	work *par.Queue

	matchInModuleCache par.Cache
}

type versionReason struct {
	version string
	reason  *query
}

func newResolver(ctx context.Context, queries []*query) *resolver {
	buildList := modload.LoadAllModules(ctx)
	initialVersion := make(map[string]string, len(buildList))
	for _, m := range buildList {
		initialVersion[m.Path] = m.Version
	}

	r := &resolver{
		work:             par.NewQueue(runtime.GOMAXPROCS(0)),
		resolvedVersion:  map[string]versionReason{},
		buildList:        buildList,
		buildListVersion: initialVersion,
		initialVersion:   initialVersion,
		nonesByPath:      map[string]*query{},
	}

	for _, q := range queries {
		if q.pattern == "all" {
			r.patternAllQueries = append(r.patternAllQueries, q)
		} else if q.patternIsLocal {
			r.localQueries = append(r.localQueries, q)
		} else if q.isWildcard() {
			r.wildcardQueries = append(r.wildcardQueries, q)
		} else {
			r.pathQueries = append(r.pathQueries, q)
		}

		if q.version == "none" {
			// Index "none" queries to make noneForPath more efficient.
			if q.isWildcard() {
				r.wildcardNones = append(r.wildcardNones, q)
			} else {
				// All "<path>@none" queries for the same path are identical; we only
				// need to index one copy.
				r.nonesByPath[q.pattern] = q
			}
		}
	}

	return r
}

// initialSelected returns the version of the module with the given path that
// was selected at the start of this 'go get' invocation.
func (r *resolver) initialSelected(mPath string) (version string) {
	v, ok := r.initialVersion[mPath]
	if !ok {
		return "none"
	}
	return v
}

// selected returns the version of the module with the given path that is
// selected in the resolver's current build list.
func (r *resolver) selected(mPath string) (version string) {
	v, ok := r.buildListVersion[mPath]
	if !ok {
		return "none"
	}
	return v
}

// noneForPath returns a "none" query matching the given module path,
// or found == false if no such query exists.
func (r *resolver) noneForPath(mPath string) (nq *query, found bool) {
	if nq = r.nonesByPath[mPath]; nq != nil {
		return nq, true
	}
	for _, nq := range r.wildcardNones {
		if nq.matchesPath(mPath) {
			return nq, true
		}
	}
	return nil, false
}

// queryModule wraps modload.Query, substituting r.checkAllowedOr to decide
// allowed versions.
func (r *resolver) queryModule(ctx context.Context, mPath, query string, selected func(string) string) (module.Version, error) {
	current := r.initialSelected(mPath)
	rev, err := modload.Query(ctx, mPath, query, current, r.checkAllowedOr(query, selected))
	if err != nil {
		return module.Version{}, err
	}
	return module.Version{Path: mPath, Version: rev.Version}, nil
}

// queryPackage wraps modload.QueryPackage, substituting r.checkAllowedOr to
// decide allowed versions.
func (r *resolver) queryPackages(ctx context.Context, pattern, query string, selected func(string) string) (pkgMods []module.Version, err error) {
	results, err := modload.QueryPackages(ctx, pattern, query, selected, r.checkAllowedOr(query, selected))
	if len(results) > 0 {
		pkgMods = make([]module.Version, 0, len(results))
		for _, qr := range results {
			pkgMods = append(pkgMods, qr.Mod)
		}
	}
	return pkgMods, err
}

// queryPattern wraps modload.QueryPattern, substituting r.checkAllowedOr to
// decide allowed versions.
func (r *resolver) queryPattern(ctx context.Context, pattern, query string, selected func(string) string) (pkgMods []module.Version, mod module.Version, err error) {
	results, modOnly, err := modload.QueryPattern(ctx, pattern, query, selected, r.checkAllowedOr(query, selected))
	if len(results) > 0 {
		pkgMods = make([]module.Version, 0, len(results))
		for _, qr := range results {
			pkgMods = append(pkgMods, qr.Mod)
		}
	}
	if modOnly != nil {
		mod = modOnly.Mod
	}
	return pkgMods, mod, err
}

// checkAllowedOr is like modload.CheckAllowed, but it always allows the requested
// and current versions (even if they are retracted or otherwise excluded).
func (r *resolver) checkAllowedOr(requested string, selected func(string) string) modload.AllowedFunc {
	return func(ctx context.Context, m module.Version) error {
		if m.Version == requested {
			return modload.CheckExclusions(ctx, m)
		}
		if (requested == "upgrade" || requested == "patch") && m.Version == selected(m.Path) {
			return nil
		}
		return modload.CheckAllowed(ctx, m)
	}
}

// matchInModule is a caching wrapper around modload.MatchInModule.
func (r *resolver) matchInModule(ctx context.Context, pattern string, m module.Version) (packages []string, err error) {
	type key struct {
		pattern string
		m       module.Version
	}
	type entry struct {
		packages []string
		err      error
	}

	e := r.matchInModuleCache.Do(key{pattern, m}, func() interface{} {
		match := modload.MatchInModule(ctx, pattern, m, imports.AnyTags())
		if len(match.Errs) > 0 {
			return entry{match.Pkgs, match.Errs[0]}
		}
		return entry{match.Pkgs, nil}
	}).(entry)

	return e.packages, e.err
}

// queryNone adds a candidate set to q for each module matching q.pattern.
// Each candidate set has only one possible module version: the matched
// module at version "none".
//
// We interpret arguments to 'go get' as packages first, and fall back to
// modules second. However, no module exists at version "none", and therefore no
// package exists at that version either: we know that the argument cannot match
// any packages, and thus it must match modules instead.
func (r *resolver) queryNone(ctx context.Context, q *query) {
	if search.IsMetaPackage(q.pattern) {
		panic(fmt.Sprintf("internal error: queryNone called with pattern %q", q.pattern))
	}

	if !q.isWildcard() {
		q.pathOnce(q.pattern, func() pathSet {
			if modload.HasModRoot() && q.pattern == modload.Target.Path {
				// The user has explicitly requested to downgrade their own module to
				// version "none". This is not an entirely unreasonable request: it
				// could plausibly mean “downgrade away everything that depends on any
				// explicit version of the main module”, or “downgrade away the
				// package with the same path as the main module, found in a module
				// with a prefix of the main module's path”.
				//
				// However, neither of those behaviors would be consistent with the
				// plain meaning of the query. To try to reduce confusion, reject the
				// query explicitly.
				return errSet(&modload.QueryMatchesMainModuleError{Pattern: q.pattern, Query: q.version})
			}

			return pathSet{mod: module.Version{Path: q.pattern, Version: "none"}}
		})
	}

	for _, curM := range r.buildList {
		if !q.matchesPath(curM.Path) {
			continue
		}
		q.pathOnce(curM.Path, func() pathSet {
			if modload.HasModRoot() && curM == modload.Target {
				return errSet(&modload.QueryMatchesMainModuleError{Pattern: q.pattern, Query: q.version})
			}
			return pathSet{mod: module.Version{Path: curM.Path, Version: "none"}}
		})
	}
}

func (r *resolver) performLocalQueries(ctx context.Context) {
	for _, q := range r.localQueries {
		q.pathOnce(q.pattern, func() pathSet {
			absDetail := ""
			if !filepath.IsAbs(q.pattern) {
				if absPath, err := filepath.Abs(q.pattern); err == nil {
					absDetail = fmt.Sprintf(" (%s)", absPath)
				}
			}

			// Absolute paths like C:\foo and relative paths like ../foo... are
			// restricted to matching packages in the main module.
			pkgPattern := modload.DirImportPath(q.pattern)
			if pkgPattern == "." {
				return errSet(fmt.Errorf("%s%s is not within module rooted at %s", q.pattern, absDetail, modload.ModRoot()))
			}

			match := modload.MatchInModule(ctx, pkgPattern, modload.Target, imports.AnyTags())
			if len(match.Errs) > 0 {
				return pathSet{err: match.Errs[0]}
			}

			if len(match.Pkgs) == 0 {
				if q.raw == "" || q.raw == "." {
					return errSet(fmt.Errorf("no package in current directory"))
				}
				if !q.isWildcard() {
					return errSet(fmt.Errorf("%s%s is not a package in module rooted at %s", q.pattern, absDetail, modload.ModRoot()))
				}
				search.WarnUnmatched([]*search.Match{match})
				return pathSet{}
			}

			return pathSet{pkgMods: []module.Version{modload.Target}}
		})
	}
}

// performWildcardQueries populates the candidates for each query whose pattern
// is a wildcard.
//
// The candidates for a given module path matching (or containing a package
// matching) a wildcard query depend only on the initial build list, but the set
// of modules may be expanded by other queries, so wildcard queries need to be
// re-evaluated whenever a potentially-matching module path is added to the
// build list.
func (r *resolver) performWildcardQueries(ctx context.Context) {
	for _, q := range r.wildcardQueries {
		q := q
		r.work.Add(func() {
			if q.version == "none" {
				r.queryNone(ctx, q)
			} else {
				r.queryWildcard(ctx, q)
			}
		})
	}
	<-r.work.Idle()
}

// queryWildcard adds a candidate set to q for each module for which:
// 	- some version of the module is already in the build list, and
// 	- that module exists at some version matching q.version, and
// 	- either the module path itself matches q.pattern, or some package within
// 	  the module at q.version matches q.pattern.
func (r *resolver) queryWildcard(ctx context.Context, q *query) {
	// For wildcard patterns, modload.QueryPattern only identifies modules
	// matching the prefix of the path before the wildcard. However, the build
	// list may already contain other modules with matching packages, and we
	// should consider those modules to satisfy the query too.
	// We want to match any packages in existing dependencies, but we only want to
	// resolve new dependencies if nothing else turns up.
	for _, curM := range r.buildList {
		if !q.canMatchInModule(curM.Path) {
			continue
		}
		q.pathOnce(curM.Path, func() pathSet {
			if _, hit := r.noneForPath(curM.Path); hit {
				// This module is being removed, so it will no longer be in the build list
				// (and thus will no longer match the pattern).
				return pathSet{}
			}

			if curM.Path == modload.Target.Path && !versionOkForMainModule(q.version) {
				if q.matchesPath(curM.Path) {
					return errSet(&modload.QueryMatchesMainModuleError{
						Pattern: q.pattern,
						Query:   q.version,
					})
				}

				packages, err := r.matchInModule(ctx, q.pattern, curM)
				if err != nil {
					return errSet(err)
				}
				if len(packages) > 0 {
					return errSet(&modload.QueryMatchesPackagesInMainModuleError{
						Pattern:  q.pattern,
						Query:    q.version,
						Packages: packages,
					})
				}

				return r.tryWildcard(ctx, q, curM)
			}

			m, err := r.queryModule(ctx, curM.Path, q.version, r.initialSelected)
			if err != nil {
				if !isNoSuchModuleVersion(err) {
					// We can't tell whether a matching version exists.
					return errSet(err)
				}
				// There is no version of curM.Path matching the query.

				// We haven't checked whether curM contains any matching packages at its
				// currently-selected version, or whether curM.Path itself matches q. If
				// either of those conditions holds, *and* no other query changes the
				// selected version of curM, then we will fail in checkWildcardVersions.
				// (This could be an error, but it's too soon to tell.)
				//
				// However, even then the transitive requirements of some other query
				// may downgrade this module out of the build list entirely, in which
				// case the pattern will no longer include it and it won't be an error.
				//
				// Either way, punt on the query rather than erroring out just yet.
				return pathSet{}
			}

			return r.tryWildcard(ctx, q, m)
		})
	}

	// Even if no modules matched, we shouldn't query for a new module to provide
	// the pattern yet: some other query may yet induce a new requirement that
	// will match the wildcard. Instead, we'll check in findMissingWildcards.
}

// tryWildcard returns a pathSet for module m matching query q.
// If m does not actually match q, tryWildcard returns an empty pathSet.
func (r *resolver) tryWildcard(ctx context.Context, q *query, m module.Version) pathSet {
	mMatches := q.matchesPath(m.Path)
	packages, err := r.matchInModule(ctx, q.pattern, m)
	if err != nil {
		return errSet(err)
	}
	if len(packages) > 0 {
		return pathSet{pkgMods: []module.Version{m}}
	}
	if mMatches {
		return pathSet{mod: m}
	}
	return pathSet{}
}

// findMissingWildcards adds a candidate set for each query in r.wildcardQueries
// that has not yet resolved to any version containing packages.
func (r *resolver) findMissingWildcards(ctx context.Context) {
	for _, q := range r.wildcardQueries {
		if q.version == "none" || q.matchesPackages {
			continue // q is not “missing”
		}
		r.work.Add(func() {
			q.pathOnce(q.pattern, func() pathSet {
				pkgMods, mod, err := r.queryPattern(ctx, q.pattern, q.version, r.initialSelected)
				if err != nil {
					if isNoSuchPackageVersion(err) && len(q.resolved) > 0 {
						// q already resolved one or more modules but matches no packages.
						// That's ok: this pattern is just a module pattern, and we don't
						// need to add any more modules to satisfy it.
						return pathSet{}
					}
					return errSet(err)
				}

				return pathSet{pkgMods: pkgMods, mod: mod}
			})
		})
	}
	<-r.work.Idle()
}

// checkWildcardVersions reports an error if any module in the build list has a
// path (or contains a package) matching a query with a wildcard pattern, but
// has a selected version that does *not* match the query.
func (r *resolver) checkWildcardVersions(ctx context.Context) {
	defer base.ExitIfErrors()

	for _, q := range r.wildcardQueries {
		for _, curM := range r.buildList {
			if !q.canMatchInModule(curM.Path) {
				continue
			}
			if !q.matchesPath(curM.Path) {
				packages, err := r.matchInModule(ctx, q.pattern, curM)
				if len(packages) == 0 {
					if err != nil {
						reportError(q, err)
					}
					continue // curM is not relevant to q.
				}
			}

			rev, err := r.queryModule(ctx, curM.Path, q.version, r.initialSelected)
			if err != nil {
				reportError(q, err)
				continue
			}
			if rev.Version == curM.Version {
				continue // curM already matches q.
			}

			if !q.matchesPath(curM.Path) {
				m := module.Version{Path: curM.Path, Version: rev.Version}
				packages, err := r.matchInModule(ctx, q.pattern, m)
				if err != nil {
					reportError(q, err)
					continue
				}
				if len(packages) == 0 {
					// curM at its original version contains a path matching q.pattern,
					// but at rev.Version it does not, so (somewhat paradoxically) if
					// we changed the version of curM it would no longer match the query.
					var version interface{} = m
					if rev.Version != q.version {
						version = fmt.Sprintf("%s@%s (%s)", m.Path, q.version, m.Version)
					}
					reportError(q, fmt.Errorf("%v matches packages in %v but not %v: specify a different version for module %s", q, curM, version, m.Path))
					continue
				}
			}

			// Since queryModule succeeded and either curM or one of the packages it
			// contains matches q.pattern, we should have either selected the version
			// of curM matching q, or reported a conflict error (and exited).
			// If we're still here and the version doesn't match,
			// something has gone very wrong.
			reportError(q, fmt.Errorf("internal error: selected %v instead of %v", curM, rev.Version))
		}
	}
}

// performPathQueries populates the candidates for each query whose pattern is
// a path literal.
//
// The candidate packages and modules for path literals depend only on the
// initial build list, not the current build list, so we only need to query path
// literals once.
func (r *resolver) performPathQueries(ctx context.Context) {
	for _, q := range r.pathQueries {
		q := q
		r.work.Add(func() {
			if q.version == "none" {
				r.queryNone(ctx, q)
			} else {
				r.queryPath(ctx, q)
			}
		})
	}
	<-r.work.Idle()
}

// queryPath adds a candidate set to q for the package with path q.pattern.
// The candidate set consists of all modules that could provide q.pattern
// and have a version matching q, plus (if it exists) the module whose path
// is itself q.pattern (at a matching version).
func (r *resolver) queryPath(ctx context.Context, q *query) {
	q.pathOnce(q.pattern, func() pathSet {
		if search.IsMetaPackage(q.pattern) || q.isWildcard() {
			panic(fmt.Sprintf("internal error: queryPath called with pattern %q", q.pattern))
		}
		if q.version == "none" {
			panic(`internal error: queryPath called with version "none"`)
		}

		if search.IsStandardImportPath(q.pattern) {
			stdOnly := module.Version{}
			packages, _ := r.matchInModule(ctx, q.pattern, stdOnly)
			if len(packages) > 0 {
				if q.rawVersion != "" {
					return errSet(fmt.Errorf("can't request explicit version %q of standard library package %s", q.version, q.pattern))
				}

				q.matchesPackages = true
				return pathSet{} // No module needed for standard library.
			}
		}

		pkgMods, mod, err := r.queryPattern(ctx, q.pattern, q.version, r.initialSelected)
		if err != nil {
			return errSet(err)
		}
		return pathSet{pkgMods: pkgMods, mod: mod}
	})
}

// performPatternAllQueries populates the candidates for each query whose
// pattern is "all".
//
// The candidate modules for a given package in "all" depend only on the initial
// build list, but we cannot follow the dependencies of a given package until we
// know which candidate is selected — and that selection may depend on the
// results of other queries. We need to re-evaluate the "all" queries whenever
// the module for one or more packages in "all" are resolved.
func (r *resolver) performPatternAllQueries(ctx context.Context) {
	if len(r.patternAllQueries) == 0 {
		return
	}

	findPackage := func(ctx context.Context, path string, m module.Version) (versionOk bool) {
		versionOk = true
		for _, q := range r.patternAllQueries {
			q.pathOnce(path, func() pathSet {
				pkgMods, err := r.queryPackages(ctx, path, q.version, r.initialSelected)
				if len(pkgMods) != 1 || pkgMods[0] != m {
					// There are candidates other than m for the given path, so we can't
					// be certain that m will actually be the module selected to provide
					// the package. Don't load its dependencies just yet, because they
					// might no longer be dependencies after we resolve the correct
					// version.
					versionOk = false
				}
				return pathSet{pkgMods: pkgMods, err: err}
			})
		}
		return versionOk
	}

	r.loadPackages(ctx, []string{"all"}, findPackage)

	// Since we built up the candidate lists concurrently, they may be in a
	// nondeterministic order. We want 'go get' to be fully deterministic,
	// including in which errors it chooses to report, so sort the candidates
	// into a deterministic-but-arbitrary order.
	for _, q := range r.patternAllQueries {
		sort.Slice(q.candidates, func(i, j int) bool {
			return q.candidates[i].path < q.candidates[j].path
		})
	}
}

// findAndUpgradeImports returns a pathSet for each package that is not yet
// in the build list but is transitively imported by the packages matching the
// given queries (which must already have been resolved).
//
// If the getU flag ("-u") is set, findAndUpgradeImports also returns a
// pathSet for each module that is not constrained by any other
// command-line argument and has an available matching upgrade.
func (r *resolver) findAndUpgradeImports(ctx context.Context, queries []*query) (upgrades []pathSet) {
	patterns := make([]string, 0, len(queries))
	for _, q := range queries {
		if q.matchesPackages {
			patterns = append(patterns, q.pattern)
		}
	}
	if len(patterns) == 0 {
		return nil
	}

	// mu guards concurrent writes to upgrades, which will be sorted
	// (to restore determinism) after loading.
	var mu sync.Mutex

	findPackage := func(ctx context.Context, path string, m module.Version) (versionOk bool) {
		version := "latest"
		if m.Path != "" {
			if getU.version == "" {
				// The user did not request that we upgrade transitive dependencies.
				return true
			}
			if _, ok := r.resolvedVersion[m.Path]; ok {
				// We cannot upgrade m implicitly because its version is determined by
				// an explicit pattern argument.
				return true
			}
			version = getU.version
		}

		// Unlike other queries, the "-u" flag upgrades relative to the build list
		// after applying changes so far, not the initial build list.
		// This is for two reasons:
		//
		// 	- The "-u" flag intentionally applies to transitive dependencies,
		// 	  which may not be known or even resolved in advance of applying
		// 	  other version changes.
		//
		// 	- The "-u" flag, unlike other arguments, does not cause version
		// 	  conflicts with other queries. (The other query always wins.)

		pkgMods, err := r.queryPackages(ctx, path, version, r.selected)
		for _, u := range pkgMods {
			if u == m {
				// The selected package version is already upgraded appropriately; there
				// is no need to change it.
				return true
			}
		}

		if err != nil {
			if isNoSuchPackageVersion(err) || (m.Path == "" && module.CheckPath(path) != nil) {
				// We can't find the package because it doesn't — or can't — even exist
				// in any module at the latest version. (Note that invalid module paths
				// could in general exist due to replacements, so we at least need to
				// run the query to check those.)
				//
				// There is no version change we can make to fix the package, so leave
				// it unresolved. Either some other query (perhaps a wildcard matching a
				// newly-added dependency for some other missing package) will fill in
				// the gaps, or we will report an error (with a better import stack) in
				// the final LoadPackages call.
				return true
			}
		}

		mu.Lock()
		upgrades = append(upgrades, pathSet{path: path, pkgMods: pkgMods, err: err})
		mu.Unlock()
		return false
	}

	r.loadPackages(ctx, patterns, findPackage)

	// Since we built up the candidate lists concurrently, they may be in a
	// nondeterministic order. We want 'go get' to be fully deterministic,
	// including in which errors it chooses to report, so sort the candidates
	// into a deterministic-but-arbitrary order.
	sort.Slice(upgrades, func(i, j int) bool {
		return upgrades[i].path < upgrades[j].path
	})
	return upgrades
}

// loadPackages loads the packages matching the given patterns, invoking the
// findPackage function for each package that may require a change to the
// build list.
//
// loadPackages invokes the findPackage function for each package loaded from a
// module outside the main module. If the module or version that supplies that
// package needs to be changed due to a query, findPackage may return false
// and the imports of that package will not be loaded.
//
// loadPackages also invokes the findPackage function for each imported package
// that is neither present in the standard library nor in any module in the
// build list.
func (r *resolver) loadPackages(ctx context.Context, patterns []string, findPackage func(ctx context.Context, path string, m module.Version) (versionOk bool)) {
	opts := modload.PackageOpts{
		Tags:          imports.AnyTags(),
		LoadTests:     *getT,
		SilenceErrors: true, // May be fixed by subsequent upgrades or downgrades.
	}

	opts.AllowPackage = func(ctx context.Context, path string, m module.Version) error {
		if m.Path == "" || m == modload.Target {
			// Packages in the standard library and main module are already at their
			// latest (and only) available versions.
			return nil
		}
		if ok := findPackage(ctx, path, m); !ok {
			return errVersionChange
		}
		return nil
	}

	_, pkgs := modload.LoadPackages(ctx, opts, patterns...)
	for _, path := range pkgs {
		const (
			parentPath  = ""
			parentIsStd = false
		)
		_, _, err := modload.Lookup(parentPath, parentIsStd, path)
		if err == nil {
			continue
		}
		if errors.Is(err, errVersionChange) {
			// We already added candidates during loading.
			continue
		}

		var (
			importMissing *modload.ImportMissingError
			ambiguous     *modload.AmbiguousImportError
		)
		if !errors.As(err, &importMissing) && !errors.As(err, &ambiguous) {
			// The package, which is a dependency of something we care about, has some
			// problem that we can't resolve with a version change.
			// Leave the error for the final LoadPackages call.
			continue
		}

		path := path
		r.work.Add(func() {
			findPackage(ctx, path, module.Version{})
		})
	}
	<-r.work.Idle()
}

// errVersionChange is a sentinel error indicating that a module's version needs
// to be updated before its dependencies can be loaded.
var errVersionChange = errors.New("version change needed")

// resolveCandidates resolves candidates sets that are attached to the given
// queries and/or needed to provide the given missing-package dependencies.
//
// resolveCandidates starts by resolving one module version from each
// unambiguous pathSet attached to the given queries.
//
// If no unambiguous query results in a change to the build list,
// resolveCandidates modifies the build list by adding one module version from
// each pathSet in missing, but does not mark those versions as resolved
// (so they can still be modified by other queries).
//
// If that still does not result in any changes to the build list,
// resolveCandidates revisits the ambiguous query candidates and resolves them
// arbitrarily in order to guarantee forward progress.
//
// If all pathSets are resolved without any changes to the build list,
// resolveCandidates returns with changed=false.
func (r *resolver) resolveCandidates(ctx context.Context, queries []*query, upgrades []pathSet) (changed bool) {
	defer base.ExitIfErrors()

	// Note: this is O(N²) with the number of pathSets in the worst case.
	//
	// We could perhaps get it down to O(N) if we were to index the pathSets
	// by module path, so that we only revisit a given pathSet when the
	// version of some module in its containingPackage list has been determined.
	//
	// However, N tends to be small, and most candidate sets will include only one
	// candidate module (so they will be resolved in the first iteration), so for
	// now we'll stick to the simple O(N²) approach.

	resolved := 0
	for {
		prevResolved := resolved

		for _, q := range queries {
			unresolved := q.candidates[:0]

			for _, cs := range q.candidates {
				if cs.err != nil {
					reportError(q, cs.err)
					resolved++
					continue
				}

				filtered, isPackage, m, unique := r.disambiguate(cs)
				if !unique {
					unresolved = append(unresolved, filtered)
					continue
				}

				if m.Path == "" {
					// The query is not viable. Choose an arbitrary candidate from
					// before filtering and “resolve” it to report a conflict.
					isPackage, m = r.chooseArbitrarily(cs)
				}
				if isPackage {
					q.matchesPackages = true
				}
				r.resolve(q, m)
				resolved++
			}

			q.candidates = unresolved
		}

		base.ExitIfErrors()
		if resolved == prevResolved {
			break // No unambiguous candidate remains.
		}
	}

	if changed := r.updateBuildList(ctx, nil); changed {
		// The build list has changed, so disregard any missing packages: they might
		// now be determined by requirements in the build list, which we would
		// prefer to use instead of arbitrary "latest" versions.
		return true
	}

	// Arbitrarily add a "latest" version that provides each missing package, but
	// do not mark the version as resolved: we still want to allow the explicit
	// queries to modify the resulting versions.
	var tentative []module.Version
	for _, cs := range upgrades {
		if cs.err != nil {
			base.Errorf("go get: %v", cs.err)
			continue
		}

		filtered, _, m, unique := r.disambiguate(cs)
		if !unique {
			_, m = r.chooseArbitrarily(filtered)
		}
		if m.Path == "" {
			// There is no viable candidate for the missing package.
			// Leave it unresolved.
			continue
		}
		tentative = append(tentative, m)
	}
	base.ExitIfErrors()
	if changed := r.updateBuildList(ctx, tentative); changed {
		return true
	}

	// The build list will be the same on the next iteration as it was on this
	// iteration, so any ambiguous queries will remain so. In order to make
	// progress, resolve them arbitrarily but deterministically.
	//
	// If that results in conflicting versions, the user can re-run 'go get'
	// with additional explicit versions for the conflicting packages or
	// modules.
	for _, q := range queries {
		for _, cs := range q.candidates {
			isPackage, m := r.chooseArbitrarily(cs)
			if isPackage {
				q.matchesPackages = true
			}
			r.resolve(q, m)
		}
	}
	return r.updateBuildList(ctx, nil)
}

// disambiguate eliminates candidates from cs that conflict with other module
// versions that have already been resolved. If there is only one (unique)
// remaining candidate, disambiguate returns that candidate, along with
// an indication of whether that result interprets cs.path as a package
//
// Note: we're only doing very simple disambiguation here. The goal is to
// reproduce the user's intent, not to find a solution that a human couldn't.
// In the vast majority of cases, we expect only one module per pathSet,
// but we want to give some minimal additional tools so that users can add an
// extra argument or two on the command line to resolve simple ambiguities.
func (r *resolver) disambiguate(cs pathSet) (filtered pathSet, isPackage bool, m module.Version, unique bool) {
	if len(cs.pkgMods) == 0 && cs.mod.Path == "" {
		panic("internal error: resolveIfUnambiguous called with empty pathSet")
	}

	for _, m := range cs.pkgMods {
		if _, ok := r.noneForPath(m.Path); ok {
			// A query with version "none" forces the candidate module to version
			// "none", so we cannot use any other version for that module.
			continue
		}

		if m.Path == modload.Target.Path {
			if m.Version == modload.Target.Version {
				return pathSet{}, true, m, true
			}
			// The main module can only be set to its own version.
			continue
		}

		vr, ok := r.resolvedVersion[m.Path]
		if !ok {
			// m is a viable answer to the query, but other answers may also
			// still be viable.
			filtered.pkgMods = append(filtered.pkgMods, m)
			continue
		}

		if vr.version != m.Version {
			// Some query forces the candidate module to a version other than this
			// one.
			//
			// The command could be something like
			//
			// 	go get example.com/foo/bar@none example.com/foo/bar/baz@latest
			//
			// in which case we *cannot* resolve the package from
			// example.com/foo/bar (because it is constrained to version
			// "none") and must fall through to module example.com/foo@latest.
			continue
		}

		// Some query forces the candidate module *to* the candidate version.
		// As a result, this candidate is the only viable choice to provide
		// its package(s): any other choice would result in an ambiguous import
		// for this path.
		//
		// For example, consider the command
		//
		// 	go get example.com/foo@latest example.com/foo/bar/baz@latest
		//
		// If modules example.com/foo and example.com/foo/bar both provide
		// package example.com/foo/bar/baz, then we *must* resolve the package
		// from example.com/foo: if we instead resolved it from
		// example.com/foo/bar, we would have two copies of the package.
		return pathSet{}, true, m, true
	}

	if cs.mod.Path != "" {
		vr, ok := r.resolvedVersion[cs.mod.Path]
		if !ok || vr.version == cs.mod.Version {
			filtered.mod = cs.mod
		}
	}

	if len(filtered.pkgMods) == 1 &&
		(filtered.mod.Path == "" || filtered.mod == filtered.pkgMods[0]) {
		// Exactly one viable module contains the package with the given path
		// (by far the common case), so we can resolve it unambiguously.
		return pathSet{}, true, filtered.pkgMods[0], true
	}

	if len(filtered.pkgMods) == 0 {
		// All modules that could provide the path as a package conflict with other
		// resolved arguments. If it can refer to a module instead, return that;
		// otherwise, this pathSet cannot be resolved (and we will return the
		// zero module.Version).
		return pathSet{}, false, filtered.mod, true
	}

	// The query remains ambiguous: there are at least two different modules
	// to which cs.path could refer.
	return filtered, false, module.Version{}, false
}

// chooseArbitrarily returns an arbitrary (but deterministic) module version
// from among those in the given set.
//
// chooseArbitrarily prefers module paths that were already in the build list at
// the start of 'go get', prefers modules that provide packages over those that
// do not, and chooses the first module meeting those criteria (so biases toward
// longer paths).
func (r *resolver) chooseArbitrarily(cs pathSet) (isPackage bool, m module.Version) {
	// Prefer to upgrade some module that was already in the build list.
	for _, m := range cs.pkgMods {
		if r.initialSelected(m.Path) != "none" {
			return true, m
		}
	}

	// Otherwise, arbitrarily choose the first module that provides the package.
	if len(cs.pkgMods) > 0 {
		return true, cs.pkgMods[0]
	}

	return false, cs.mod
}

// checkPackagesAndRetractions reloads packages for the given patterns and
// reports missing and ambiguous package errors. It also reports loads and
// reports retractions for resolved modules and modules needed to build
// named packages.
//
// We skip missing-package errors earlier in the process, since we want to
// resolve pathSets ourselves, but at that point, we don't have enough context
// to log the package-import chains leading to each error.
func (r *resolver) checkPackagesAndRetractions(ctx context.Context, pkgPatterns []string) {
	defer base.ExitIfErrors()

	// Build a list of modules to load retractions for. Start with versions
	// selected based on command line queries.
	//
	// This is a subset of the build list. If the main module has a lot of
	// dependencies, loading retractions for the entire build list would be slow.
	relevantMods := make(map[module.Version]struct{})
	for path, reason := range r.resolvedVersion {
		relevantMods[module.Version{Path: path, Version: reason.version}] = struct{}{}
	}

	// Reload packages, reporting errors for missing and ambiguous imports.
	if len(pkgPatterns) > 0 {
		// LoadPackages will print errors (since it has more context) but will not
		// exit, since we need to load retractions later.
		pkgOpts := modload.PackageOpts{
			LoadTests:             *getT,
			ResolveMissingImports: false,
			AllowErrors:           true,
		}
		matches, pkgs := modload.LoadPackages(ctx, pkgOpts, pkgPatterns...)
		for _, m := range matches {
			if len(m.Errs) > 0 {
				base.SetExitStatus(1)
				break
			}
		}
		for _, pkg := range pkgs {
			if _, _, err := modload.Lookup("", false, pkg); err != nil {
				base.SetExitStatus(1)
				if ambiguousErr := (*modload.AmbiguousImportError)(nil); errors.As(err, &ambiguousErr) {
					for _, m := range ambiguousErr.Modules {
						relevantMods[m] = struct{}{}
					}
				}
			}
			if m := modload.PackageModule(pkg); m.Path != "" {
				relevantMods[m] = struct{}{}
			}
		}
	}

	// Load and report retractions.
	type retraction struct {
		m   module.Version
		err error
	}
	retractions := make([]retraction, 0, len(relevantMods))
	for m := range relevantMods {
		retractions = append(retractions, retraction{m: m})
	}
	sort.Slice(retractions, func(i, j int) bool {
		return retractions[i].m.Path < retractions[j].m.Path
	})
	for i := 0; i < len(retractions); i++ {
		i := i
		r.work.Add(func() {
			err := modload.CheckRetractions(ctx, retractions[i].m)
			if retractErr := (*modload.ModuleRetractedError)(nil); errors.As(err, &retractErr) {
				retractions[i].err = err
			}
		})
	}
	<-r.work.Idle()
	var retractPath string
	for _, r := range retractions {
		if r.err != nil {
			fmt.Fprintf(os.Stderr, "go: warning: %v\n", r.err)
			if retractPath == "" {
				retractPath = r.m.Path
			} else {
				retractPath = "<module>"
			}
		}
	}
	if retractPath != "" {
		fmt.Fprintf(os.Stderr, "go: to switch to the latest unretracted version, run:\n\tgo get %s@latest", retractPath)
	}
}

// reportChanges logs version changes to os.Stderr.
//
// reportChanges only logs changes to modules named on the command line and to
// explicitly required modules in go.mod. Most changes to indirect requirements
// are not relevant to the user and are not logged.
//
// reportChanges should be called after WriteGoMod.
func (r *resolver) reportChanges(oldReqs, newReqs []module.Version) {
	type change struct {
		path, old, new string
	}
	changes := make(map[string]change)

	// Collect changes in modules matched by command line arguments.
	for path, reason := range r.resolvedVersion {
		old := r.initialVersion[path]
		new := reason.version
		if old != new && (old != "" || new != "none") {
			changes[path] = change{path, old, new}
		}
	}

	// Collect changes to explicit requirements in go.mod.
	for _, req := range oldReqs {
		path := req.Path
		old := req.Version
		new := r.buildListVersion[path]
		if old != new {
			changes[path] = change{path, old, new}
		}
	}
	for _, req := range newReqs {
		path := req.Path
		old := r.initialVersion[path]
		new := req.Version
		if old != new {
			changes[path] = change{path, old, new}
		}
	}

	sortedChanges := make([]change, 0, len(changes))
	for _, c := range changes {
		sortedChanges = append(sortedChanges, c)
	}
	sort.Slice(sortedChanges, func(i, j int) bool {
		return sortedChanges[i].path < sortedChanges[j].path
	})
	for _, c := range sortedChanges {
		if c.old == "" {
			fmt.Fprintf(os.Stderr, "go get: added %s %s\n", c.path, c.new)
		} else if c.new == "none" || c.new == "" {
			fmt.Fprintf(os.Stderr, "go get: removed %s %s\n", c.path, c.old)
		} else if semver.Compare(c.new, c.old) > 0 {
			fmt.Fprintf(os.Stderr, "go get: upgraded %s %s => %s\n", c.path, c.old, c.new)
		} else {
			fmt.Fprintf(os.Stderr, "go get: downgraded %s %s => %s\n", c.path, c.old, c.new)
		}
	}

	// TODO(golang.org/issue/33284): attribute changes to command line arguments.
	// For modules matched by command line arguments, this probably isn't
	// necessary, but it would be useful for unmatched direct dependencies of
	// the main module.
}

// resolve records that module m must be at its indicated version (which may be
// "none") due to query q. If some other query forces module m to be at a
// different version, resolve reports a conflict error.
func (r *resolver) resolve(q *query, m module.Version) {
	if m.Path == "" {
		panic("internal error: resolving a module.Version with an empty path")
	}

	if m.Path == modload.Target.Path && m.Version != modload.Target.Version {
		reportError(q, &modload.QueryMatchesMainModuleError{
			Pattern: q.pattern,
			Query:   q.version,
		})
		return
	}

	vr, ok := r.resolvedVersion[m.Path]
	if ok && vr.version != m.Version {
		reportConflict(q, m, vr)
		return
	}
	r.resolvedVersion[m.Path] = versionReason{m.Version, q}
	q.resolved = append(q.resolved, m)
}

// updateBuildList updates the module loader's global build list to be
// consistent with r.resolvedVersion, and to include additional modules
// provided that they do not conflict with the resolved versions.
//
// If the additional modules conflict with the resolved versions, they will be
// downgraded to a non-conflicting version (possibly "none").
func (r *resolver) updateBuildList(ctx context.Context, additions []module.Version) (changed bool) {
	if len(additions) == 0 && len(r.resolvedVersion) == r.buildListResolvedVersions {
		return false
	}

	defer base.ExitIfErrors()

	resolved := make([]module.Version, 0, len(r.resolvedVersion))
	for mPath, rv := range r.resolvedVersion {
		if mPath != modload.Target.Path {
			resolved = append(resolved, module.Version{Path: mPath, Version: rv.version})
		}
	}

	if err := modload.EditBuildList(ctx, additions, resolved); err != nil {
		var constraint *modload.ConstraintError
		if !errors.As(err, &constraint) {
			base.Errorf("go get: %v", err)
			return false
		}

		reason := func(m module.Version) string {
			rv, ok := r.resolvedVersion[m.Path]
			if !ok {
				panic(fmt.Sprintf("internal error: can't find reason for requirement on %v", m))
			}
			return rv.reason.ResolvedString(module.Version{Path: m.Path, Version: rv.version})
		}
		for _, c := range constraint.Conflicts {
			base.Errorf("go get: %v requires %v, not %v", reason(c.Source), c.Dep, reason(c.Constraint))
		}
		return false
	}

	buildList := modload.LoadAllModules(ctx)
	r.buildListResolvedVersions = len(r.resolvedVersion)
	if reflect.DeepEqual(r.buildList, buildList) {
		return false
	}
	r.buildList = buildList
	r.buildListVersion = make(map[string]string, len(r.buildList))
	for _, m := range r.buildList {
		r.buildListVersion[m.Path] = m.Version
	}
	return true
}

func reqsFromGoMod(f *modfile.File) []module.Version {
	reqs := make([]module.Version, len(f.Require))
	for i, r := range f.Require {
		reqs[i] = r.Mod
	}
	return reqs
}

// isNoSuchModuleVersion reports whether err indicates that the requested module
// does not exist at the requested version, either because the module does not
// exist at all or because it does not include that specific version.
func isNoSuchModuleVersion(err error) bool {
	var noMatch *modload.NoMatchingVersionError
	return errors.Is(err, os.ErrNotExist) || errors.As(err, &noMatch)
}

// isNoSuchPackageVersion reports whether err indicates that the requested
// package does not exist at the requested version, either because no module
// that could contain it exists at that version, or because every such module
// that does exist does not actually contain the package.
func isNoSuchPackageVersion(err error) bool {
	var noPackage *modload.PackageNotInModuleError
	return isNoSuchModuleVersion(err) || errors.As(err, &noPackage)
}
