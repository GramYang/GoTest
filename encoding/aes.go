package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	aesEnc := AesEncrypt{}
	arrEncrypt, err := aesEnc.Encrypt("abcde")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(arrEncrypt)
	strMsg, err := aesEnc.Decrypt(arrEncrypt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strMsg)
}

type AesEncrypt struct{}

func (*AesEncrypt) getKey() []byte {
	strKey := "1234567890123456"
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("des key 长度不能小于16")
	}
	//密钥长度只能是16字节，24字节，32字节这三个长度
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		return arrKey[:32]
	}
	if keyLen >= 24 {
		return arrKey[:24]
	}
	return arrKey[:16]
}

//AES加密数据块分组长度必须为128bit
func (a *AesEncrypt) Encrypt(strMsg string) ([]byte, error) {
	key := a.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMsg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMsg))
	return encrypted, nil
}

func (a *AesEncrypt) Decrypt(src []byte) (strDes string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	key := a.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}
