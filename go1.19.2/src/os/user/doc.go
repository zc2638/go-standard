// Package user allows user account lookups by name or id.
//
// For most Unix systems, this package has two internal implementations of
// resolving user and group ids to names, and listing supplementary group IDs.
// One is written in pure Go and parses /etc/passwd and /etc/group. The other
// is cgo-based and relies on the standard C library (libc) routines such as
// getpwuid_r, getgrnam_r, and getgrouplist.
//
// When cgo is available, and the required routines are implemented in libc
// for a particular platform, cgo-based (libc-backed) code is used.
// This can be overridden by using osusergo build tag, which enforces
// the pure Go implementation.
package user
