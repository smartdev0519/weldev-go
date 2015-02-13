package gc

import (
	"cmd/internal/obj"
	"strconv"
	"strings"
)

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (n *Node) Line() string {
	return obj.Linklinefmt(Ctxt, int(n.Lineno), false, false)
}

func atoi(s string) int {
	// NOTE: Not strconv.Atoi, accepts hex and octal prefixes.
	n, _ := strconv.ParseInt(s, 0, 0)
	return int(n)
}

func isalnum(c int) bool {
	return isalpha(c) || isdigit(c)
}

func isalpha(c int) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

func isdigit(c int) bool {
	return '0' <= c && c <= '9'
}

func plan9quote(s string) string {
	if s == "" {
		goto needquote
	}
	for i := 0; i < len(s); i++ {
		if s[i] <= ' ' || s[i] == '\'' {
			goto needquote
		}
	}
	return s

needquote:
	return "'" + strings.Replace(s, "'", "''", -1) + "'"
}

// simulation of int(*s++) in C
func intstarstringplusplus(s string) (int, string) {
	if s == "" {
		return 0, ""
	}
	return int(s[0]), s[1:]
}

// strings.Compare, introduced in Go 1.5.
func stringsCompare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
