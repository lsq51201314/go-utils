package utils

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc/credentials"
)

func NewServer(cafile, pemfiile, keyfile string) (c credentials.TransportCredentials, err error) {
	var cert tls.Certificate
	if cert, err = tls.LoadX509KeyPair(pemfiile, keyfile); err != nil {
		return
	}
	certPool := x509.NewCertPool()
	var ca []byte
	if ca, err = os.ReadFile(cafile); err != nil {
		return
	}
	certPool.AppendCertsFromPEM(ca)
	c = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return
}

func NewClient(cafile, pemfiile, keyfile string) (c credentials.TransportCredentials, err error) {
	var cert tls.Certificate
	if cert, err = tls.LoadX509KeyPair(pemfiile, keyfile); err != nil {
		return
	}
	certPool := x509.NewCertPool()
	var ca []byte
	if ca, err = os.ReadFile(cafile); err != nil {
		return
	}
	certPool.AppendCertsFromPEM(ca)
	c = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	})
	return
}
