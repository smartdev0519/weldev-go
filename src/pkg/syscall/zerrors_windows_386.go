// mkerrors_windows.sh -f -m32
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

// Go names for Windows errors.
const (
	ENOENT  = ERROR_FILE_NOT_FOUND
	ENOTDIR = ERROR_DIRECTORY
)

// Windows reserves errors >= 1<<29 for application use.
const APPLICATION_ERROR = 1 << 29

// Invented values to support what package os and others expects.
const (
	E2BIG = APPLICATION_ERROR + iota
	EACCES
	EADDRINUSE
	EADDRNOTAVAIL
	EADV
	EAFNOSUPPORT
	EAGAIN
	EALREADY
	EBADE
	EBADF
	EBADFD
	EBADMSG
	EBADR
	EBADRQC
	EBADSLT
	EBFONT
	EBUSY
	ECANCELED
	ECHILD
	ECHRNG
	ECOMM
	ECONNABORTED
	ECONNREFUSED
	ECONNRESET
	EDEADLK
	EDEADLOCK
	EDESTADDRREQ
	EDOM
	EDOTDOT
	EDQUOT
	EEXIST
	EFAULT
	EFBIG
	EHOSTDOWN
	EHOSTUNREACH
	EIDRM
	EILSEQ
	EINPROGRESS
	EINTR
	EINVAL
	EIO
	EISCONN
	EISDIR
	EISNAM
	EKEYEXPIRED
	EKEYREJECTED
	EKEYREVOKED
	EL2HLT
	EL2NSYNC
	EL3HLT
	EL3RST
	ELIBACC
	ELIBBAD
	ELIBEXEC
	ELIBMAX
	ELIBSCN
	ELNRNG
	ELOOP
	EMEDIUMTYPE
	EMFILE
	EMLINK
	EMSGSIZE
	EMULTIHOP
	ENAMETOOLONG
	ENAVAIL
	ENETDOWN
	ENETRESET
	ENETUNREACH
	ENFILE
	ENOANO
	ENOBUFS
	ENOCSI
	ENODATA
	ENODEV
	ENOEXEC
	ENOKEY
	ENOLCK
	ENOLINK
	ENOMEDIUM
	ENOMEM
	ENOMSG
	ENONET
	ENOPKG
	ENOPROTOOPT
	ENOSPC
	ENOSR
	ENOSTR
	ENOSYS
	ENOTBLK
	ENOTCONN
	ENOTEMPTY
	ENOTNAM
	ENOTRECOVERABLE
	ENOTSOCK
	ENOTSUP
	ENOTTY
	ENOTUNIQ
	ENXIO
	EOPNOTSUPP
	EOVERFLOW
	EOWNERDEAD
	EPERM
	EPFNOSUPPORT
	EPIPE
	EPROTO
	EPROTONOSUPPORT
	EPROTOTYPE
	ERANGE
	EREMCHG
	EREMOTE
	EREMOTEIO
	ERESTART
	EROFS
	ESHUTDOWN
	ESOCKTNOSUPPORT
	ESPIPE
	ESRCH
	ESRMNT
	ESTALE
	ESTRPIPE
	ETIME
	ETIMEDOUT
	ETOOMANYREFS
	ETXTBSY
	EUCLEAN
	EUNATCH
	EUSERS
	EWOULDBLOCK
	EXDEV
	EXFULL
	EWINDOWS
)

// Error strings for invented errors
var errors = [...]string{
	E2BIG - APPLICATION_ERROR:           "argument list too long",
	EACCES - APPLICATION_ERROR:          "permission denied",
	EADDRINUSE - APPLICATION_ERROR:      "address already in use",
	EADDRNOTAVAIL - APPLICATION_ERROR:   "cannot assign requested address",
	EADV - APPLICATION_ERROR:            "advertise error",
	EAFNOSUPPORT - APPLICATION_ERROR:    "address family not supported by protocol",
	EAGAIN - APPLICATION_ERROR:          "resource temporarily unavailable",
	EALREADY - APPLICATION_ERROR:        "operation already in progress",
	EBADE - APPLICATION_ERROR:           "invalid exchange",
	EBADF - APPLICATION_ERROR:           "bad file descriptor",
	EBADFD - APPLICATION_ERROR:          "file descriptor in bad state",
	EBADMSG - APPLICATION_ERROR:         "bad message",
	EBADR - APPLICATION_ERROR:           "invalid request descriptor",
	EBADRQC - APPLICATION_ERROR:         "invalid request code",
	EBADSLT - APPLICATION_ERROR:         "invalid slot",
	EBFONT - APPLICATION_ERROR:          "bad font file format",
	EBUSY - APPLICATION_ERROR:           "device or resource busy",
	ECANCELED - APPLICATION_ERROR:       "operation canceled",
	ECHILD - APPLICATION_ERROR:          "no child processes",
	ECHRNG - APPLICATION_ERROR:          "channel number out of range",
	ECOMM - APPLICATION_ERROR:           "communication error on send",
	ECONNABORTED - APPLICATION_ERROR:    "software caused connection abort",
	ECONNREFUSED - APPLICATION_ERROR:    "connection refused",
	ECONNRESET - APPLICATION_ERROR:      "connection reset by peer",
	EDEADLK - APPLICATION_ERROR:         "resource deadlock avoided",
	EDEADLOCK - APPLICATION_ERROR:       "resource deadlock avoided",
	EDESTADDRREQ - APPLICATION_ERROR:    "destination address required",
	EDOM - APPLICATION_ERROR:            "numerical argument out of domain",
	EDOTDOT - APPLICATION_ERROR:         "RFS specific error",
	EDQUOT - APPLICATION_ERROR:          "disk quota exceeded",
	EEXIST - APPLICATION_ERROR:          "file exists",
	EFAULT - APPLICATION_ERROR:          "bad address",
	EFBIG - APPLICATION_ERROR:           "file too large",
	EHOSTDOWN - APPLICATION_ERROR:       "host is down",
	EHOSTUNREACH - APPLICATION_ERROR:    "no route to host",
	EIDRM - APPLICATION_ERROR:           "identifier removed",
	EILSEQ - APPLICATION_ERROR:          "invalid or incomplete multibyte or wide character",
	EINPROGRESS - APPLICATION_ERROR:     "operation now in progress",
	EINTR - APPLICATION_ERROR:           "interrupted system call",
	EINVAL - APPLICATION_ERROR:          "invalid argument",
	EIO - APPLICATION_ERROR:             "input/output error",
	EISCONN - APPLICATION_ERROR:         "transport endpoint is already connected",
	EISDIR - APPLICATION_ERROR:          "is a directory",
	EISNAM - APPLICATION_ERROR:          "is a named type file",
	EKEYEXPIRED - APPLICATION_ERROR:     "key has expired",
	EKEYREJECTED - APPLICATION_ERROR:    "key was rejected by service",
	EKEYREVOKED - APPLICATION_ERROR:     "key has been revoked",
	EL2HLT - APPLICATION_ERROR:          "level 2 halted",
	EL2NSYNC - APPLICATION_ERROR:        "level 2 not synchronized",
	EL3HLT - APPLICATION_ERROR:          "level 3 halted",
	EL3RST - APPLICATION_ERROR:          "level 3 reset",
	ELIBACC - APPLICATION_ERROR:         "can not access a needed shared library",
	ELIBBAD - APPLICATION_ERROR:         "accessing a corrupted shared library",
	ELIBEXEC - APPLICATION_ERROR:        "cannot exec a shared library directly",
	ELIBMAX - APPLICATION_ERROR:         "attempting to link in too many shared libraries",
	ELIBSCN - APPLICATION_ERROR:         ".lib section in a.out corrupted",
	ELNRNG - APPLICATION_ERROR:          "link number out of range",
	ELOOP - APPLICATION_ERROR:           "too many levels of symbolic links",
	EMEDIUMTYPE - APPLICATION_ERROR:     "wrong medium type",
	EMFILE - APPLICATION_ERROR:          "too many open files",
	EMLINK - APPLICATION_ERROR:          "too many links",
	EMSGSIZE - APPLICATION_ERROR:        "message too long",
	EMULTIHOP - APPLICATION_ERROR:       "multihop attempted",
	ENAMETOOLONG - APPLICATION_ERROR:    "file name too long",
	ENAVAIL - APPLICATION_ERROR:         "no XENIX semaphores available",
	ENETDOWN - APPLICATION_ERROR:        "network is down",
	ENETRESET - APPLICATION_ERROR:       "network dropped connection on reset",
	ENETUNREACH - APPLICATION_ERROR:     "network is unreachable",
	ENFILE - APPLICATION_ERROR:          "too many open files in system",
	ENOANO - APPLICATION_ERROR:          "no anode",
	ENOBUFS - APPLICATION_ERROR:         "no buffer space available",
	ENOCSI - APPLICATION_ERROR:          "no CSI structure available",
	ENODATA - APPLICATION_ERROR:         "no data available",
	ENODEV - APPLICATION_ERROR:          "no such device",
	ENOEXEC - APPLICATION_ERROR:         "exec format error",
	ENOKEY - APPLICATION_ERROR:          "required key not available",
	ENOLCK - APPLICATION_ERROR:          "no locks available",
	ENOLINK - APPLICATION_ERROR:         "link has been severed",
	ENOMEDIUM - APPLICATION_ERROR:       "no medium found",
	ENOMEM - APPLICATION_ERROR:          "cannot allocate memory",
	ENOMSG - APPLICATION_ERROR:          "no message of desired type",
	ENONET - APPLICATION_ERROR:          "machine is not on the network",
	ENOPKG - APPLICATION_ERROR:          "package not installed",
	ENOPROTOOPT - APPLICATION_ERROR:     "protocol not available",
	ENOSPC - APPLICATION_ERROR:          "no space left on device",
	ENOSR - APPLICATION_ERROR:           "out of streams resources",
	ENOSTR - APPLICATION_ERROR:          "device not a stream",
	ENOSYS - APPLICATION_ERROR:          "function not implemented",
	ENOTBLK - APPLICATION_ERROR:         "block device required",
	ENOTCONN - APPLICATION_ERROR:        "transport endpoint is not connected",
	ENOTEMPTY - APPLICATION_ERROR:       "directory not empty",
	ENOTNAM - APPLICATION_ERROR:         "not a XENIX named type file",
	ENOTRECOVERABLE - APPLICATION_ERROR: "state not recoverable",
	ENOTSOCK - APPLICATION_ERROR:        "socket operation on non-socket",
	ENOTSUP - APPLICATION_ERROR:         "operation not supported",
	ENOTTY - APPLICATION_ERROR:          "inappropriate ioctl for device",
	ENOTUNIQ - APPLICATION_ERROR:        "name not unique on network",
	ENXIO - APPLICATION_ERROR:           "no such device or address",
	EOPNOTSUPP - APPLICATION_ERROR:      "operation not supported",
	EOVERFLOW - APPLICATION_ERROR:       "value too large for defined data type",
	EOWNERDEAD - APPLICATION_ERROR:      "owner died",
	EPERM - APPLICATION_ERROR:           "operation not permitted",
	EPFNOSUPPORT - APPLICATION_ERROR:    "protocol family not supported",
	EPIPE - APPLICATION_ERROR:           "broken pipe",
	EPROTO - APPLICATION_ERROR:          "protocol error",
	EPROTONOSUPPORT - APPLICATION_ERROR: "protocol not supported",
	EPROTOTYPE - APPLICATION_ERROR:      "protocol wrong type for socket",
	ERANGE - APPLICATION_ERROR:          "numerical result out of range",
	EREMCHG - APPLICATION_ERROR:         "remote address changed",
	EREMOTE - APPLICATION_ERROR:         "object is remote",
	EREMOTEIO - APPLICATION_ERROR:       "remote I/O error",
	ERESTART - APPLICATION_ERROR:        "interrupted system call should be restarted",
	EROFS - APPLICATION_ERROR:           "read-only file system",
	ESHUTDOWN - APPLICATION_ERROR:       "cannot send after transport endpoint shutdown",
	ESOCKTNOSUPPORT - APPLICATION_ERROR: "socket type not supported",
	ESPIPE - APPLICATION_ERROR:          "illegal seek",
	ESRCH - APPLICATION_ERROR:           "no such process",
	ESRMNT - APPLICATION_ERROR:          "srmount error",
	ESTALE - APPLICATION_ERROR:          "stale NFS file handle",
	ESTRPIPE - APPLICATION_ERROR:        "streams pipe error",
	ETIME - APPLICATION_ERROR:           "timer expired",
	ETIMEDOUT - APPLICATION_ERROR:       "connection timed out",
	ETOOMANYREFS - APPLICATION_ERROR:    "too many references: cannot splice",
	ETXTBSY - APPLICATION_ERROR:         "text file busy",
	EUCLEAN - APPLICATION_ERROR:         "structure needs cleaning",
	EUNATCH - APPLICATION_ERROR:         "protocol driver not attached",
	EUSERS - APPLICATION_ERROR:          "too many users",
	EWOULDBLOCK - APPLICATION_ERROR:     "resource temporarily unavailable",
	EXDEV - APPLICATION_ERROR:           "invalid cross-device link",
	EXFULL - APPLICATION_ERROR:          "exchange full",
	EWINDOWS - APPLICATION_ERROR:        "not supported by windows",
}
