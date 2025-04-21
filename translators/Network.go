package translators

import (
	"net"
	"net/netip"
)

type (
	// Represents a cmd line argument that will be translated to a [netip.Addr]
	// type.
	Addr struct{}

	// Represents a cmd line argument that should be considered a port value.
	// Nothing more than an alias for BuiltinUint16
	Port = BuiltinUint16

	// Represents a cmd line argument that will be translated to a
	// [netip.AddrPort] type.
	AddrPort struct{}

	// Represents a cmd line argument that will be translated to a
	// [netip.Prefix] type.
	AddrPrefix struct{}

	// Represents a cmd line argument that will be translated to a
	// [net.HardwareAddr] (a.k.a. a MAC address) type.
	HardwareAddr struct{}
)

func (_ Addr) Translate(arg string) (netip.Addr, error) {
	return netip.ParseAddr(arg)
}
func (_ Addr) Reset() {}
func (_ Addr) HelpAddendum() string {
	return "Expecting a valid IPv4 or IPv6 address (i.e. 192.0.2.1 or 2001:db8::68)"
}

func (_ AddrPort) Translate(arg string) (netip.AddrPort, error) {
	return netip.ParseAddrPort(arg)
}
func (_ AddrPort) Reset() {}
func (_ AddrPort) HelpAddendum() string {
	return "Expecting a valid IPv4 or IPv6 address with a port (i.e. 192.0.2.1:1000 or 2001:db8::68:1000)"
}

func (_ AddrPrefix) Translate(arg string) (netip.Prefix, error) {
	return netip.ParsePrefix(arg)
}
func (_ AddrPrefix) Reset() {}
func (_ AddrPrefix) HelpAddendum() string {
	return "Expecting an IP address prefix (i.e. 192.168.1.0/24 or 2001:db8::/32)"
}

func (_ HardwareAddr) Translate(arg string) (net.HardwareAddr, error) {
	return net.ParseMAC(arg)
}
func (_ HardwareAddr) Reset() {}
func (_ HardwareAddr) HelpAddendum() string {
	return "Expecting a valid hardware MAC address"
}
