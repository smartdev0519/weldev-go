// mkerrors_nacl.sh /home/rsc/pub/nacl/native_client/src/trusted/service_runtime/include/sys/errno.h
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

const (
	EPERM           = 1
	ENOENT          = 2
	ESRCH           = 3
	EINTR           = 4
	EIO             = 5
	ENXIO           = 6
	E2BIG           = 7
	ENOEXEC         = 8
	EBADF           = 9
	ECHILD          = 10
	EAGAIN          = 11
	ENOMEM          = 12
	EACCES          = 13
	EFAULT          = 14
	EBUSY           = 16
	EEXIST          = 17
	EXDEV           = 18
	ENODEV          = 19
	ENOTDIR         = 20
	EISDIR          = 21
	EINVAL          = 22
	ENFILE          = 23
	EMFILE          = 24
	ENOTTY          = 25
	EFBIG           = 27
	ENOSPC          = 28
	ESPIPE          = 29
	EROFS           = 30
	EMLINK          = 31
	EPIPE           = 32
	ENAMETOOLONG    = 36
	ENOSYS          = 38
	EDQUOT          = 122
	EDOM            = 33
	ERANGE          = 34
	ENOMSG          = 35
	ECHRNG          = 37
	EL3HLT          = 39
	EL3RST          = 40
	ELNRNG          = 41
	EUNATCH         = 42
	ENOCSI          = 43
	EL2HLT          = 44
	EDEADLK         = 45
	ENOLCK          = 46
	EBADE           = 50
	EBADR           = 51
	EXFULL          = 52
	ENOANO          = 53
	EBADRQC         = 54
	EBADSLT         = 55
	EBFONT          = 57
	ENOSTR          = 60
	ENODATA         = 61
	ETIME           = 62
	ENOSR           = 63
	ENONET          = 64
	ENOPKG          = 65
	EREMOTE         = 66
	ENOLINK         = 67
	EADV            = 68
	ESRMNT          = 69
	ECOMM           = 70
	EPROTO          = 71
	EMULTIHOP       = 74
	ELBIN           = 75
	EDOTDOT         = 76
	EBADMSG         = 77
	EFTYPE          = 79
	ENOTUNIQ        = 80
	EBADFD          = 81
	EREMCHG         = 82
	ELIBACC         = 83
	ELIBBAD         = 84
	ELIBSCN         = 85
	ELIBMAX         = 86
	ELIBEXEC        = 87
	ENMFILE         = 89
	ENOTEMPTY       = 90
	ELOOP           = 92
	EOPNOTSUPP      = 95
	EPFNOSUPPORT    = 96
	ECONNRESET      = 104
	ENOBUFS         = 105
	EAFNOSUPPORT    = 106
	EPROTOTYPE      = 107
	ENOTSOCK        = 108
	ENOPROTOOPT     = 109
	ESHUTDOWN       = 110
	ECONNREFUSED    = 111
	EADDRINUSE      = 112
	ECONNABORTED    = 113
	ENETUNREACH     = 114
	ENETDOWN        = 115
	ETIMEDOUT       = 116
	EHOSTDOWN       = 117
	EHOSTUNREACH    = 118
	EINPROGRESS     = 119
	EALREADY        = 120
	EDESTADDRREQ    = 121
	EPROTONOSUPPORT = 123
	ESOCKTNOSUPPORT = 124
	EADDRNOTAVAIL   = 125
	ENETRESET       = 126
	EISCONN         = 127
	ENOTCONN        = 128
	ETOOMANYREFS    = 129
	EPROCLIM        = 130
	EUSERS          = 131
	ESTALE          = 133
	ENOMEDIUM       = 135
	ENOSHARE        = 136
	ECASECLASH      = 137
	EILSEQ          = 138
	EOVERFLOW       = 139
	ECANCELED       = 140
	EL2NSYNC        = 88
	EIDRM           = 91
	EMSGSIZE        = 132
	ENACL           = 99 /* otherwise unused */
)


// Error table
var errors = [...]string{
	EPERM:           "operation not permitted",
	ENOENT:          "no such file or directory",
	ESRCH:           "no such process",
	EINTR:           "interrupted system call",
	EIO:             "I/O error",
	ENXIO:           "no such device or address",
	E2BIG:           "argument list too long",
	ENOEXEC:         "exec format error",
	EBADF:           "bad file number",
	ECHILD:          "no child processes",
	EAGAIN:          "try again",
	ENOMEM:          "out of memory",
	EACCES:          "permission denied",
	EFAULT:          "bad address",
	EBUSY:           "device or resource busy",
	EEXIST:          "file exists",
	EXDEV:           "cross-device link",
	ENODEV:          "no such device",
	ENOTDIR:         "not a directory",
	EISDIR:          "is a directory",
	EINVAL:          "invalid argument",
	ENFILE:          "file table overflow",
	EMFILE:          "too many open files",
	ENOTTY:          "not a typewriter",
	EFBIG:           "file too large",
	ENOSPC:          "no space left on device",
	ESPIPE:          "illegal seek",
	EROFS:           "read-only file system",
	EMLINK:          "too many links",
	EPIPE:           "broken pipe",
	ENAMETOOLONG:    "file name too long",
	ENOSYS:          "function not implemented",
	EDQUOT:          "quota exceeded",
	EDOM:            "math arg out of domain of func",
	ERANGE:          "math result not representable",
	ENOMSG:          "no message of desired type",
	ECHRNG:          "channel number out of range",
	EL3HLT:          "level 3 halted",
	EL3RST:          "level 3 reset",
	ELNRNG:          "link number out of range",
	EUNATCH:         "protocol driver not attached",
	ENOCSI:          "no CSI structure available",
	EL2HLT:          "level 2 halted",
	EDEADLK:         "deadlock condition",
	ENOLCK:          "no record locks available",
	EBADE:           "invalid exchange",
	EBADR:           "invalid request descriptor",
	EXFULL:          "exchange full",
	ENOANO:          "no anode",
	EBADRQC:         "invalid request code",
	EBADSLT:         "invalid slot",
	EBFONT:          "bad font file fmt",
	ENOSTR:          "device not a stream",
	ENODATA:         "no data (for no delay io)",
	ETIME:           "timer expired",
	ENOSR:           "out of streams resources",
	ENONET:          "machine is not on the network",
	ENOPKG:          "package not installed",
	EREMOTE:         "the object is remote",
	ENOLINK:         "the link has been severed",
	EADV:            "advertise error",
	ESRMNT:          "srmount error",
	ECOMM:           "communication error on send",
	EPROTO:          "protocol error",
	EMULTIHOP:       "multihop attempted",
	ELBIN:           "inode is remote (not really error)",
	EDOTDOT:         "cross mount point (not really error)",
	EBADMSG:         "trying to read unreadable message",
	EFTYPE:          "inappropriate file type or format",
	ENOTUNIQ:        "given log. name not unique",
	EBADFD:          "f.d. invalid for this operation",
	EREMCHG:         "remote address changed",
	ELIBACC:         "can't access a needed shared lib",
	ELIBBAD:         "accessing a corrupted shared lib",
	ELIBSCN:         ".lib section in a.out corrupted",
	ELIBMAX:         "attempting to link in too many libs",
	ELIBEXEC:        "attempting to exec a shared library",
	ENMFILE:         "no more files",
	ENOTEMPTY:       "directory not empty",
	ELOOP:           "too many symbolic links",
	EOPNOTSUPP:      "operation not supported on transport endpoint",
	EPFNOSUPPORT:    "protocol family not supported",
	ECONNRESET:      "connection reset by peer",
	ENOBUFS:         "no buffer space available",
	EAFNOSUPPORT:    "address family not supported by protocol family",
	EPROTOTYPE:      "protocol wrong type for socket",
	ENOTSOCK:        "socket operation on non-socket",
	ENOPROTOOPT:     "protocol not available",
	ESHUTDOWN:       "can't send after socket shutdown",
	ECONNREFUSED:    "connection refused",
	EADDRINUSE:      "address already in use",
	ECONNABORTED:    "connection aborted",
	ENETUNREACH:     "network is unreachable",
	ENETDOWN:        "network interface is not configured",
	ETIMEDOUT:       "connection timed out",
	EHOSTDOWN:       "host is down",
	EHOSTUNREACH:    "host is unreachable",
	EINPROGRESS:     "connection already in progress",
	EALREADY:        "socket already connected",
	EDESTADDRREQ:    "destination address required",
	EPROTONOSUPPORT: "unknown protocol",
	ESOCKTNOSUPPORT: "socket type not supported",
	EADDRNOTAVAIL:   "address not available",
	EISCONN:         "socket is already connected",
	ENOTCONN:        "socket is not connected",
	ENOMEDIUM:       "no medium (in tape drive)",
	ENOSHARE:        "no such host or network path",
	ECASECLASH:      "filename exists with different case",
	EOVERFLOW:       "value too large for defined data type",
	ECANCELED:       "operation canceled.",
	EL2NSYNC:        "level 2 not synchronized",
	EIDRM:           "identifier removed",
	EMSGSIZE:        "message too long",
	ENACL:           "not supported by native client",
}
