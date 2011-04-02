// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"runtime"
	"syscall"
)

func epipecheck(file *File, e syscall.Error) {
}


// DevNull is the name of the operating system's ``null device.''
// On Unix-like systems, it is "/dev/null"; on Windows, "NUL".
const DevNull = "/dev/null"

// Open opens the named file with specified flag (O_RDONLY etc.) and perm.
// If successful, methods on the returned File can be used for I/O.
// It returns the File and an Error, if any.
func Open(name string, flag int, perm uint32) (file *File, err Error) {
	var fd int
	var e syscall.Error

	syscall.ForkLock.RLock()
	if flag&O_CREAT == O_CREAT {
		fd, e = syscall.Create(name, flag & ^O_CREAT, perm)
	} else {
		fd, e = syscall.Open(name, flag)
	}
	syscall.ForkLock.RUnlock()

	if e != nil {
		return nil, &PathError{"open", name, e}
	}

	return NewFile(fd, name), nil
}

// Close closes the File, rendering it unusable for I/O.
// It returns an Error, if any.
func (file *File) Close() Error {
	if file == nil || file.fd < 0 {
		return Ebadfd
	}
	var err Error
	syscall.ForkLock.RLock()
	if e := syscall.Close(file.fd); e != nil {
		err = &PathError{"close", file.name, e}
	}
	syscall.ForkLock.RUnlock()
	file.fd = -1 // so it can't be closed again

	// no need for a finalizer anymore
	runtime.SetFinalizer(file, nil)
	return err
}

// Stat returns the FileInfo structure describing file.
// It returns the FileInfo and an error, if any.
func (file *File) Stat() (fi *FileInfo, err Error) {
	return dirstat(file)
}

// Truncate changes the size of the file.
// It does not change the I/O offset.
func (f *File) Truncate(size int64) Error {
	var d Dir
	d.Null()

	d.Length = uint64(size)

	if e := syscall.Fwstat(f.fd, pdir(nil, &d)); iserror(e) {
		return &PathError{"truncate", f.name, e}
	}
	return nil
}

// Chmod changes the mode of the file to mode.
func (f *File) Chmod(mode uint32) Error {
	var d Dir
	d.Null()

	d.Mode = mode & 0777

	if e := syscall.Fwstat(f.fd, pdir(nil, &d)); iserror(e) {
		return &PathError{"chmod", f.name, e}
	}
	return nil
}

// Sync commits the current contents of the file to stable storage.
// Typically, this means flushing the file system's in-memory copy
// of recently written data to disk.
func (f *File) Sync() (err Error) {
	if f == nil {
		return EINVAL
	}

	var d Dir
	d.Null()

	if e := syscall.Fwstat(f.fd, pdir(nil, &d)); iserror(e) {
		return NewSyscallError("fsync", e)
	}
	return nil
}

// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link's target.
func Truncate(name string, size int64) Error {
	var d Dir
	d.Null()

	d.Length = uint64(size)

	if e := syscall.Wstat(name, pdir(nil, &d)); iserror(e) {
		return &PathError{"truncate", name, e}
	}
	return nil
}

// Remove removes the named file or directory.
func Remove(name string) Error {
	if e := syscall.Remove(name); iserror(e) {
		return &PathError{"remove", name, e}
	}
	return nil
}

// Rename renames a file.
func Rename(oldname, newname string) Error {
	var d Dir
	d.Null()

	d.Name = newname

	if e := syscall.Wstat(oldname, pdir(nil, &d)); iserror(e) {
		return &PathError{"rename", oldname, e}
	}
	return nil
}

// Chmod changes the mode of the named file to mode.
func Chmod(name string, mode uint32) Error {
	var d Dir
	d.Null()

	d.Mode = mode & 0777

	if e := syscall.Wstat(name, pdir(nil, &d)); iserror(e) {
		return &PathError{"chmod", name, e}
	}
	return nil
}

// ChownPlan9 changes the uid and gid strings of the named file.
func ChownPlan9(name, uid, gid string) Error {
	var d Dir
	d.Null()

	d.Uid = uid
	d.Gid = gid

	if e := syscall.Wstat(name, pdir(nil, &d)); iserror(e) {
		return &PathError{"chown_plan9", name, e}
	}
	return nil
}

// Chtimes changes the access and modification times of the named
// file, similar to the Unix utime() or utimes() functions.
//
// The argument times are in nanoseconds, although the underlying
// filesystem may truncate or round the values to a more
// coarse time unit.
func Chtimes(name string, atimeNs int64, mtimeNs int64) Error {
	var d Dir
	d.Null()

	d.Atime = uint32(atimeNs / 1e9)
	d.Mtime = uint32(mtimeNs / 1e9)

	if e := syscall.Wstat(name, pdir(nil, &d)); iserror(e) {
		return &PathError{"chtimes", name, e}
	}
	return nil
}

func Pipe() (r *File, w *File, err Error) {
	var p [2]int

	syscall.ForkLock.RLock()
	if e := syscall.Pipe(p[0:]); iserror(e) {
		syscall.ForkLock.RUnlock()
		return nil, nil, NewSyscallError("pipe", e)
	}
	syscall.ForkLock.RUnlock()

	return NewFile(p[0], "|0"), NewFile(p[1], "|1"), nil
}


// not supported on Plan 9

// Link creates a hard link.
func Link(oldname, newname string) Error {
	return EPLAN9
}

func Symlink(oldname, newname string) Error {
	return EPLAN9
}

func Readlink(name string) (string, Error) {
	return "", EPLAN9
}

func Chown(name string, uid, gid int) Error {
	return EPLAN9
}

func Lchown(name string, uid, gid int) Error {
	return EPLAN9
}

func (f *File) Chown(uid, gid int) Error {
	return EPLAN9
}
