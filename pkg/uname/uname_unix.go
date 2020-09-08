package uname

import (
	"bytes"

	"golang.org/x/sys/unix"
)

// Uname extends unix.Utsname to provide convience methods
// to return uname information as strings.
type Uname struct {
	unix.Utsname
}

// New calls unix.Uname
func New() (*Uname, error) {
	var u Uname
	u.Utsname = unix.Utsname{}
	return &u, unix.Uname(&u.Utsname)
}

// Machine returns the system's machine name as a string
func (u *Uname) Machine() string {
	return toString(u.Utsname.Machine[:])
}

// Nodename returns the system's node name as a string
func (u *Uname) Nodename() string {
	return toString(u.Utsname.Nodename[:])
}

// Release returns the system's release as a string
func (u *Uname) Release() string {
	return toString(u.Utsname.Release[:])
}

// Version returns the system's version as a string
func (u *Uname) Version() string {
	return toString(u.Utsname.Version[:])
}

// Sysname returns the system's name as a string.
func (u *Uname) Sysname() string {
	return toString(u.Utsname.Sysname[:])
}

func toString(d []byte) string {
	return string(d[:bytes.IndexByte(d[:], 0)])
}
