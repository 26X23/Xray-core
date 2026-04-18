package dns

import (
	"net"
	"github.com/xtls/xray-core/transport/internet/socket"
)

func (c *Config) UDP() {
}

func (c *Config) WrapPacketConnClient(raw net.PacketConn, sockopt *socket.SocketConfig, level int, levelCount int) (net.PacketConn, error) {
	return NewConnClient(c, raw)
}

func (c *Config) WrapPacketConnServer(raw net.PacketConn, level int, levelCount int) (net.PacketConn, error) {
	return NewConnServer(c, raw)
}

func (c *Config) HeaderConn() {
}
