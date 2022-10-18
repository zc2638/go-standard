// Package mail implements parsing of mail messages.
//
// For the most part, this package follows the syntax as specified by RFC 5322 and
// extended by RFC 6532.
// Notable divergences:
//   - Obsolete address formats are not parsed, including addresses with
//     embedded route information.
//   - The full range of spacing (the CFWS syntax element) is not supported,
//     such as breaking addresses across lines.
//   - No unicode normalization is performed.
//   - The special characters ()[]:;@\, are allowed to appear unquoted in names.
package mail
