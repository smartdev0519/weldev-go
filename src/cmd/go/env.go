// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var cmdEnv = &Command{
	Run:       runEnv,
	UsageLine: "env [var ...]",
	Short:     "print Go environment information",
	Long: `
Env prints Go environment information.

By default env prints information as a shell script
(on Windows, a batch file).  If one or more variable
names is given as arguments,  env prints the value of
each named variable on its own line.
	`,
}

type envVar struct {
	name, value string
}

func mkEnv() []envVar {
	var b builder
	b.init()

	env := []envVar{
		{"GOARCH", goarch},
		{"GOBIN", gobin},
		{"GOCHAR", archChar},
		{"GOEXE", exeSuffix},
		{"GOHOSTARCH", runtime.GOARCH},
		{"GOHOSTOS", runtime.GOOS},
		{"GOOS", goos},
		{"GOPATH", os.Getenv("GOPATH")},
		{"GOROOT", goroot},
		{"GOTOOLDIR", toolDir},
	}

	if goos != "plan9" {
		cmd := b.gccCmd(".")
		env = append(env, envVar{"CC", cmd[0]})
		env = append(env, envVar{"GOGCCFLAGS", strings.Join(cmd[3:], " ")})
	}

	if buildContext.CgoEnabled {
		env = append(env, envVar{"CGO_ENABLED", "1"})
	} else {
		env = append(env, envVar{"CGO_ENABLED", "0"})
	}

	return env
}

func findEnv(env []envVar, name string) string {
	for _, e := range env {
		if e.name == name {
			return e.value
		}
	}
	return ""
}

func runEnv(cmd *Command, args []string) {
	env := mkEnv()
	if len(args) > 0 {
		for _, name := range args {
			fmt.Printf("%s\n", findEnv(env, name))
		}
		return
	}

	switch runtime.GOOS {
	default:
		for _, e := range env {
			fmt.Printf("%s=\"%s\"\n", e.name, e.value)
		}
	case "plan9":
		for _, e := range env {
			fmt.Printf("%s='%s'\n", e.name, strings.Replace(e.value, "'", "''", -1))
		}
	case "windows":
		for _, e := range env {
			fmt.Printf("set %s=%s\n", e.name, e.value)
		}
	}
}
