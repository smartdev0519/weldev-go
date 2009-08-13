// mkerrors.sh
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// godefs -gsyscall _errors.c

// MACHINE GENERATED - DO NOT EDIT.

package syscall

// Constants
const (
	EMULTIHOP = 0x5f;
	EAFNOSUPPORT = 0x2f;
	EACCES = 0xd;
	EDESTADDRREQ = 0x27;
	EILSEQ = 0x5c;
	ESPIPE = 0x1d;
	EMLINK = 0x1f;
	EPROGUNAVAIL = 0x4a;
	ENOTTY = 0x19;
	EBADF = 0x9;
	ERANGE = 0x22;
	ECANCELED = 0x59;
	ETXTBSY = 0x1a;
	ENOMEM = 0xc;
	EINPROGRESS = 0x24;
	ENOTEMPTY = 0x42;
	ENOTBLK = 0xf;
	EPROTOTYPE = 0x29;
	ENOMSG = 0x5b;
	ERPCMISMATCH = 0x49;
	ENOTDIR = 0x14;
	EALREADY = 0x25;
	ETIMEDOUT = 0x3c;
	ENEEDAUTH = 0x51;
	ENODATA = 0x60;
	EINTR = 0x4;
	ENOLINK = 0x61;
	EPERM = 0x1;
	ENETDOWN = 0x32;
	ESTALE = 0x46;
	ENOTSOCK = 0x26;
	ENOSR = 0x62;
	EAUTH = 0x50;
	ECHILD = 0xa;
	EPIPE = 0x20;
	ENOATTR = 0x5d;
	EBADMSG = 0x5e;
	EREMOTE = 0x47;
	ETOOMANYREFS = 0x3b;
	EPFNOSUPPORT = 0x2e;
	EPROCUNAVAIL = 0x4c;
	EADDRINUSE = 0x30;
	ENETRESET = 0x34;
	EISDIR = 0x15;
	EIDRM = 0x5a;
	EDEVERR = 0x53;
	EINVAL = 0x16;
	ESHUTDOWN = 0x3a;
	EPWROFF = 0x52;
	EOVERFLOW = 0x54;
	EBUSY = 0x10;
	EPROCLIM = 0x43;
	EPROTO = 0x64;
	ENODEV = 0x13;
	EROFS = 0x1e;
	E2BIG = 0x7;
	EDEADLK = 0xb;
	ECONNRESET = 0x36;
	EBADMACHO = 0x58;
	ENXIO = 0x6;
	EBADRPC = 0x48;
	ENAMETOOLONG = 0x3f;
	ELAST = 0x67;
	ESOCKTNOSUPPORT = 0x2c;
	EADDRNOTAVAIL = 0x31;
	ETIME = 0x65;
	EPROTONOSUPPORT = 0x2b;
	EIO = 0x5;
	ENETUNREACH = 0x33;
	EXDEV = 0x12;
	EDQUOT = 0x45;
	ENOSPC = 0x1c;
	ENOEXEC = 0x8;
	EMSGSIZE = 0x28;
	EFTYPE = 0x4f;
	EDOM = 0x21;
	ENOSTR = 0x63;
	EFBIG = 0x1b;
	ESRCH = 0x3;
	EHOSTDOWN = 0x40;
	ENOLCK = 0x4d;
	ENFILE = 0x17;
	ENOSYS = 0x4e;
	EBADARCH = 0x56;
	ENOTCONN = 0x39;
	ENOTSUP = 0x2d;
	ECONNABORTED = 0x35;
	EISCONN = 0x38;
	ESHLIBVERS = 0x57;
	EUSERS = 0x44;
	ENOPROTOOPT = 0x2a;
	EMFILE = 0x18;
	ELOOP = 0x3e;
	ENOBUFS = 0x37;
	EFAULT = 0xe;
	EWOULDBLOCK = 0x23;
	EBADEXEC = 0x55;
	ENOPOLICY = 0x67;
	ECONNREFUSED = 0x3d;
	EAGAIN = 0x23;
	EEXIST = 0x11;
	EPROGMISMATCH = 0x4b;
	ENOENT = 0x2;
	EHOSTUNREACH = 0x41;
	EOPNOTSUPP = 0x66;
	SIGBUS = 0xa;
	SIGTTIN = 0x15;
	SIGPROF = 0x1b;
	SIGFPE = 0x8;
	SIGHUP = 0x1;
	SIGTTOU = 0x16;
	SIGUSR1 = 0x1e;
	SIGURG = 0x10;
	SIGQUIT = 0x3;
	SIGIO = 0x17;
	SIGABRT = 0x6;
	SIGINFO = 0x1d;
	SIGUSR2 = 0x1f;
	SIGTRAP = 0x5;
	SIGVTALRM = 0x1a;
	SIGSEGV = 0xb;
	SIGCONT = 0x13;
	SIGPIPE = 0xd;
	SIGXFSZ = 0x19;
	SIGCHLD = 0x14;
	SIGSYS = 0xc;
	SIGSTOP = 0x11;
	SIGALRM = 0xe;
	SIGTSTP = 0x12;
	SIGEMT = 0x7;
	SIGKILL = 0x9;
	SIGXCPU = 0x18;
	SIGILL = 0x4;
	SIGINT = 0x2;
	SIGIOT = 0x6;
	SIGTERM = 0xf;
	SIGWINCH = 0x1c;
)

// Types


// Error table
var errors = [...]string {
	95: "EMULTIHOP (Reserved)",
	47: "address family not supported by protocol family",
	13: "permission denied",
	39: "destination address required",
	92: "illegal byte sequence",
	29: "illegal seek",
	31: "too many links",
	74: "RPC prog. not avail",
	25: "inappropriate ioctl for device",
	9: "bad file descriptor",
	34: "result too large",
	89: "operation canceled",
	26: "text file busy",
	12: "cannot allocate memory",
	36: "operation now in progress",
	66: "directory not empty",
	15: "block device required",
	41: "protocol wrong type for socket",
	91: "no message of desired type",
	73: "RPC version wrong",
	20: "not a directory",
	37: "operation already in progress",
	60: "operation timed out",
	81: "need authenticator",
	96: "no message available on STREAM",
	4: "interrupted system call",
	97: "ENOLINK (Reserved)",
	1: "operation not permitted",
	50: "network is down",
	70: "stale NFS file handle",
	38: "socket operation on non-socket",
	98: "no STREAM resources",
	80: "authentication error",
	10: "no child processes",
	32: "broken pipe",
	93: "attribute not found",
	94: "bad message",
	71: "too many levels of remote in path",
	59: "too many references: can't splice",
	46: "protocol family not supported",
	76: "bad procedure for program",
	48: "address already in use",
	52: "network dropped connection on reset",
	21: "is a directory",
	90: "identifier removed",
	83: "device error",
	22: "invalid argument",
	58: "can't send after socket shutdown",
	82: "device power is off",
	84: "value too large to be stored in data type",
	16: "resource busy",
	67: "too many processes",
	100: "protocol error",
	19: "operation not supported by device",
	30: "read-only file system",
	7: "argument list too long",
	11: "resource deadlock avoided",
	54: "connection reset by peer",
	88: "malformed Mach-o file",
	6: "device not configured",
	72: "RPC struct is bad",
	63: "file name too long",
	103: "policy not found",
	44: "socket type not supported",
	49: "can't assign requested address",
	101: "STREAM ioctl timeout",
	43: "protocol not supported",
	5: "input/output error",
	51: "network is unreachable",
	18: "cross-device link",
	69: "disc quota exceeded",
	28: "no space left on device",
	8: "exec format error",
	40: "message too long",
	79: "inappropriate file type or format",
	33: "numerical argument out of domain",
	99: "not a STREAM",
	27: "file too large",
	3: "no such process",
	64: "host is down",
	77: "no locks available",
	23: "too many open files in system",
	78: "function not implemented",
	86: "bad CPU type in executable",
	57: "socket is not connected",
	45: "operation not supported",
	53: "software caused connection abort",
	56: "socket is already connected",
	87: "shared library version mismatch",
	68: "too many users",
	42: "protocol not available",
	24: "too many open files",
	62: "too many levels of symbolic links",
	55: "no buffer space available",
	14: "bad address",
	35: "resource temporarily unavailable",
	85: "bad executable (or shared library)",
	61: "connection refused",
	17: "file exists",
	75: "program version wrong",
	2: "no such file or directory",
	65: "no route to host",
	102: "operation not supported on socket",
}

