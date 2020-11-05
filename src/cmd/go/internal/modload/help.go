// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import "cmd/go/internal/base"

// TODO(rsc): The "module code layout" section needs to be written.

var HelpModules = &base.Command{
	UsageLine: "modules",
	Short:     "modules, module versions, and more",
	Long: `
A module is a collection of related Go packages.
Modules are the unit of source code interchange and versioning.
The go command has direct support for working with modules,
including recording and resolving dependencies on other modules.
Modules replace the old GOPATH-based approach to specifying
which source files are used in a given build.

Module support

The go command includes support for Go modules. Module-aware mode is active
by default whenever a go.mod file is found in the current directory or in
any parent directory.

The quickest way to take advantage of module support is to check out your
repository, create a go.mod file (described in the next section) there, and run
go commands from within that file tree.

For more fine-grained control, the go command continues to respect
a temporary environment variable, GO111MODULE, which can be set to one
of three string values: off, on, or auto (the default).
If GO111MODULE=on, then the go command requires the use of modules,
never consulting GOPATH. We refer to this as the command
being module-aware or running in "module-aware mode".
If GO111MODULE=off, then the go command never uses
module support. Instead it looks in vendor directories and GOPATH
to find dependencies; we now refer to this as "GOPATH mode."
If GO111MODULE=auto or is unset, then the go command enables or disables
module support based on the current directory.
Module support is enabled only when the current directory contains a
go.mod file or is below a directory containing a go.mod file.

In module-aware mode, GOPATH no longer defines the meaning of imports
during a build, but it still stores downloaded dependencies (in GOPATH/pkg/mod)
and installed commands (in GOPATH/bin, unless GOBIN is set).

Defining a module

A module is defined by a tree of Go source files with a go.mod file
in the tree's root directory. The directory containing the go.mod file
is called the module root. Typically the module root will also correspond
to a source code repository root (but in general it need not).
The module is the set of all Go packages in the module root and its
subdirectories, but excluding subtrees with their own go.mod files.

The "module path" is the import path prefix corresponding to the module root.
The go.mod file defines the module path and lists the specific versions
of other modules that should be used when resolving imports during a build,
by giving their module paths and versions.

For example, this go.mod declares that the directory containing it is the root
of the module with path example.com/m, and it also declares that the module
depends on specific versions of golang.org/x/text and gopkg.in/yaml.v2:

	module example.com/m

	require (
		golang.org/x/text v0.3.0
		gopkg.in/yaml.v2 v2.1.0
	)

The go.mod file can also specify replacements and excluded versions
that only apply when building the module directly; they are ignored
when the module is incorporated into a larger build.
For more about the go.mod file, see 'go help go.mod'.

To start a new module, simply create a go.mod file in the root of the
module's directory tree, containing only a module statement.
The 'go mod init' command can be used to do this:

	go mod init example.com/m

In a project already using an existing dependency management tool like
godep, glide, or dep, 'go mod init' will also add require statements
matching the existing configuration.

Once the go.mod file exists, no additional steps are required:
go commands like 'go build', 'go test', or even 'go list' will automatically
add new dependencies as needed to satisfy imports.

The main module and the build list

The "main module" is the module containing the directory where the go command
is run. The go command finds the module root by looking for a go.mod in the
current directory, or else the current directory's parent directory,
or else the parent's parent directory, and so on.

The main module's go.mod file defines the precise set of packages available
for use by the go command, through require, replace, and exclude statements.
Dependency modules, found by following require statements, also contribute
to the definition of that set of packages, but only through their go.mod
files' require statements: any replace and exclude statements in dependency
modules are ignored. The replace and exclude statements therefore allow the
main module complete control over its own build, without also being subject
to complete control by dependencies.

The set of modules providing packages to builds is called the "build list".
The build list initially contains only the main module. Then the go command
adds to the list the exact module versions required by modules already
on the list, recursively, until there is nothing left to add to the list.
If multiple versions of a particular module are added to the list,
then at the end only the latest version (according to semantic version
ordering) is kept for use in the build.

The 'go list' command provides information about the main module
and the build list. For example:

	go list -m              # print path of main module
	go list -m -f={{.Dir}}  # print root directory of main module
	go list -m all          # print build list

Maintaining module requirements

The go.mod file is meant to be readable and editable by both programmers and
tools. Most updates to dependencies can be performed using "go get" and
"go mod tidy". Other module-aware build commands may be invoked using the
-mod=mod flag to automatically add missing requirements and fix inconsistencies.

The "go get" command updates go.mod to change the module versions used in a
build. An upgrade of one module may imply upgrading others, and similarly a
downgrade of one module may imply downgrading others. The "go get" command
makes these implied changes as well. See "go help module-get".

The "go mod" command provides other functionality for use in maintaining
and understanding modules and go.mod files. See "go help mod", particularly
"go help mod tidy" and "go help mod edit".

As part of maintaining the require statements in go.mod, the go command
tracks which ones provide packages imported directly by the current module
and which ones provide packages only used indirectly by other module
dependencies. Requirements needed only for indirect uses are marked with a
"// indirect" comment in the go.mod file. Indirect requirements may be
automatically removed from the go.mod file once they are implied by other
direct requirements. Indirect requirements only arise when using modules
that fail to state some of their own dependencies or when explicitly
upgrading a module's dependencies ahead of its own stated requirements.

The -mod build flag provides additional control over the updating and use of
go.mod for commands that build packages like "go build" and "go test".

If invoked with -mod=readonly (the default in most situations), the go command
reports an error if a package named on the command line or an imported package
is not provided by any module in the build list computed from the main module's
requirements. The go command also reports an error if a module's checksum is
missing from go.sum (see Module downloading and verification). Either go.mod or
go.sum must be updated in these situations.

If invoked with -mod=mod, the go command automatically updates go.mod and
go.sum, fixing inconsistencies and adding missing requirements and checksums
as needed. If the go command finds an unfamiliar import, it looks up the
module containing that import and adds a requirement for the latest version
of that module to go.mod. In most cases, therefore, one may add an import to
source code and run "go build", "go test", or even "go list" with -mod=mod:
as part of analyzing the package, the go command will resolve the import and
update the go.mod file.

If invoked with -mod=vendor, the go command loads packages from the main
module's vendor directory instead of downloading modules to and loading packages
from the module cache. The go command assumes the vendor directory holds
correct copies of dependencies, and it does not compute the set of required
module versions from go.mod files. However, the go command does check that
vendor/modules.txt (generated by "go mod vendor") contains metadata consistent
with go.mod.

If the go command is not invoked with a -mod flag, and the vendor directory
is present, and the "go" version in go.mod is 1.14 or higher, the go command
will act as if it were invoked with -mod=vendor. Otherwise, the -mod flag
defaults to -mod=readonly.

Note that neither "go get" nor the "go mod" subcommands accept the -mod flag.

Pseudo-versions

The go.mod file and the go command more generally use semantic versions as
the standard form for describing module versions, so that versions can be
compared to determine which should be considered earlier or later than another.
A module version like v1.2.3 is introduced by tagging a revision in the
underlying source repository. Untagged revisions can be referred to
using a "pseudo-version" like v0.0.0-yyyymmddhhmmss-abcdefabcdef,
where the time is the commit time in UTC and the final suffix is the prefix
of the commit hash. The time portion ensures that two pseudo-versions can
be compared to determine which happened later, the commit hash identifes
the underlying commit, and the prefix (v0.0.0- in this example) is derived from
the most recent tagged version in the commit graph before this commit.

There are three pseudo-version forms:

vX.0.0-yyyymmddhhmmss-abcdefabcdef is used when there is no earlier
versioned commit with an appropriate major version before the target commit.
(This was originally the only form, so some older go.mod files use this form
even for commits that do follow tags.)

vX.Y.Z-pre.0.yyyymmddhhmmss-abcdefabcdef is used when the most
recent versioned commit before the target commit is vX.Y.Z-pre.

vX.Y.(Z+1)-0.yyyymmddhhmmss-abcdefabcdef is used when the most
recent versioned commit before the target commit is vX.Y.Z.

Pseudo-versions never need to be typed by hand: the go command will accept
the plain commit hash and translate it into a pseudo-version (or a tagged
version if available) automatically. This conversion is an example of a
module query.

Module queries

The go command accepts a "module query" in place of a module version
both on the command line and in the main module's go.mod file.
(After evaluating a query found in the main module's go.mod file,
the go command updates the file to replace the query with its result.)

A fully-specified semantic version, such as "v1.2.3",
evaluates to that specific version.

A semantic version prefix, such as "v1" or "v1.2",
evaluates to the latest available tagged version with that prefix.

A semantic version comparison, such as "<v1.2.3" or ">=v1.5.6",
evaluates to the available tagged version nearest to the comparison target
(the latest version for < and <=, the earliest version for > and >=).

The string "latest" matches the latest available tagged version,
or else the underlying source repository's latest untagged revision.

The string "upgrade" is like "latest", but if the module is
currently required at a later version than the version "latest"
would select (for example, a newer pre-release version), "upgrade"
will select the later version instead.

The string "patch" matches the latest available tagged version
of a module with the same major and minor version numbers as the
currently required version. If no version is currently required,
"patch" is equivalent to "latest".

A revision identifier for the underlying source repository, such as
a commit hash prefix, revision tag, or branch name, selects that
specific code revision. If the revision is also tagged with a
semantic version, the query evaluates to that semantic version.
Otherwise the query evaluates to a pseudo-version for the commit.
Note that branches and tags with names that are matched by other
query syntax cannot be selected this way. For example, the query
"v2" means the latest version starting with "v2", not the branch
named "v2".

All queries prefer release versions to pre-release versions.
For example, "<v1.2.3" will prefer to return "v1.2.2"
instead of "v1.2.3-pre1", even though "v1.2.3-pre1" is nearer
to the comparison target.

Module versions disallowed by exclude statements in the
main module's go.mod are considered unavailable and cannot
be returned by queries.

For example, these commands are all valid:

	go get github.com/gorilla/mux@latest    # same (@latest is default for 'go get')
	go get github.com/gorilla/mux@v1.6.2    # records v1.6.2
	go get github.com/gorilla/mux@e3702bed2 # records v1.6.2
	go get github.com/gorilla/mux@c856192   # records v0.0.0-20180517173623-c85619274f5d
	go get github.com/gorilla/mux@master    # records current meaning of master

Module compatibility and semantic versioning

The go command requires that modules use semantic versions and expects that
the versions accurately describe compatibility: it assumes that v1.5.4 is a
backwards-compatible replacement for v1.5.3, v1.4.0, and even v1.0.0.
More generally the go command expects that packages follow the
"import compatibility rule", which says:

"If an old package and a new package have the same import path,
the new package must be backwards compatible with the old package."

Because the go command assumes the import compatibility rule,
a module definition can only set the minimum required version of one
of its dependencies: it cannot set a maximum or exclude selected versions.
Still, the import compatibility rule is not a guarantee: it may be that
v1.5.4 is buggy and not a backwards-compatible replacement for v1.5.3.
Because of this, the go command never updates from an older version
to a newer version of a module unasked.

In semantic versioning, changing the major version number indicates a lack
of backwards compatibility with earlier versions. To preserve import
compatibility, the go command requires that modules with major version v2
or later use a module path with that major version as the final element.
For example, version v2.0.0 of example.com/m must instead use module path
example.com/m/v2, and packages in that module would use that path as
their import path prefix, as in example.com/m/v2/sub/pkg. Including the
major version number in the module path and import paths in this way is
called "semantic import versioning". Pseudo-versions for modules with major
version v2 and later begin with that major version instead of v0, as in
v2.0.0-20180326061214-4fc5987536ef.

As a special case, module paths beginning with gopkg.in/ continue to use the
conventions established on that system: the major version is always present,
and it is preceded by a dot instead of a slash: gopkg.in/yaml.v1
and gopkg.in/yaml.v2, not gopkg.in/yaml and gopkg.in/yaml/v2.

The go command treats modules with different module paths as unrelated:
it makes no connection between example.com/m and example.com/m/v2.
Modules with different major versions can be used together in a build
and are kept separate by the fact that their packages use different
import paths.

In semantic versioning, major version v0 is for initial development,
indicating no expectations of stability or backwards compatibility.
Major version v0 does not appear in the module path, because those
versions are preparation for v1.0.0, and v1 does not appear in the
module path either.

Code written before the semantic import versioning convention
was introduced may use major versions v2 and later to describe
the same set of unversioned import paths as used in v0 and v1.
To accommodate such code, if a source code repository has a
v2.0.0 or later tag for a file tree with no go.mod, the version is
considered to be part of the v1 module's available versions
and is given an +incompatible suffix when converted to a module
version, as in v2.0.0+incompatible. The +incompatible tag is also
applied to pseudo-versions derived from such versions, as in
v2.0.1-0.yyyymmddhhmmss-abcdefabcdef+incompatible.

In general, having a dependency in the build list (as reported by 'go list -m all')
on a v0 version, pre-release version, pseudo-version, or +incompatible version
is an indication that problems are more likely when upgrading that
dependency, since there is no expectation of compatibility for those.

See https://research.swtch.com/vgo-import for more information about
semantic import versioning, and see https://semver.org/ for more about
semantic versioning.

Module code layout

For now, see https://research.swtch.com/vgo-module for information
about how source code in version control systems is mapped to
module file trees.

Module downloading and verification

The go command can fetch modules from a proxy or connect to source control
servers directly, according to the setting of the GOPROXY environment
variable (see 'go help env'). The default setting for GOPROXY is
"https://proxy.golang.org,direct", which means to try the
Go module mirror run by Google and fall back to a direct connection
if the proxy reports that it does not have the module (HTTP error 404 or 410).
See https://proxy.golang.org/privacy for the service's privacy policy.

If GOPROXY is set to the string "direct", downloads use a direct connection to
source control servers. Setting GOPROXY to "off" disallows downloading modules
from any source. Otherwise, GOPROXY is expected to be list of module proxy URLs
separated by either comma (,) or pipe (|) characters, which control error
fallback behavior. For each request, the go command tries each proxy in
sequence. If there is an error, the go command will try the next proxy in the
list if the error is a 404 or 410 HTTP response or if the current proxy is
followed by a pipe character, indicating it is safe to fall back on any error.

The GOPRIVATE and GONOPROXY environment variables allow bypassing
the proxy for selected modules. See 'go help module-private' for details.

No matter the source of the modules, the go command checks downloads against
known checksums, to detect unexpected changes in the content of any specific
module version from one day to the next. This check first consults the current
module's go.sum file but falls back to the Go checksum database, controlled by
the GOSUMDB and GONOSUMDB environment variables. See 'go help module-auth'
for details.

See 'go help goproxy' for details about the proxy protocol and also
the format of the cached downloaded packages.

Modules and vendoring

When using modules, the go command typically satisfies dependencies by
downloading modules from their sources and using those downloaded copies
(after verification, as described in the previous section). Vendoring may
be used to allow interoperation with older versions of Go, or to ensure
that all files used for a build are stored together in a single file tree.

The command 'go mod vendor' constructs a directory named vendor in the main
module's root directory that contains copies of all packages needed to support
builds and tests of packages in the main module. 'go mod vendor' also
creates the file vendor/modules.txt that contains metadata about vendored
packages and module versions. This file should be kept consistent with go.mod:
when vendoring is used, 'go mod vendor' should be run after go.mod is updated.

If the vendor directory is present in the main module's root directory, it will
be used automatically if the "go" version in the main module's go.mod file is
1.14 or higher. Build commands like 'go build' and 'go test' will load packages
from the vendor directory instead of accessing the network or the local module
cache. To explicitly enable vendoring, invoke the go command with the flag
-mod=vendor. To disable vendoring, use the flag -mod=mod.

Unlike vendoring in GOPATH, the go command ignores vendor directories in
locations other than the main module's root directory.
	`,
}

var HelpGoMod = &base.Command{
	UsageLine: "go.mod",
	Short:     "the go.mod file",
	Long: `
A module version is defined by a tree of source files, with a go.mod
file in its root. When the go command is run, it looks in the current
directory and then successive parent directories to find the go.mod
marking the root of the main (current) module.

The go.mod file itself is line-oriented, with // comments but
no /* */ comments. Each line holds a single directive, made up of a
verb followed by arguments. For example:

	module my/thing
	go 1.12
	require other/thing v1.0.2
	require new/thing/v2 v2.3.4
	exclude old/thing v1.2.3
	replace bad/thing v1.4.5 => good/thing v1.4.5
	retract v1.5.6

The verbs are
	module, to define the module path;
	go, to set the expected language version;
	require, to require a particular module at a given version or later;
	exclude, to exclude a particular module version from use;
	replace, to replace a module version with a different module version; and
	retract, to indicate a previously released version should not be used.
Exclude and replace apply only in the main module's go.mod and are ignored
in dependencies.  See https://golang.org/ref/mod for details.

The leading verb can be factored out of adjacent lines to create a block,
like in Go imports:

	require (
		new/thing v2.3.4
		old/thing v1.2.3
	)

The go.mod file is designed both to be edited directly and to be
easily updated by tools. The 'go mod edit' command can be used to
parse and edit the go.mod file from programs and tools.
See 'go help mod edit'.

The go command automatically updates go.mod each time it uses the
module graph, to make sure go.mod always accurately reflects reality
and is properly formatted. For example, consider this go.mod file:

        module M

        require (
                A v1
                B v1.0.0
                C v1.0.0
                D v1.2.3
                E dev
        )

        exclude D v1.2.3

The update rewrites non-canonical version identifiers to semver form,
so A's v1 becomes v1.0.0 and E's dev becomes the pseudo-version for the
latest commit on the dev branch, perhaps v0.0.0-20180523231146-b3f5c0f6e5f1.

The update modifies requirements to respect exclusions, so the
requirement on the excluded D v1.2.3 is updated to use the next
available version of D, perhaps D v1.2.4 or D v1.3.0.

The update removes redundant or misleading requirements.
For example, if A v1.0.0 itself requires B v1.2.0 and C v1.0.0,
then go.mod's requirement of B v1.0.0 is misleading (superseded by
A's need for v1.2.0), and its requirement of C v1.0.0 is redundant
(implied by A's need for the same version), so both will be removed.
If module M contains packages that directly import packages from B or
C, then the requirements will be kept but updated to the actual
versions being used.

Finally, the update reformats the go.mod in a canonical formatting, so
that future mechanical changes will result in minimal diffs.

Because the module graph defines the meaning of import statements, any
commands that load packages also use and therefore update go.mod,
including go build, go get, go install, go list, go test, go mod graph,
go mod tidy, and go mod why.

The expected language version, set by the go directive, determines
which language features are available when compiling the module.
Language features available in that version will be available for use.
Language features removed in earlier versions, or added in later versions,
will not be available. Note that the language version does not affect
build tags, which are determined by the Go release being used.
	`,
}
