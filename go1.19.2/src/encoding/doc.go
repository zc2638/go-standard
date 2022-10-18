// Package encoding defines interfaces shared by other packages that
// convert data to and from byte-level and textual representations.
// Packages that check for these interfaces include encoding/gob,
// encoding/json, and encoding/xml. As a result, implementing an
// interface once can make a type useful in multiple encodings.
// Standard types that implement these interfaces include time.Time and net.IP.
// The interfaces come in pairs that produce and consume encoded data.

package encoding
