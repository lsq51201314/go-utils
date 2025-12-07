package gsign

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"
)

// 自签证书
type GSign struct {
	country            string
	province           string
	locality           string
	organization       string
	organizationalUnit string
	commonName         string
	domains            []string
	ips                []string
}

// 参数顺序：Country、Province、Locality、Organization、OrganizationalUnit、CommonName
func New(params ...string) (*GSign, error) {
	var obj GSign
	if len(params) > 0 {
		obj.country = params[0]
	} else {
		obj.country = "cn"
	}

	if len(params) > 1 {
		obj.province = params[1]
	} else {
		obj.province = "fujian"
	}

	if len(params) > 2 {
		obj.locality = params[2]
	} else {
		obj.locality = "putian"
	}

	if len(params) > 3 {
		obj.organization = params[3]
	} else {
		obj.organization = "No.9"
	}

	if len(params) > 4 {
		obj.organizationalUnit = params[4]
	} else {
		obj.organizationalUnit = "CA"
	}

	if len(params) > 5 {
		obj.commonName = params[5]
	} else {
		obj.commonName = "No.9 CA"
	}
	obj.domains = make([]string, 0)
	obj.ips = make([]string, 0)
	return &obj, nil
}

// 添加域名
func (t *GSign) AddDomain(domain string, reset ...bool) {
	if len(reset) > 0 && reset[0] {
		t.domains = make([]string, 0)
	}
	t.domains = append(t.domains, domain)
}

// 添加地址
func (t *GSign) AddIP(ip string, reset ...bool) {
	if len(reset) > 0 && reset[0] {
		t.ips = make([]string, 0)
	}
	t.ips = append(t.ips, ip)
}

func (t *GSign) generateCA(duration int) (*x509.Certificate, *rsa.PrivateKey, error) {
	caKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	caTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2024),
		Subject: pkix.Name{
			Country:            []string{t.country},
			Province:           []string{t.province},
			Locality:           []string{t.locality},
			Organization:       []string{t.organization},
			OrganizationalUnit: []string{t.organizationalUnit},
			CommonName:         t.commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(duration, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	caCertDER, err := x509.CreateCertificate(rand.Reader, caTemplate, caTemplate, &caKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}
	caCert, err := x509.ParseCertificate(caCertDER)
	if err != nil {
		return nil, nil, err
	}
	return caCert, caKey, nil
}

func (t *GSign) generateClientCert(duration int, caCert *x509.Certificate, caKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error) {
	clientKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	clientTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2024),
		Subject: pkix.Name{
			Country:            []string{t.country},
			Province:           []string{t.province},
			Locality:           []string{t.locality},
			Organization:       []string{t.organization},
			OrganizationalUnit: []string{t.organizationalUnit},
			CommonName:         t.commonName,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(duration, 0, 0),
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		DNSNames:    t.domains,
		IPAddresses: make([]net.IP, 0),
	}
	for _, v := range t.ips {
		clientTemplate.IPAddresses = append(clientTemplate.IPAddresses, net.ParseIP(v))
	}
	clientCertDER, err := x509.CreateCertificate(rand.Reader, clientTemplate, caCert, &clientKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}
	clientCert, err := x509.ParseCertificate(clientCertDER)
	if err != nil {
		return nil, nil, err
	}
	return clientCert, clientKey, nil
}

// 默认1年
func (t *GSign) Generate(duration ...int) (capem, cakey, sigpem, sigkey string, err error) {
	dur := 1
	if len(duration) > 0 && duration[0] > 0 {
		dur = duration[0]
	}
	caCert, caKey, err := t.generateCA(dur)
	if err != nil {
		return "", "", "", "", err
	}
	clientCert, clientKey, err := t.generateClientCert(dur, caCert, caKey)
	if err != nil {
		return "", "", "", "", err
	}
	capem = string(pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCert.Raw,
	}))
	cakey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caKey),
	}))
	sigpem = string(pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCert.Raw,
	}))
	sigkey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(clientKey),
	}))
	return
}
