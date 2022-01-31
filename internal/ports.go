package internal

import (
	"fmt"
	"github.com/cakturk/go-netstat/netstat"
)

type ListeningPort struct {
	IP          string
	Port        uint16
	Protocol    string
	ProcessName string
	ProcessId   int
	Owner       uint32
}

func LocalV4TcpFilter(e *netstat.SockTabEntry) bool {
	return e.LocalAddr.IP.IsLoopback() &&
		e.LocalAddr.IP.To4() != nil &&
		e.State == netstat.Listen

}

func LocalV6TcpFilter(e *netstat.SockTabEntry) bool {
	return e.LocalAddr.IP.IsLoopback() &&
		e.LocalAddr.IP.To4() == nil &&
		e.State == netstat.Listen

}

func LocalV4UdpFilter(e *netstat.SockTabEntry) bool {
	return e.LocalAddr.IP.IsLoopback() &&
		e.LocalAddr.IP.To4() != nil

}

func LocalV6UdpFilter(e *netstat.SockTabEntry) bool {
	return e.LocalAddr.IP.IsLoopback() &&
		e.LocalAddr.IP.To4() == nil

}

func ListenersV4Loopback() (ports []ListeningPort, err error) {
	udpListeners, err := EnumeratePorts("udp", 4)
	fmt.Printf("v4udp: %d\n", len(udpListeners))
	if err != nil {
		return nil, err
	}

	tcpListeners, err := EnumeratePorts("tcp", 4)
	fmt.Printf("v4tcp: %d\n", len(tcpListeners))
	if err != nil {
		return nil, err
	}

	idx := 0
	count := len(udpListeners) + len(tcpListeners)
	ports = make([]ListeningPort, count)
	for _, listener := range udpListeners {
		ports[idx] = ListeningPort{
			IP:          listener.LocalAddr.IP.String(),
			Port:        listener.LocalAddr.Port,
			ProcessName: "unkown",
			ProcessId:   -1,
			Owner:       listener.UID,
			Protocol:    "udp",
		}
		if listener.Process != nil {
			ports[idx].ProcessName = listener.Process.Name
			ports[idx].ProcessId = listener.Process.Pid
		}
		idx++
	}

	for _, listener := range tcpListeners {
		ports[idx] = ListeningPort{
			IP:          listener.LocalAddr.IP.String(),
			Port:        listener.LocalAddr.Port,
			ProcessName: "unkown",
			ProcessId:   -1,
			Owner:       listener.UID,
			Protocol:    "tcp",
		}
		if listener.Process != nil {
			ports[idx].ProcessName = listener.Process.Name
			ports[idx].ProcessId = listener.Process.Pid
		}
		idx++
	}
	return ports, nil
}

func ListenersV6Loopback() (ports []ListeningPort, err error) {
	udpListeners, err := EnumeratePorts("udp", 6)
	fmt.Printf("v6udp: %d\n", len(udpListeners))
	if err != nil {
		return nil, err
	}

	tcpListeners, err := EnumeratePorts("tcp", 6)
	fmt.Printf("v6tcp: %d\n", len(tcpListeners))
	if err != nil {
		return nil, err
	}

	idx := 0
	count := len(udpListeners) + len(tcpListeners)
	ports = make([]ListeningPort, count)
	for _, listener := range udpListeners {
		ports[idx] = ListeningPort{
			IP:          listener.LocalAddr.IP.String(),
			Port:        listener.LocalAddr.Port,
			ProcessName: "unkown",
			ProcessId:   -1,
			Owner:       listener.UID,
			Protocol:    "udp",
		}
		if listener.Process != nil {
			ports[idx].ProcessName = listener.Process.Name
			ports[idx].ProcessId = listener.Process.Pid
		}
		idx++
	}

	for _, listener := range tcpListeners {
		ports[idx] = ListeningPort{
			IP:          listener.LocalAddr.IP.String(),
			Port:        listener.LocalAddr.Port,
			ProcessName: "unkown",
			ProcessId:   -1,
			Owner:       listener.UID,
			Protocol:    "tcp",
		}
		if listener.Process != nil {
			ports[idx].ProcessName = listener.Process.Name
			ports[idx].ProcessId = listener.Process.Pid
		}
		idx++
	}
	return ports, nil
}

func EnumeratePorts(protocol string, ipVersion int) (ports []netstat.SockTabEntry, err error) {
	switch protocol {
	case "tcp":
		if ipVersion == 4 {
			return netstat.TCPSocks(LocalV4TcpFilter)
		} else if ipVersion == 6 {
			return netstat.TCP6Socks(LocalV6TcpFilter)
		}
	case "udp":
		if ipVersion == 4 {
			return netstat.UDPSocks(LocalV4UdpFilter)
		} else if ipVersion == 6 {
			return netstat.UDP6Socks(LocalV6UdpFilter)
		}
	}
	return nil, nil
}
