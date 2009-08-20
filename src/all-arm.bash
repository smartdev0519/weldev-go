#!/bin/bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e
bash make-arm.bash

# TODO(kaib): add in proper tests
#bash run.bash
5g -o hello.5 /home/kaib/work/go/hello.go
5l -o 5.out hello.5
qemu-arm -cpu cortex-a8 5.out|grep -q "Hello World"
