package sha256

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestEncrypt(t *testing.T) {
	str := "aaaaaaaaaaaaaaaaaaaaaaaaa"
	key := "bbbbbbbbbbbbbbbbb"
	enc, err := Encrypt([]byte(str), key)
	if err != nil {
		t.Error(err)
	}
	log.Printf("enc:%s", enc)
}

func TestDecrypt(t *testing.T) {
	enc := "S1daWEhHSFlZR0hWREJWWI7OG3mB+rKov/U9mclsa72JBcVymmlzXwlYQE0Ch3WV"
	key := "bbbbbbbbbbbbbbbbb"
	dsc, err := Decrypt(enc, key)
	if err != nil {
		t.Error(err)
	}
	log.Printf("dsc:%s", dsc)
	if dsc != "aaaaaaaaaaaaaaaaaaaaaaaaa" {
		t.Error("error")
	}
}
