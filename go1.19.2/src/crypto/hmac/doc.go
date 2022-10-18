// Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as
// defined in U.S. Federal Information Processing Standards Publication 198.
// An HMAC is a cryptographic hash that uses a key to sign a message.
// The receiver verifies the hash by recomputing it using the same key.
//
// Receivers should be careful to use Equal to compare MACs in order to avoid
// timing side-channels:
//
//	 // ValidMAC reports whether messageMAC is a valid HMAC tag for message.
//	 func ValidMAC(message, messageMAC, key []byte) bool {
//		 mac := hmac.New(sha256.New, key)
//		 mac.Write(message)
//		 expectedMAC := mac.Sum(nil)
//		 return hmac.Equal(messageMAC, expectedMAC)
//	 }

package hmac
