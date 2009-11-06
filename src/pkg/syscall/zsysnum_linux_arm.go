// hand generated

package syscall

const (
	SYS_SYSCALL_BASE	= 0;

	SYS_RESTART_SYSCALL		= (SYS_SYSCALL_BASE + 0);
	SYS_EXIT			= (SYS_SYSCALL_BASE + 1);
	SYS_FORK			= (SYS_SYSCALL_BASE + 2);
	SYS_READ			= (SYS_SYSCALL_BASE + 3);
	SYS_WRITE			= (SYS_SYSCALL_BASE + 4);
	SYS_OPEN			= (SYS_SYSCALL_BASE + 5);
	SYS_CLOSE			= (SYS_SYSCALL_BASE + 6);
	SYS_CREAT			= (SYS_SYSCALL_BASE + 8);
	SYS_LINK			= (SYS_SYSCALL_BASE + 9);
	SYS_UNLINK			= (SYS_SYSCALL_BASE + 10);
	SYS_EXECVE			= (SYS_SYSCALL_BASE + 11);
	SYS_CHDIR			= (SYS_SYSCALL_BASE + 12);
	SYS_TIME			= (SYS_SYSCALL_BASE + 13);
	SYS_MKNOD			= (SYS_SYSCALL_BASE + 14);
	SYS_CHMOD			= (SYS_SYSCALL_BASE + 15);
	SYS_LCHOWN			= (SYS_SYSCALL_BASE + 16);
	SYS_LSEEK			= (SYS_SYSCALL_BASE + 19);
	SYS_GETPID			= (SYS_SYSCALL_BASE + 20);
	SYS_MOUNT			= (SYS_SYSCALL_BASE + 21);
	SYS_UMOUNT			= (SYS_SYSCALL_BASE + 22);
	SYS_SETUID			= (SYS_SYSCALL_BASE + 23);
	SYS_GETUID			= (SYS_SYSCALL_BASE + 24);
	SYS_STIME			= (SYS_SYSCALL_BASE + 25);
	SYS_PTRACE			= (SYS_SYSCALL_BASE + 26);
	SYS_ALARM			= (SYS_SYSCALL_BASE + 27);
	SYS_PAUSE			= (SYS_SYSCALL_BASE + 29);
	SYS_UTIME			= (SYS_SYSCALL_BASE + 30);
	SYS_ACCESS			= (SYS_SYSCALL_BASE + 33);
	SYS_NICE			= (SYS_SYSCALL_BASE + 34);
	SYS_SYNC			= (SYS_SYSCALL_BASE + 36);
	SYS_KILL			= (SYS_SYSCALL_BASE + 37);
	SYS_RENAME			= (SYS_SYSCALL_BASE + 38);
	SYS_MKDIR			= (SYS_SYSCALL_BASE + 39);
	SYS_RMDIR			= (SYS_SYSCALL_BASE + 40);
	SYS_DUP				= (SYS_SYSCALL_BASE + 41);
	SYS_PIPE			= (SYS_SYSCALL_BASE + 42);
	SYS_TIMES			= (SYS_SYSCALL_BASE + 43);
	SYS_BRK				= (SYS_SYSCALL_BASE + 45);
	SYS_SETGID			= (SYS_SYSCALL_BASE + 46);
	SYS_GETGID			= (SYS_SYSCALL_BASE + 47);
	SYS_GETEUID			= (SYS_SYSCALL_BASE + 49);
	SYS_GETEGID			= (SYS_SYSCALL_BASE + 50);
	SYS_ACCT			= (SYS_SYSCALL_BASE + 51);
	SYS_UMOUNT2			= (SYS_SYSCALL_BASE + 52);
	SYS_IOCTL			= (SYS_SYSCALL_BASE + 54);
	SYS_FCNTL			= (SYS_SYSCALL_BASE + 55);
	SYS_SETPGID			= (SYS_SYSCALL_BASE + 57);
	SYS_UMASK			= (SYS_SYSCALL_BASE + 60);
	SYS_CHROOT			= (SYS_SYSCALL_BASE + 61);
	SYS_USTAT			= (SYS_SYSCALL_BASE + 62);
	SYS_DUP2			= (SYS_SYSCALL_BASE + 63);
	SYS_GETPPID			= (SYS_SYSCALL_BASE + 64);
	SYS_GETPGRP			= (SYS_SYSCALL_BASE + 65);
	SYS_SETSID			= (SYS_SYSCALL_BASE + 66);
	SYS_SIGACTION			= (SYS_SYSCALL_BASE + 67);
	SYS_SETREUID			= (SYS_SYSCALL_BASE + 70);
	SYS_SETREGID			= (SYS_SYSCALL_BASE + 71);
	SYS_SIGSUSPEND			= (SYS_SYSCALL_BASE + 72);
	SYS_SIGPENDING			= (SYS_SYSCALL_BASE + 73);
	SYS_SETHOSTNAME			= (SYS_SYSCALL_BASE + 74);
	SYS_SETRLIMIT			= (SYS_SYSCALL_BASE + 75);
	SYS_GETRLIMIT			= (SYS_SYSCALL_BASE + 76);
	SYS_GETRUSAGE			= (SYS_SYSCALL_BASE + 77);
	SYS_GETTIMEOFDAY		= (SYS_SYSCALL_BASE + 78);
	SYS_SETTIMEOFDAY		= (SYS_SYSCALL_BASE + 79);
	SYS_GETGROUPS			= (SYS_SYSCALL_BASE + 80);
	SYS_SETGROUPS			= (SYS_SYSCALL_BASE + 81);
	SYS_SELECT			= (SYS_SYSCALL_BASE + 82);
	SYS_SYMLINK			= (SYS_SYSCALL_BASE + 83);
	SYS_READLINK			= (SYS_SYSCALL_BASE + 85);
	SYS_USELIB			= (SYS_SYSCALL_BASE + 86);
	SYS_SWAPON			= (SYS_SYSCALL_BASE + 87);
	SYS_REBOOT			= (SYS_SYSCALL_BASE + 88);
	SYS_READDIR			= (SYS_SYSCALL_BASE + 89);
	SYS_MMAP			= (SYS_SYSCALL_BASE + 90);
	SYS_MUNMAP			= (SYS_SYSCALL_BASE + 91);
	SYS_TRUNCATE			= (SYS_SYSCALL_BASE + 92);
	SYS_FTRUNCATE			= (SYS_SYSCALL_BASE + 93);
	SYS_FCHMOD			= (SYS_SYSCALL_BASE + 94);
	SYS_FCHOWN			= (SYS_SYSCALL_BASE + 95);
	SYS_GETPRIORITY			= (SYS_SYSCALL_BASE + 96);
	SYS_SETPRIORITY			= (SYS_SYSCALL_BASE + 97);
	SYS_STATFS			= (SYS_SYSCALL_BASE + 99);
	SYS_FSTATFS			= (SYS_SYSCALL_BASE + 100);
	SYS_SOCKETCALL			= (SYS_SYSCALL_BASE + 102);
	SYS_SYSLOG			= (SYS_SYSCALL_BASE + 103);
	SYS_SETITIMER			= (SYS_SYSCALL_BASE + 104);
	SYS_GETITIMER			= (SYS_SYSCALL_BASE + 105);
	SYS_STAT			= (SYS_SYSCALL_BASE + 106);
	SYS_LSTAT			= (SYS_SYSCALL_BASE + 107);
	SYS_FSTAT			= (SYS_SYSCALL_BASE + 108);
	SYS_VHANGUP			= (SYS_SYSCALL_BASE + 111);
	SYS_SYSCALL			= (SYS_SYSCALL_BASE + 113);
	SYS_WAIT4			= (SYS_SYSCALL_BASE + 114);
	SYS_SWAPOFF			= (SYS_SYSCALL_BASE + 115);
	SYS_SYSINFO			= (SYS_SYSCALL_BASE + 116);
	SYS_IPC				= (SYS_SYSCALL_BASE + 117);
	SYS_FSYNC			= (SYS_SYSCALL_BASE + 118);
	SYS_SIGRETURN			= (SYS_SYSCALL_BASE + 119);
	SYS_CLONE			= (SYS_SYSCALL_BASE + 120);
	SYS_SETDOMAINNAME		= (SYS_SYSCALL_BASE + 121);
	SYS_UNAME			= (SYS_SYSCALL_BASE + 122);
	SYS_ADJTIMEX			= (SYS_SYSCALL_BASE + 124);
	SYS_MPROTECT			= (SYS_SYSCALL_BASE + 125);
	SYS_SIGPROCMASK			= (SYS_SYSCALL_BASE + 126);
	SYS_INIT_MODULE			= (SYS_SYSCALL_BASE + 128);
	SYS_DELETE_MODULE		= (SYS_SYSCALL_BASE + 129);
	SYS_QUOTACTL			= (SYS_SYSCALL_BASE + 131);
	SYS_GETPGID			= (SYS_SYSCALL_BASE + 132);
	SYS_FCHDIR			= (SYS_SYSCALL_BASE + 133);
	SYS_BDFLUSH			= (SYS_SYSCALL_BASE + 134);
	SYS_SYSFS			= (SYS_SYSCALL_BASE + 135);
	SYS_PERSONALITY			= (SYS_SYSCALL_BASE + 136);
	SYS_SETFSUID			= (SYS_SYSCALL_BASE + 138);
	SYS_SETFSGID			= (SYS_SYSCALL_BASE + 139);
	SYS__LLSEEK			= (SYS_SYSCALL_BASE + 140);
	SYS_GETDENTS			= (SYS_SYSCALL_BASE + 141);
	SYS__NEWSELECT			= (SYS_SYSCALL_BASE + 142);
	SYS_FLOCK			= (SYS_SYSCALL_BASE + 143);
	SYS_MSYNC			= (SYS_SYSCALL_BASE + 144);
	SYS_READV			= (SYS_SYSCALL_BASE + 145);
	SYS_WRITEV			= (SYS_SYSCALL_BASE + 146);
	SYS_GETSID			= (SYS_SYSCALL_BASE + 147);
	SYS_FDATASYNC			= (SYS_SYSCALL_BASE + 148);
	SYS__SYSCTL			= (SYS_SYSCALL_BASE + 149);
	SYS_MLOCK			= (SYS_SYSCALL_BASE + 150);
	SYS_MUNLOCK			= (SYS_SYSCALL_BASE + 151);
	SYS_MLOCKALL			= (SYS_SYSCALL_BASE + 152);
	SYS_MUNLOCKALL			= (SYS_SYSCALL_BASE + 153);
	SYS_SCHED_SETPARAM		= (SYS_SYSCALL_BASE + 154);
	SYS_SCHED_GETPARAM		= (SYS_SYSCALL_BASE + 155);
	SYS_SCHED_SETSCHEDULER		= (SYS_SYSCALL_BASE + 156);
	SYS_SCHED_GETSCHEDULER		= (SYS_SYSCALL_BASE + 157);
	SYS_SCHED_YIELD			= (SYS_SYSCALL_BASE + 158);
	SYS_SCHED_GET_PRIORITY_MAX	= (SYS_SYSCALL_BASE + 159);
	SYS_SCHED_GET_PRIORITY_MIN	= (SYS_SYSCALL_BASE + 160);
	SYS_SCHED_RR_GET_INTERVAL	= (SYS_SYSCALL_BASE + 161);
	SYS_NANOSLEEP			= (SYS_SYSCALL_BASE + 162);
	SYS_MREMAP			= (SYS_SYSCALL_BASE + 163);
	SYS_SETRESUID			= (SYS_SYSCALL_BASE + 164);
	SYS_GETRESUID			= (SYS_SYSCALL_BASE + 165);
	SYS_POLL			= (SYS_SYSCALL_BASE + 168);
	SYS_NFSSERVCTL			= (SYS_SYSCALL_BASE + 169);
	SYS_SETRESGID			= (SYS_SYSCALL_BASE + 170);
	SYS_GETRESGID			= (SYS_SYSCALL_BASE + 171);
	SYS_PRCTL			= (SYS_SYSCALL_BASE + 172);
	SYS_RT_SIGRETURN		= (SYS_SYSCALL_BASE + 173);
	SYS_RT_SIGACTION		= (SYS_SYSCALL_BASE + 174);
	SYS_RT_SIGPROCMASK		= (SYS_SYSCALL_BASE + 175);
	SYS_RT_SIGPENDING		= (SYS_SYSCALL_BASE + 176);
	SYS_RT_SIGTIMEDWAIT		= (SYS_SYSCALL_BASE + 177);
	SYS_RT_SIGQUEUEINFO		= (SYS_SYSCALL_BASE + 178);
	SYS_RT_SIGSUSPEND		= (SYS_SYSCALL_BASE + 179);
	SYS_PREAD64			= (SYS_SYSCALL_BASE + 180);
	SYS_PWRITE64			= (SYS_SYSCALL_BASE + 181);
	SYS_CHOWN			= (SYS_SYSCALL_BASE + 182);
	SYS_GETCWD			= (SYS_SYSCALL_BASE + 183);
	SYS_CAPGET			= (SYS_SYSCALL_BASE + 184);
	SYS_CAPSET			= (SYS_SYSCALL_BASE + 185);
	SYS_SIGALTSTACK			= (SYS_SYSCALL_BASE + 186);
	SYS_SENDFILE			= (SYS_SYSCALL_BASE + 187);
	SYS_VFORK			= (SYS_SYSCALL_BASE + 190);
	SYS_UGETRLIMIT			= (SYS_SYSCALL_BASE + 191);
	SYS_MMAP2			= (SYS_SYSCALL_BASE + 192);
	SYS_TRUNCATE64			= (SYS_SYSCALL_BASE + 193);
	SYS_FTRUNCATE64			= (SYS_SYSCALL_BASE + 194);
	SYS_STAT64			= (SYS_SYSCALL_BASE + 195);
	SYS_LSTAT64			= (SYS_SYSCALL_BASE + 196);
	SYS_FSTAT64			= (SYS_SYSCALL_BASE + 197);
	SYS_LCHOWN32			= (SYS_SYSCALL_BASE + 198);
	SYS_GETUID32			= (SYS_SYSCALL_BASE + 199);
	SYS_GETGID32			= (SYS_SYSCALL_BASE + 200);
	SYS_GETEUID32			= (SYS_SYSCALL_BASE + 201);
	SYS_GETEGID32			= (SYS_SYSCALL_BASE + 202);
	SYS_SETREUID32			= (SYS_SYSCALL_BASE + 203);
	SYS_SETREGID32			= (SYS_SYSCALL_BASE + 204);
	SYS_GETGROUPS32			= (SYS_SYSCALL_BASE + 205);
	SYS_SETGROUPS32			= (SYS_SYSCALL_BASE + 206);
	SYS_FCHOWN32			= (SYS_SYSCALL_BASE + 207);
	SYS_SETRESUID32			= (SYS_SYSCALL_BASE + 208);
	SYS_GETRESUID32			= (SYS_SYSCALL_BASE + 209);
	SYS_SETRESGID32			= (SYS_SYSCALL_BASE + 210);
	SYS_GETRESGID32			= (SYS_SYSCALL_BASE + 211);
	SYS_CHOWN32			= (SYS_SYSCALL_BASE + 212);
	SYS_SETUID32			= (SYS_SYSCALL_BASE + 213);
	SYS_SETGID32			= (SYS_SYSCALL_BASE + 214);
	SYS_SETFSUID32			= (SYS_SYSCALL_BASE + 215);
	SYS_SETFSGID32			= (SYS_SYSCALL_BASE + 216);
	SYS_GETDENTS64			= (SYS_SYSCALL_BASE + 217);
	SYS_PIVOT_ROOT			= (SYS_SYSCALL_BASE + 218);
	SYS_MINCORE			= (SYS_SYSCALL_BASE + 219);
	SYS_MADVISE			= (SYS_SYSCALL_BASE + 220);
	SYS_FCNTL64			= (SYS_SYSCALL_BASE + 221);
	SYS_GETTID			= (SYS_SYSCALL_BASE + 224);
	SYS_READAHEAD			= (SYS_SYSCALL_BASE + 225);
	SYS_SETXATTR			= (SYS_SYSCALL_BASE + 226);
	SYS_LSETXATTR			= (SYS_SYSCALL_BASE + 227);
	SYS_FSETXATTR			= (SYS_SYSCALL_BASE + 228);
	SYS_GETXATTR			= (SYS_SYSCALL_BASE + 229);
	SYS_LGETXATTR			= (SYS_SYSCALL_BASE + 230);
	SYS_FGETXATTR			= (SYS_SYSCALL_BASE + 231);
	SYS_LISTXATTR			= (SYS_SYSCALL_BASE + 232);
	SYS_LLISTXATTR			= (SYS_SYSCALL_BASE + 233);
	SYS_FLISTXATTR			= (SYS_SYSCALL_BASE + 234);
	SYS_REMOVEXATTR			= (SYS_SYSCALL_BASE + 235);
	SYS_LREMOVEXATTR		= (SYS_SYSCALL_BASE + 236);
	SYS_FREMOVEXATTR		= (SYS_SYSCALL_BASE + 237);
	SYS_TKILL			= (SYS_SYSCALL_BASE + 238);
	SYS_SENDFILE64			= (SYS_SYSCALL_BASE + 239);
	SYS_FUTEX			= (SYS_SYSCALL_BASE + 240);
	SYS_SCHED_SETAFFINITY		= (SYS_SYSCALL_BASE + 241);
	SYS_SCHED_GETAFFINITY		= (SYS_SYSCALL_BASE + 242);
	SYS_IO_SETUP			= (SYS_SYSCALL_BASE + 243);
	SYS_IO_DESTROY			= (SYS_SYSCALL_BASE + 244);
	SYS_IO_GETEVENTS		= (SYS_SYSCALL_BASE + 245);
	SYS_IO_SUBMIT			= (SYS_SYSCALL_BASE + 246);
	SYS_IO_CANCEL			= (SYS_SYSCALL_BASE + 247);
	SYS_EXIT_GROUP			= (SYS_SYSCALL_BASE + 248);
	SYS_LOOKUP_DCOOKIE		= (SYS_SYSCALL_BASE + 249);
	SYS_EPOLL_CREATE		= (SYS_SYSCALL_BASE + 250);
	SYS_EPOLL_CTL			= (SYS_SYSCALL_BASE + 251);
	SYS_EPOLL_WAIT			= (SYS_SYSCALL_BASE + 252);
	SYS_REMAP_FILE_PAGES		= (SYS_SYSCALL_BASE + 253);
	SYS_SET_TID_ADDRESS		= (SYS_SYSCALL_BASE + 256);
	SYS_TIMER_CREATE		= (SYS_SYSCALL_BASE + 257);
	SYS_TIMER_SETTIME		= (SYS_SYSCALL_BASE + 258);
	SYS_TIMER_GETTIME		= (SYS_SYSCALL_BASE + 259);
	SYS_TIMER_GETOVERRUN		= (SYS_SYSCALL_BASE + 260);
	SYS_TIMER_DELETE		= (SYS_SYSCALL_BASE + 261);
	SYS_CLOCK_SETTIME		= (SYS_SYSCALL_BASE + 262);
	SYS_CLOCK_GETTIME		= (SYS_SYSCALL_BASE + 263);
	SYS_CLOCK_GETRES		= (SYS_SYSCALL_BASE + 264);
	SYS_CLOCK_NANOSLEEP		= (SYS_SYSCALL_BASE + 265);
	SYS_STATFS64			= (SYS_SYSCALL_BASE + 266);
	SYS_FSTATFS64			= (SYS_SYSCALL_BASE + 267);
	SYS_TGKILL			= (SYS_SYSCALL_BASE + 268);
	SYS_UTIMES			= (SYS_SYSCALL_BASE + 269);
	SYS_ARM_FADVISE64_64		= (SYS_SYSCALL_BASE + 270);
	SYS_PCICONFIG_IOBASE		= (SYS_SYSCALL_BASE + 271);
	SYS_PCICONFIG_READ		= (SYS_SYSCALL_BASE + 272);
	SYS_PCICONFIG_WRITE		= (SYS_SYSCALL_BASE + 273);
	SYS_MQ_OPEN			= (SYS_SYSCALL_BASE + 274);
	SYS_MQ_UNLINK			= (SYS_SYSCALL_BASE + 275);
	SYS_MQ_TIMEDSEND		= (SYS_SYSCALL_BASE + 276);
	SYS_MQ_TIMEDRECEIVE		= (SYS_SYSCALL_BASE + 277);
	SYS_MQ_NOTIFY			= (SYS_SYSCALL_BASE + 278);
	SYS_MQ_GETSETATTR		= (SYS_SYSCALL_BASE + 279);
	SYS_WAITID			= (SYS_SYSCALL_BASE + 280);
	SYS_SOCKET			= (SYS_SYSCALL_BASE + 281);
	SYS_BIND			= (SYS_SYSCALL_BASE + 282);
	SYS_CONNECT			= (SYS_SYSCALL_BASE + 283);
	SYS_LISTEN			= (SYS_SYSCALL_BASE + 284);
	SYS_ACCEPT			= (SYS_SYSCALL_BASE + 285);
	SYS_GETSOCKNAME			= (SYS_SYSCALL_BASE + 286);
	SYS_GETPEERNAME			= (SYS_SYSCALL_BASE + 287);
	SYS_SOCKETPAIR			= (SYS_SYSCALL_BASE + 288);
	SYS_SEND			= (SYS_SYSCALL_BASE + 289);
	SYS_SENDTO			= (SYS_SYSCALL_BASE + 290);
	SYS_RECV			= (SYS_SYSCALL_BASE + 291);
	SYS_RECVFROM			= (SYS_SYSCALL_BASE + 292);
	SYS_SHUTDOWN			= (SYS_SYSCALL_BASE + 293);
	SYS_SETSOCKOPT			= (SYS_SYSCALL_BASE + 294);
	SYS_GETSOCKOPT			= (SYS_SYSCALL_BASE + 295);
	SYS_SENDMSG			= (SYS_SYSCALL_BASE + 296);
	SYS_RECVMSG			= (SYS_SYSCALL_BASE + 297);
	SYS_SEMOP			= (SYS_SYSCALL_BASE + 298);
	SYS_SEMGET			= (SYS_SYSCALL_BASE + 299);
	SYS_SEMCTL			= (SYS_SYSCALL_BASE + 300);
	SYS_MSGSND			= (SYS_SYSCALL_BASE + 301);
	SYS_MSGRCV			= (SYS_SYSCALL_BASE + 302);
	SYS_MSGGET			= (SYS_SYSCALL_BASE + 303);
	SYS_MSGCTL			= (SYS_SYSCALL_BASE + 304);
	SYS_SHMAT			= (SYS_SYSCALL_BASE + 305);
	SYS_SHMDT			= (SYS_SYSCALL_BASE + 306);
	SYS_SHMGET			= (SYS_SYSCALL_BASE + 307);
	SYS_SHMCTL			= (SYS_SYSCALL_BASE + 308);
	SYS_ADD_KEY			= (SYS_SYSCALL_BASE + 309);
	SYS_REQUEST_KEY			= (SYS_SYSCALL_BASE + 310);
	SYS_KEYCTL			= (SYS_SYSCALL_BASE + 311);
	SYS_SEMTIMEDOP			= (SYS_SYSCALL_BASE + 312);
	SYS_VSERVER			= (SYS_SYSCALL_BASE + 313);
	SYS_IOPRIO_SET			= (SYS_SYSCALL_BASE + 314);
	SYS_IOPRIO_GET			= (SYS_SYSCALL_BASE + 315);
	SYS_INOTIFY_INIT		= (SYS_SYSCALL_BASE + 316);
	SYS_INOTIFY_ADD_WATCH		= (SYS_SYSCALL_BASE + 317);
	SYS_INOTIFY_RM_WATCH		= (SYS_SYSCALL_BASE + 318);
	SYS_MBIND			= (SYS_SYSCALL_BASE + 319);
	SYS_GET_MEMPOLICY		= (SYS_SYSCALL_BASE + 320);
	SYS_SET_MEMPOLICY		= (SYS_SYSCALL_BASE + 321);
	SYS_OPENAT			= (SYS_SYSCALL_BASE + 322);
	SYS_MKDIRAT			= (SYS_SYSCALL_BASE + 323);
	SYS_MKNODAT			= (SYS_SYSCALL_BASE + 324);
	SYS_FCHOWNAT			= (SYS_SYSCALL_BASE + 325);
	SYS_FUTIMESAT			= (SYS_SYSCALL_BASE + 326);
	SYS_FSTATAT64			= (SYS_SYSCALL_BASE + 327);
	SYS_UNLINKAT			= (SYS_SYSCALL_BASE + 328);
	SYS_RENAMEAT			= (SYS_SYSCALL_BASE + 329);
	SYS_LINKAT			= (SYS_SYSCALL_BASE + 330);
	SYS_SYMLINKAT			= (SYS_SYSCALL_BASE + 331);
	SYS_READLINKAT			= (SYS_SYSCALL_BASE + 332);
	SYS_FCHMODAT			= (SYS_SYSCALL_BASE + 333);
	SYS_FACCESSAT			= (SYS_SYSCALL_BASE + 334);
	SYS_UNSHARE			= (SYS_SYSCALL_BASE + 337);
	SYS_SET_ROBUST_LIST		= (SYS_SYSCALL_BASE + 338);
	SYS_GET_ROBUST_LIST		= (SYS_SYSCALL_BASE + 339);
	SYS_SPLICE			= (SYS_SYSCALL_BASE + 340);
	SYS_ARM_SYNC_FILE_RANGE		= (SYS_SYSCALL_BASE + 341);
	SYS_SYNC_FILE_RANGE2		= SYS_ARM_SYNC_FILE_RANGE;
	SYS_TEE				= (SYS_SYSCALL_BASE + 342);
	SYS_VMSPLICE			= (SYS_SYSCALL_BASE + 343);
	SYS_MOVE_PAGES			= (SYS_SYSCALL_BASE + 344);
	SYS_GETCPU			= (SYS_SYSCALL_BASE + 345);
	SYS_KEXEC_LOAD			= (SYS_SYSCALL_BASE + 347);
	SYS_UTIMENSAT			= (SYS_SYSCALL_BASE + 348);
	SYS_SIGNALFD			= (SYS_SYSCALL_BASE + 349);
	SYS_TIMERFD_CREATE		= (SYS_SYSCALL_BASE + 350);
	SYS_EVENTFD			= (SYS_SYSCALL_BASE + 351);
	SYS_FALLOCATE			= (SYS_SYSCALL_BASE + 352);
	SYS_TIMERFD_SETTIME		= (SYS_SYSCALL_BASE + 353);
	SYS_TIMERFD_GETTIME		= (SYS_SYSCALL_BASE + 354);
	SYS_SIGNALFD4			= (SYS_SYSCALL_BASE + 355);
	SYS_EVENTFD2			= (SYS_SYSCALL_BASE + 356);
	SYS_EPOLL_CREATE1		= (SYS_SYSCALL_BASE + 357);
	SYS_DUP3			= (SYS_SYSCALL_BASE + 358);
	SYS_PIPE2			= (SYS_SYSCALL_BASE + 359);
	SYS_INOTIFY_INIT1		= (SYS_SYSCALL_BASE + 360);
)

func _darwin_system_call_conflict()	{}
