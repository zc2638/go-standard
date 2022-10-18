// Package image implements a basic 2-D image library.
//
// The fundamental interface is called Image. An Image contains colors, which
// are described in the image/color package.
//
// Values of the Image interface are created either by calling functions such
// as NewRGBA and NewPaletted, or by calling Decode on an io.Reader containing
// image data in a format such as GIF, JPEG or PNG. Decoding any particular
// image format requires the prior registration of a decoder function.
// Registration is typically automatic as a side effect of initializing that
// format's package so that, to decode a PNG image, it suffices to have
//
//	import _ "image/png"
//
// in a program's main package. The _ means to import a package purely for its
// initialization side effects.
//
// See "The Go image package" for more details:
// https://golang.org/doc/articles/image_package.html
package image
