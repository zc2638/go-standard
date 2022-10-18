// Package tzdata provides an embedded copy of the timezone database.
// If this package is imported anywhere in the program, then if
// the time package cannot find tzdata files on the system,
// it will use this embedded information.
//
// Importing this package will increase the size of a program by about
// 450 KB.
//
// This package should normally be imported by a program's main package,
// not by a library. Libraries normally shouldn't decide whether to
// include the timezone database in a program.
//
// This package will be automatically imported if you build with
// -tags timetzdata.
package tzdata
