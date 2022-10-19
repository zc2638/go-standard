// Package zip provides support for reading and writing ZIP archives.
//
// See: https://www.pkware.com/appnote
//
// This package does not support disk spanning.
//
// A note about ZIP64:
//
// To be backwards compatible the FileHeader has both 32 and 64 bit Size
// fields. The 64 bit fields will always contain the correct value and
// for normal archives both fields will be the same. For files requiring
// the ZIP64 format the 32 bit fields will be 0xffffffff and the 64 bit
// fields must be used instead.
package zip
