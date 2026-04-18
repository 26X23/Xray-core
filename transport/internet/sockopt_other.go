//go:build js || netbsd || openbsd || solaris
// +build js netbsd openbsd solaris

package internet

import "github.com/xtls/xray-core/transport/internet/socket"

func ApplyOutboundSocketOptions(network string, address string, fd uintptr, config *socket.SocketConfig) error {
	return nil
}

func applyInboundSocketOptions(network string, fd uintptr, config *socket.SocketConfig) error {
	return nil
}

func bindAddr(fd uintptr, ip []byte, port uint32) error {
	return nil
}

func setReuseAddr(fd uintptr) error {
	return nil
}

func setReusePort(fd uintptr) error {
	return nil
}
