package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func RsaDecrypt(src []byte, key string) ([]byte, error) {
	file, err := os.Open(key)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, _ = file.Read(buf)
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("decode pem failure")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	result, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	if err != nil {
		return nil, err
	}
	return result, nil
}
