package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

//rsa常规解密
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
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
}

//rsa分段加密后分段解密
func RsaDecrypt1(src []byte, key string) ([]byte, error) {
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
	if err != nil {
		return nil, err
	}
	keySize, srcSize := privateKey.Size(), len(src)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		//endIndex的初值是keySize，并以keySize累加
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	return buffer.Bytes(), nil
}
