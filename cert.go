package utils

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc/credentials"
)

// 证书实例
type Cert struct {
	Server credentials.TransportCredentials
	Client credentials.TransportCredentials
}

// 服务证书
func (c *Cert) NewServer(cafile, pemfiile, keyfile string) (err error) {
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
	c.Server = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return
}

// 客户证书
func (c *Cert) NewClient(cafile, pemfiile, keyfile string) (err error) {
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
	c.Client = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	})
	return
}
