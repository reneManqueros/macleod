package models

import (
	"crypto/tls"
	"io"
	"log"
	"net"
)

type Server struct {
	Listener net.Listener
}

func (s *Server) New() {
	tlsConfig := &tls.Config{
		GetConfigForClient: getConfigForClient,
	}
	var err error
	s.Listener, err = tls.Listen("tcp", Config.ListenAddress, tlsConfig)
	if err != nil {
		log.Fatalf("error in tls.Listen: %s", err)
	}
}

func (s *Server) Serve() {
	for {
		clientConn, err := s.Listener.Accept()
		if err != nil {
			log.Printf("error in listener.Accept: %s", err)
			break
		}

		go s.Handle(clientConn)
	}
}

func getConfigForClient(chi *tls.ClientHelloInfo) (*tls.Config, error) {
	domain := chi.ServerName
	certificate, key, err := Config.GetCertificatesForDomain(domain)
	if err != nil {
		return nil, err
	}

	cert, err := tls.LoadX509KeyPair(certificate, key)
	if err != nil {
		log.Fatalf("error in tls.LoadX509KeyPair: %s", err)
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}, nil
}

func (s *Server) Handle(clientConn net.Conn) {
	tlsConn, ok := clientConn.(*tls.Conn)
	if ok {
		err := tlsConn.Handshake()
		if err != nil {
			log.Printf("error in tls.Handshake: %s", err)
			_ = clientConn.Close()
			return
		}
		domain := tlsConn.ConnectionState().ServerName
		backend, err := Config.GetBackendForDomain(domain)
		if err != nil {
			log.Println(err)
			return
		}
		backendConn, err := net.Dial("tcp", backend)
		if err != nil {
			log.Printf("error in net.Dial: %s", err)
			_ = clientConn.Close()
			return
		}

		go s.CopyIO(clientConn, backendConn)
		go s.CopyIO(backendConn, clientConn)
	}
}

func (s *Server) CopyIO(from, to io.ReadWriteCloser) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered while tunneling")
		}
	}()

	_, _ = io.Copy(from, to)
	_ = to.Close()
	_ = from.Close()
}
