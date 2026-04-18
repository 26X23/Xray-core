package socket

import (
	"github.com/xtls/xray-core/common/net"
)

var strategy = [][]byte{
	//              name        strategy,   prefer, fallback
	{0, 0, 0}, //   AsIs        none,       /,      /
	{1, 0, 0}, //   UseIP       use,        both,   none
	{1, 4, 0}, //   UseIPv4     use,        4,      none
	{1, 6, 0}, //   UseIPv6     use,        6,      none
	{1, 4, 6}, //   UseIPv4v6   use,        4,      6
	{1, 6, 4}, //   UseIPv6v4   use,        6,      4
	{2, 0, 0}, //   ForceIP     force,      both,   none
	{2, 4, 0}, //   ForceIPv4   force,      4,      none
	{2, 6, 0}, //   ForceIPv6   force,      6,      none
	{2, 4, 6}, //   ForceIPv4v6 force,      4,      6
	{2, 6, 4}, //   ForceIPv6v4 force,      6,      4
}

func (m SocketConfig_TProxyMode) IsEnabled() bool {
	return m != SocketConfig_Off
}

func (s DomainStrategy) HasStrategy() bool {
	return strategy[s][0] != 0
}

func (s DomainStrategy) ForceIP() bool {
	return strategy[s][0] == 2
}

func (s DomainStrategy) PreferIP4() bool {
	return strategy[s][1] == 4 || strategy[s][1] == 0
}

func (s DomainStrategy) PreferIP6() bool {
	return strategy[s][1] == 6 || strategy[s][1] == 0
}

func (s DomainStrategy) HasFallback() bool {
	return strategy[s][2] != 0
}

func (s DomainStrategy) FallbackIP4() bool {
	return strategy[s][2] == 4
}

func (s DomainStrategy) FallbackIP6() bool {
	return strategy[s][2] == 6
}

func (s DomainStrategy) GetDynamicStrategy(addrFamily net.AddressFamily) DomainStrategy {
	if addrFamily.IsDomain() {
		return s
	}
	switch s {
	case DomainStrategy_USE_IP:
		if addrFamily.IsIPv4() {
			return DomainStrategy_USE_IP46
		} else if addrFamily.IsIPv6() {
			return DomainStrategy_USE_IP64
		}
	case DomainStrategy_FORCE_IP:
		if addrFamily.IsIPv4() {
			return DomainStrategy_FORCE_IP46
		} else if addrFamily.IsIPv6() {
			return DomainStrategy_FORCE_IP64
		}
	default:
	}
	return s
}

func (v *SocketConfig) ParseTFOValue() int {
	if v.Tfo == 0 {
		return -1
	}
	tfo := int(v.Tfo)
	if tfo < 0 {
		tfo = 0
	}
	return tfo
}
