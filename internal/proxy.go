package internal

import "context"

type Proxy interface {
	Start(context.Context) error
	Stop(context.Context) error
}

func NewTcpProxy(ctx context.Context, port ListeningPort) (proxy *TcpProxy, err error) {
	return &TcpProxy{}, nil
}

type TcpProxy struct {
}

func (p *TcpProxy) Start(ctx context.Context) error {
	return nil
}

func (p *TcpProxy) Stop(ctx context.Context) error {
	return nil
}

func NewUdpProxy(ctx context.Context, port ListeningPort) (proxy *UdpProxy, err error) {
	return &UdpProxy{}, nil
}

type UdpProxy struct {
}

func (p *UdpProxy) Start(ctx context.Context) error {
	return nil
}

func (p *UdpProxy) Stop(ctx context.Context) error {
	return nil
}
