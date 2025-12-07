package gsign

import (
	"os"
	"testing"
)

func TestGsign(t *testing.T) {
	s, _ := New()
	s.AddDomain("domain.com")
	s.AddIP("127.0.0.1")
	ca, cak, pem, key, _ := s.Generate()
	os.WriteFile("ca.crt", []byte(ca), 0777)
	os.WriteFile("ca.key", []byte(cak), 0777)
	os.WriteFile("domain.pem", []byte(pem), 0777)
	os.WriteFile("domain.key", []byte(key), 0777)
}
