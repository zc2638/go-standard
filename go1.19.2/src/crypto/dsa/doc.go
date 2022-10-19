// Package dsa implements the Digital Signature Algorithm, as defined in FIPS 186-3.
//
// The DSA operations in this package are not implemented using constant-time algorithms.
//
// Deprecated: DSA is a legacy algorithm, and modern alternatives such as
// Ed25519 (implemented by package crypto/ed25519) should be used instead. Keys
// with 1024-bit moduli (L1024N160 parameters) are cryptographically weak, while
// bigger keys are not widely supported. Note that FIPS 186-5 no longer approves
// DSA for signature generation.
package dsa
