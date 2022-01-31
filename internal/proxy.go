package internal

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net"
)

type Proxy interface {
	Start(context.Context) error
	Stop(context.Context) error
}

func NewTcpProxy(ctx context.Context, port ListeningPort) (proxy *TcpProxy, err error) {
	return &TcpProxy{
		ip:       viper.GetString("ip"),
		portinfo: port,
	}, nil
}

type TcpProxy struct {
	ip       string
	listener *net.TCPListener
	portinfo ListeningPort
	running  bool
}

func (p *TcpProxy) Start(ctx context.Context) (err error) {
	p.listener, err = net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP(p.ip),
		Port: int(p.portinfo.Port),
	})
	if err != nil {
		return err
	}
	p.running = true
	go p.listenTcp(err)
	return nil
}

func (p *TcpProxy) Stop(ctx context.Context) (err error) {
	p.running = false
	return p.listener.Close()
}

func (p *TcpProxy) listenTcp(err error) {
	defer p.listener.Close()
	log.Infof("Listening on %s:%d\n", p.ip, p.portinfo.Port)
	for p.running {
		conn, err := p.listener.Accept()
		if err != nil {
			log.Errorf("Error accepting connection: %s\n", err)
			continue
		}
		go p.handleTcp(conn)
	}
}

func (p *TcpProxy) handleTcp(conn net.Conn) {
	local, err := net.Dial("tcp", fmt.Sprintf("%s:%d", p.portinfo.IP, p.portinfo.Port))
	if err != nil {
		log.Errorf("Error connecting to local service %s:%d: %s\n", p.portinfo.IP, p.portinfo.Port, err)
		return
	}
	go func() {
		_, err := io.Copy(local, conn)
		if err != nil {
			log.Errorf("Error copying data: %s\n", err)
		}
		log.Warn("Client disconnected")
		local.Close()
		conn.Close()
	}()

	_, err = io.Copy(conn, local)
	if err != nil {
		log.Errorf("Error copying data: %s\n", err)
	}
	log.Warn("[***] Service disconnected")
	local.Close()
	conn.Close()
}

func NewUdpProxy(ctx context.Context, port ListeningPort) (proxy *UdpProxy, err error) {
	return &UdpProxy{}, nil
}

type UdpProxy struct {
}

func (p *UdpProxy) Start(ctx context.Context) (err error) {
	return nil
}

func (p *UdpProxy) Stop(ctx context.Context) (err error) {
	return nil
}
