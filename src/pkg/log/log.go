// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple logging package. It defines a type, Logger, with simple
// methods for formatting output to one or two destinations. It also
// has a predefined 'standard' Logger accessible through helper
// functions Print[f|ln], Exit[f|ln], and Panic[f|ln], which are
// easier to use than creating a Logger manually.  That logger writes
// to standard error and prints the date and time of each logged
// message.
// The Exit functions call os.Exit(1) after writing the log message.
// The Panic functions call panic after writing the log message.
package log

import (
	"fmt"
	"io"
	"runtime"
	"os"
	"time"
)

// These flags define the output Loggers produce.
const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/0123 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota // the date: 2009/0123
	Ltime                     // the time: 01:23:23
	Lmicroseconds             // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                 // full file name and line number: /a/b/c/d.go:23
	Lshortfile                // final file name element and line number: d.go:23. overrides Llongfile
	lallBits      = Ldate | Ltime | Lmicroseconds | Llongfile | Lshortfile
)

// Logger represents an active logging object.
type Logger struct {
	out    io.Writer // destination for output
	prefix string    // prefix to write at beginning of each line
	flag   int       // properties
}

// New creates a new Logger.   The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{out, prefix, flag}
}

var (
	std    = New(os.Stderr, "", Ldate|Ltime)
	stdout = New(os.Stdout, "", Ldate|Ltime) // Deprecated.
)

// Cheap integer to fixed-width decimal ASCII.  Use a negative width to avoid zero-padding
func itoa(i int, wid int) string {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		return "0"
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}

	return string(b[bp:])
}

func (l *Logger) formatHeader(ns int64, calldepth int) string {
	h := l.prefix
	if l.flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		t := time.SecondsToLocalTime(ns / 1e9)
		if l.flag&(Ldate) != 0 {
			h += itoa(int(t.Year), 4) + "/" + itoa(t.Month, 2) + "/" + itoa(t.Day, 2) + " "
		}
		if l.flag&(Ltime|Lmicroseconds) != 0 {
			h += itoa(t.Hour, 2) + ":" + itoa(t.Minute, 2) + ":" + itoa(t.Second, 2)
			if l.flag&Lmicroseconds != 0 {
				h += "." + itoa(int(ns%1e9)/1e3, 6)
			}
			h += " "
		}
	}
	if l.flag&(Lshortfile|Llongfile) != 0 {
		_, file, line, ok := runtime.Caller(calldepth)
		if ok {
			if l.flag&Lshortfile != 0 {
				short := file
				for i := len(file) - 1; i > 0; i-- {
					if file[i] == '/' {
						short = file[i+1:]
						break
					}
				}
				file = short
			}
		} else {
			file = "???"
			line = 0
		}
		h += file + ":" + itoa(line, -1) + ": "
	}
	return h
}

// Output writes the output for a logging event.  The string s contains the text to print after
// the time stamp;  calldepth is used to recover the PC.  It is provided for generality, although
// at the moment on all pre-defined paths it will be 2.
func (l *Logger) Output(calldepth int, s string) os.Error {
	now := time.Nanoseconds() // get this early.
	newline := "\n"
	if len(s) > 0 && s[len(s)-1] == '\n' {
		newline = ""
	}
	s = l.formatHeader(now, calldepth+1) + s + newline
	_, err := io.WriteString(l.out, s)
	return err
}

// Printf prints to the logger in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf(format, v...))
}

// Print prints to the logger in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) { l.Output(2, fmt.Sprint(v...)) }

// Println prints to the logger in the manner of fmt.Println.
func (l *Logger) Println(v ...interface{}) { l.Output(2, fmt.Sprintln(v...)) }

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	std.out = w
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	std.flag = flag & lallBits
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	std.prefix = prefix
}

// These functions write to the standard logger.

// Print prints to the standard logger in the manner of fmt.Print.
func Print(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
}

// Printf prints to the standard logger in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(format, v...))
}

// Println prints to the standard logger in the manner of fmt.Println.
func Println(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
}

// Exit is equivalent to Print() followed by a call to os.Exit(1).
func Exit(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Exitf is equivalent to Printf() followed by a call to os.Exit(1).
func Exitf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Exitln is equivalent to Println() followed by a call to os.Exit(1).
func Exitln(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(2, s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(2, s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	std.Output(2, s)
	panic(s)
}

// Everything from here on is deprecated and will be removed after the next release.

// Logf is analogous to Printf() for a Logger.
// Deprecated.
func (l *Logger) Logf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf(format, v...))
}

// Log is analogous to Print() for a Logger.
// Deprecated.
func (l *Logger) Log(v ...interface{}) { l.Output(2, fmt.Sprintln(v...)) }

// Stdout is a helper function for easy logging to stdout. It is analogous to Print().
// Deprecated.
func Stdout(v ...interface{}) { stdout.Output(2, fmt.Sprint(v...)) }

// Stderr is a helper function for easy logging to stderr. It is analogous to Fprint(os.Stderr).
// Deprecated.
func Stderr(v ...interface{}) { std.Output(2, fmt.Sprintln(v...)) }

// Stdoutf is a helper functions for easy formatted logging to stdout. It is analogous to Printf().
// Deprecated.
func Stdoutf(format string, v ...interface{}) { stdout.Output(2, fmt.Sprintf(format, v...)) }

// Stderrf is a helper function for easy formatted logging to stderr. It is analogous to Fprintf(os.Stderr).
// Deprecated.
func Stderrf(format string, v ...interface{}) { std.Output(2, fmt.Sprintf(format, v...)) }

// Crash is equivalent to Stderr() followed by a call to panic().
// Deprecated.
func Crash(v ...interface{}) { Panicln(v...) }

// Crashf is equivalent to Stderrf() followed by a call to panic().
// Deprecated.
func Crashf(format string, v ...interface{}) { Panicf(format, v...) }
