#!/usr/bin/env bash
# Copyright 2015 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# This directory is intended to test the use of Go with sanitizers
# like msan, asan, etc.  See https://github.com/google/sanitizers .

set -e

# The sanitizers were originally developed with clang, so prefer it.
CC=cc
if test "$(type -p clang)" != ""; then
  CC=clang
fi
export CC

if $CC -fsanitize=memory 2>&1 | grep "unrecognized" >& /dev/null; then
  echo "skipping msan test: -fsanitize=memory not supported"
  exit 0
fi

# The memory sanitizer in versions of clang before 3.6 don't work with Go.
if $CC --version | grep clang >& /dev/null; then
  ver=$($CC --version | sed -e 's/.* version \([0-9.-]*\).*/\1/')
  major=$(echo $ver | sed -e 's/\([0-9]*\).*/\1/')
  minor=$(echo $ver | sed -e 's/[0-9]*\.\([0-9]*\).*/\1/')
  if test $major -lt 3 || test $major -eq 3 -a $minor -lt 6; then
    echo "skipping msan test; clang version $major.$minor older than 3.6"
    exit 0
  fi
fi

go run msan.go
