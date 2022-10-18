// Package os provides a platform-independent interface to operating system
// functionality. The design is Unix-like, although the error handling is
// Go-like; failing calls return values of type error rather than error numbers.
// Often, more information is available within the error. For example,
// if a call that takes a file name fails, such as Open or Stat, the error
// will include the failing file name when printed and will be of type
// *PathError, which may be unpacked for more information.
//
// The os interface is intended to be uniform across all operating systems.
// Features not generally available appear in the system-specific package syscall.
//
// Here is a simple example, opening a file and reading some of it.
//
//	file, err := os.Open("file.go") // For read access.
//	if err != nil {
//		log.Fatal(err)
//	}
//
// If the open fails, the error string will be self-explanatory, like
//
//	open file.go: no such file or directory
//
// The file's data can then be read into a slice of bytes. Read and
// Write take their byte counts from the length of the argument slice.
//
//	data := make([]byte, 100)
//	count, err := file.Read(data)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("read %d bytes: %q\n", count, data[:count])
//
// Note: The maximum number of concurrent operations on a File may be limited by
// the OS or the system. The number should be high, but exceeding it may degrade
// performance or cause other issues.
package os
