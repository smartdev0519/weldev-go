#!/bin/bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e
GOBIN="${GOBIN:-$HOME/bin}"
export MAKEFLAGS=-j4

if ! test -f $GOROOT/include/u.h
then
	echo '$GOROOT is not set correctly or not exported' 1>&2
	exit 1
fi

if ! test -d $GOBIN
then
	echo '$GOBIN is not a directory or does not exist' 1>&2
	echo 'create it or set $GOBIN differently' 1>&2
	exit 1
fi

bash clean.bash

rm -f $GOBIN/quietgcc
CC=${CC:-gcc}
sed -e "s|@CC@|$CC|" < quietgcc.bash > $GOBIN/quietgcc
chmod +x $GOBIN/quietgcc

for i in lib9 libbio libmach cmd pkg libcgo cmd/cgo cmd/ebnflint cmd/godoc cmd/gofmt cmd/goyacc cmd/hgpatch
do
	case "$i-$GOOS" in
	libcgo-nacl)
		;;
	*)
		# The ( ) here are to preserve the current directory
		# for the next round despite the cd $i below.
		# set -e does not apply to ( ) so we must explicitly
		# test the exit status.
		(
			echo; echo; echo %%%% making $i %%%%; echo
			cd $i
			case $i in
			cmd)
				bash make.bash
				;;
			*)
				make install
			esac
		)  || exit 1
	esac
done

case "`uname`" in
Darwin)
	echo;
	echo %%% run sudo.bash to install debuggers
	echo
esac
