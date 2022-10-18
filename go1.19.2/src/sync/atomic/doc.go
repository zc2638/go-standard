// Package atomic provides low-level atomic memory primitives
// useful for implementing synchronization algorithms.
//
// These functions require great care to be used correctly.
// Except for special, low-level applications, synchronization is better
// done with channels or the facilities of the sync package.
// Share memory by communicating;
// don't communicate by sharing memory.
//
// The swap operation, implemented by the SwapT functions, is the atomic
// equivalent of:
//
//	old = *addr
//	*addr = new
//	return old
//
// The compare-and-swap operation, implemented by the CompareAndSwapT
// functions, is the atomic equivalent of:
//
//	if *addr == old {
//		*addr = new
//		return true
//	}
//	return false
//
// The add operation, implemented by the AddT functions, is the atomic
// equivalent of:
//
//	*addr += delta
//	return *addr
//
// The load and store operations, implemented by the LoadT and StoreT
// functions, are the atomic equivalents of "return *addr" and
// "*addr = val".
//
// In the terminology of the Go memory model, if the effect of
// an atomic operation A is observed by atomic operation B,
// then A “synchronizes before” B.
// Additionally, all the atomic operations executed in a program
// behave as though executed in some sequentially consistent order.
// This definition provides the same semantics as
// C++'s sequentially consistent atomics and Java's volatile variables.
package atomic

// BUG(rsc): On 386, the 64-bit functions use instructions unavailable before the Pentium MMX.
//
// On non-Linux ARM, the 64-bit functions use instructions unavailable before the ARMv6k core.
//
// On ARM, 386, and 32-bit MIPS, it is the caller's responsibility to arrange
// for 64-bit alignment of 64-bit words accessed atomically via the primitive
// atomic functions (types Int64 and Uint64 are automatically aligned).
// The first word in an allocated struct, array, or slice; in a global
// variable; or in a local variable (because the subject of all atomic operations
// will escape to the heap) can be relied upon to be 64-bit aligned.
