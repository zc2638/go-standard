// Package netip defines an IP address type that's a small value type.
// Building on that Addr type, the package also defines AddrPort (an
// IP address and a port), and Prefix (an IP address and a bit length
// prefix).
//
// Compared to the net.IP type, this package's Addr type takes less
// memory, is immutable, and is comparable (supports == and being a
// map key).
package netip
