#!/usr/bin/env bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Generate Go code listing errors and other #defined constant
# values (ENAMETOOLONG etc.), by asking the preprocessor
# about the definitions.

unset LANG
export LC_ALL=C
export LC_CTYPE=C

case "$GOARCH" in
arm)
	GCC=arm-gcc
	;;
*)
	GCC=gcc
	;;
esac

uname=$(uname)

includes_Linux='
#define _LARGEFILE_SOURCE
#define _LARGEFILE64_SOURCE
#define _FILE_OFFSET_BITS 64
#define _GNU_SOURCE

#include <sys/types.h>
#include <sys/epoll.h>
#include <linux/ptrace.h>
#include <linux/wait.h>
'

includes_Darwin='
#define __DARWIN_UNIX03 0
#define KERNEL
#define _DARWIN_USE_64_BIT_INODE
#include <sys/wait.h>
#include <sys/event.h>
'

includes_FreeBSD='
#include <sys/wait.h>
#include <sys/event.h>
'

includes='
#include <sys/types.h>
#include <fcntl.h>
#include <dirent.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <netinet/ip6.h>
#include <netinet/tcp.h>
#include <errno.h>
#include <sys/signal.h>
#include <signal.h>
'

# Write godefs input.
(
	indirect="includes_$(uname)"
	echo "${!indirect} $includes"
	echo
	echo 'enum {'

	# The gcc command line prints all the #defines
	# it encounters while processing the input
	echo "${!indirect} $includes" | $GCC -x c - -E -dM |
	awk '
		$1 != "#define" || $2 ~ /\(/ {next}
		
		$2 ~ /^(SIGEV_|SIGSTKSZ|SIGRT(MIN|MAX))/ {next}

		$2 ~ /^E[A-Z0-9_]+$/ ||
		$2 ~ /^SIG[^_]/ ||
		$2 ~ /^(AF|SOCK|SO|SOL|IPPROTO|IP|IPV6|TCP|EVFILT|EV|SHUT|PROT|MAP)_/ ||
		$2 == "SOMAXCONN" ||
		$2 == "NAME_MAX" ||
		$2 ~ /^(O|F|FD|NAME|S|PTRACE)_/ ||
		$2 ~ /^W[A-Z0-9]+$/ {printf("\t$%s = %s,\n", $2, $2)}
		
		$2 ~ /^__W[A-Z0-9]+$/ {printf("\t$%s = %s,\n", substr($2,3), $2)}
		
		{next}
	' | sort

	echo '};'
) >_const.c

# Pull out just the error names for later.
errors=$(
	echo '#include <errno.h>' | $GCC -x c - -E -dM |
	awk '$1=="#define" && $2 ~ /^E[A-Z0-9_]+$/ { print $2 }' |
	sort
)

echo '// mkerrors.sh' "$@"
echo '// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT'
echo
godefs -gsyscall "$@" _const.c

# Run C program to print error strings.
(
	/bin/echo "
#include <stdio.h>
#include <errno.h>
#include <ctype.h>
#include <string.h>

#define nelem(x) (sizeof(x)/sizeof((x)[0]))

enum { A = 'A', Z = 'Z', a = 'a', z = 'z' }; // avoid need for single quotes below

int errors[] = {
"
	for i in $errors
	do
		/bin/echo '	'$i,
	done

	# Use /bin/echo to avoid builtin echo,
	# which interprets \n itself
	/bin/echo '
};

int
main(void)
{
	int i, j, e;
	char buf[1024];

	printf("\n\n// Error table\n");
	printf("var errors = [...]string {\n");
	for(i=0; i<nelem(errors); i++) {
		e = errors[i];
		for(j=0; j<i; j++)
			if(errors[j] == e)	// duplicate value
				goto next;
		strcpy(buf, strerror(e));
		// lowercase first letter: Bad -> bad, but STREAM -> STREAM.
		if(A <= buf[0] && buf[0] <= Z && a <= buf[1] && buf[1] <= z)
			buf[0] += a - A;
		printf("\t%d: \"%s\",\n", e, buf);
	next:;
	}
	printf("}\n\n");
	return 0;
}

'
) >_errors.c

gcc -o _errors _errors.c && ./_errors && rm -f _errors.c _errors _const.c
