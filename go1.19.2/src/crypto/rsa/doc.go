// Package rsa implements RSA encryption as specified in PKCS #1 and RFC 8017.
//
// RSA is a single, fundamental operation that is used in this package to
// implement either public-key encryption or public-key signatures.
//
// The original specification for encryption and signatures with RSA is PKCS #1
// and the terms "RSA encryption" and "RSA signatures" by default refer to
// PKCS #1 version 1.5. However, that specification has flaws and new designs
// should use version 2, usually called by just OAEP and PSS, where
// possible.
//
// Two sets of interfaces are included in this package. When a more abstract
// interface isn't necessary, there are functions for encrypting/decrypting
// with v1.5/OAEP and signing/verifying with v1.5/PSS. If one needs to abstract
// over the public key primitive, the PrivateKey type implements the
// Decrypter and Signer interfaces from the crypto package.
//
// The RSA operations in this package are not implemented using constant-time algorithms.
package rsa
