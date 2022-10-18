// Package lzw implements the Lempel-Ziv-Welch compressed data format,
// described in T. A. Welch, “A Technique for High-Performance Data
// Compression”, Computer, 17(6) (June 1984), pp 8-19.
//
// In particular, it implements LZW as used by the GIF and PDF file
// formats, which means variable-width codes up to 12 bits and the first
// two non-literal codes are a clear code and an EOF code.
//
// The TIFF file format uses a similar but incompatible version of the LZW
// algorithm. See the golang.org/x/image/tiff/lzw package for an
// implementation.

package lzw
