// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepath

import (
	"strings"
	"syscall"
)

func isSlash(c uint8) bool {
	return c == '\\' || c == '/'
}

// IsAbs returns true if the path is absolute.
func IsAbs(path string) (b bool) {
	l := volumeNameLen(path)
	if l == 0 {
		return false
	}
	path = path[l:]
	if path == "" {
		return false
	}
	return isSlash(path[0])
}

// volumeNameLen returns length of the leading volume name on Windows.
// It returns 0 elsewhere.
func volumeNameLen(path string) int {
	if len(path) < 2 {
		return 0
	}
	// with drive letter
	c := path[0]
	if path[1] == ':' && ('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z') {
		return 2
	}
	// is it UNC
	if l := len(path); l >= 5 && isSlash(path[0]) && isSlash(path[1]) &&
		!isSlash(path[2]) && path[2] != '.' {
		// first, leading `\\` and next shouldn't be `\`. its server name.
		for n := 3; n < l-1; n++ {
			// second, next '\' shouldn't be repeated.
			if isSlash(path[n]) {
				n++
				// third, following something characters. its share name.
				if !isSlash(path[n]) {
					if path[n] == '.' {
						break
					}
					for ; n < l; n++ {
						if isSlash(path[n]) {
							break
						}
					}
					return n
				}
				break
			}
		}
	}
	return 0
}

// HasPrefix exists for historical compatibility and should not be used.
func HasPrefix(p, prefix string) bool {
	if strings.HasPrefix(p, prefix) {
		return true
	}
	return strings.HasPrefix(strings.ToLower(p), strings.ToLower(prefix))
}

func splitList(path string) []string {
	// The same implementation is used in LookPath in os/exec;
	// consider changing os/exec when changing this.

	if path == "" {
		return []string{}
	}

	// Split path, respecting but preserving quotes.
	list := []string{}
	start := 0
	quo := false
	for i := 0; i < len(path); i++ {
		switch c := path[i]; {
		case c == '"':
			quo = !quo
		case c == ListSeparator && !quo:
			list = append(list, path[start:i])
			start = i + 1
		}
	}
	list = append(list, path[start:])

	// Remove quotes.
	for i, s := range list {
		if strings.Contains(s, `"`) {
			list[i] = strings.Replace(s, `"`, ``, -1)
		}
	}

	return list
}

func abs(path string) (string, error) {
	return syscall.FullPath(path)
}

func join(elem []string) string {
	for i, e := range elem {
		if e != "" {
			return joinNonEmpty(elem[i:])
		}
	}
	return ""
}

// joinNonEmpty is like join, but it assumes that the first element is non-empty.
func joinNonEmpty(elem []string) string {
	// The following logic prevents Join from inadvertently creating a
	// UNC path on Windows. Unless the first element is a UNC path, Join
	// shouldn't create a UNC path. See golang.org/issue/9167.
	p := Clean(strings.Join(elem, string(Separator)))
	if !isUNC(p) {
		return p
	}
	// p == UNC only allowed when the first element is a UNC path.
	head := Clean(elem[0])
	if isUNC(head) {
		return p
	}
	// head + tail == UNC, but joining two non-UNC paths should not result
	// in a UNC path. Undo creation of UNC path.
	tail := Clean(strings.Join(elem[1:], string(Separator)))
	if head[len(head)-1] == Separator {
		return head + tail
	}
	return head + string(Separator) + tail
}

// isUNC returns true if path is a UNC path.
func isUNC(path string) bool {
	return volumeNameLen(path) > 2
}
