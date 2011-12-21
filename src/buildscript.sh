#!/bin/sh
# Copyright 2011 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

. ./buildinfo.sh

for sys in $GOOSARCHES
do
	export GOOS=$(echo $sys | sed 's/_.*//')
	export GOARCH=$(echo $sys | sed 's/.*_//')
	targ=buildscript_${GOOS}_$GOARCH.sh
	rm -f $targ

	(echo '#!/usr/bin/env bash
# AUTO-GENERATED by buildscript.sh; DO NOT EDIT.
# This script builds the go command (written in Go),
# and then the go command can build the rest of the tree.

export GOOS='$GOOS'
export GOARCH='$GOARCH'
export WORK=$(mktemp -d -t go-build.XXXXXX)
trap "rm -rf $WORK" EXIT SIGINT SIGTERM
set -e

'
	# Save script printed by go install but make shell safe
	# by quoting variable expansions.  On Windows, rewrite
	# \ paths into / paths.  This avoids the \ being interpreted
	# as a shell escape but also makes sure that we generate the
	# same scripts on Unix and Windows systems.
	go install -a -n cmd/go | sed '
		s/$GOBIN/"$GOBIN"/g
		s/$GOROOT/"$GOROOT"/g
		s/$WORK/"$WORK"/g
		s;\\;/;g
	'
	)>$targ
	chmod +x $targ
done
