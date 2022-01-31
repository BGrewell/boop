package internal

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func NewProxyController(ctx context.Context) (controller *ProxyController, err error) {
	p := &ProxyController{
		proxies:     make([]Proxy, 0),
		exitOnError: false,
	}
	v4, err := ListenersV4Loopback()
	if err != nil {
		return nil, err
	}
	v6, err := ListenersV6Loopback()
	if err != nil {
		return nil, err
	}
	ports := append(v4, v6...)
	err = p.proxyPorts(ctx, ports)
	if err != nil {
		return nil, err
	}

	return p, nil
}

type ProxyController struct {
	proxies     []Proxy
	exitOnError bool
}

func (pc *ProxyController) Start(ctx context.Context) (err error) {
	for _, proxy := range pc.proxies {
		err = proxy.Start(ctx)
		if err != nil {
			if pc.exitOnError {
				return err
			}
			log.Warnf("%s\n", err)
		}
	}
	return nil
}

func (pc *ProxyController) Stop(ctx context.Context) (err error) {
	for _, proxy := range pc.proxies {
		err = proxy.Stop(ctx)
		if pc.exitOnError {
			return err
		}
		log.Warnf("%s\n", err)
	}
	return nil
}

func (pc *ProxyController) proxyPorts(ctx context.Context, ports []ListeningPort) (err error) {
	for _, port := range ports {
		switch port.Protocol {
		case "tcp":
			p, err := NewTcpProxy(ctx, port)
			if err != nil {
				return err
			}
			pc.proxies = append(pc.proxies, p)
		case "udp":
			p, err := NewUdpProxy(ctx, port)
			if err != nil {
				return err
			}
			pc.proxies = append(pc.proxies, p)
		}
	}
	return nil
}
