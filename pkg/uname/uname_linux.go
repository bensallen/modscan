package uname

// Domainname returns the system's domainname as a string.
func (u *Uname) Domainname() string {
	return toString(u.Utsname.Domainname[:])
}
