// Package zlib implements reading and writing of zlib format compressed data,
// as specified in RFC 1950.
//
// The implementation provides filters that uncompress during reading
// and compress during writing.  For example, to write compressed data
// to a buffer:
//
// 	 var b bytes.Buffer
//	 w := zlib.NewWriter(&b)
//	 w.Write([]byte("hello, world\n"))
//	 w.Close()
//
// and to read that data back:
// 
//	 r, err := zlib.NewReader(&b)
//	 io.Copy(os.Stdout, r)
//	 r.Close()

package zlib
