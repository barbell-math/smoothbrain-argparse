package translators

import (
	"net"
	"net/netip"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestAddrTranslator(t *testing.T) {
	var _t Translator[netip.Addr]

	_t = Addr{}
	val, err := _t.Translate("3.3.3.9")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val.String(), "3.3.3.9")

	_, err = _t.Translate("asdf")
	sbtest.NotNil(t, err)
}

func TestAddrPortTranslator(t *testing.T) {
	var _t Translator[netip.AddrPort]

	_t = AddrPort{}
	val, err := _t.Translate("3.3.3.9:3000")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val.String(), "3.3.3.9:3000")

	_, err = _t.Translate("asdf")
	sbtest.NotNil(t, err)
}

func TestAddrPrefixTranslator(t *testing.T) {
	var _t Translator[netip.Prefix]

	_t = AddrPrefix{}
	val, err := _t.Translate("3.3.3.9/24")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val.String(), "3.3.3.9/24")

	_, err = _t.Translate("asdf")
	sbtest.NotNil(t, err)
}

func TestHardwareAddrTranslator(t *testing.T) {
	var _t Translator[net.HardwareAddr]

	_t = HardwareAddr{}
	val, err := _t.Translate("00:00:5e:00:53:01")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val.String(), "00:00:5e:00:53:01")

	_, err = _t.Translate("asdf")
	sbtest.NotNil(t, err)
}
