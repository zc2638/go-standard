// Package maphash provides hash functions on byte sequences.
// These hash functions are intended to be used to implement hash tables or
// other data structures that need to map arbitrary strings or byte
// sequences to a uniform distribution on unsigned 64-bit integers.
// Each different instance of a hash table or data structure should use its own Seed.
//
// The hash functions are not cryptographically secure.
// (See crypto/sha256 and crypto/sha512 for cryptographic use.)
package maphash
