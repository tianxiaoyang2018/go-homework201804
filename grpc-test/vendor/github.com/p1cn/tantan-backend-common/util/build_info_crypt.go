package util

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
)

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}
func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func BuildInfoEncrypt(s, secret string) ([]byte, error) {
	if len(secret) < 32 {
		return nil, errors.New("length of SecretKey less than 32")
	}

	buf := &bytes.Buffer{}
	w := gzip.NewWriter(buf)

	w.Write([]byte(s))
	w.Close()

	aes_key := []byte(secret)[0:32]

	block, err := aes.NewCipher(aes_key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, 16)
	io.ReadFull(rand.Reader, iv)

	mode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := PKCS7Padding(buf.Bytes(), block.BlockSize())

	mode.CryptBlocks(ciphertext, ciphertext)
	r := append(iv, ciphertext...)
	return r, nil
}

func BuildInfoDecrypt(bs []byte, secret string) (string, error) {
	if len(secret) < 32 || len(bs) < 16 {
		return "", errors.New("length of SecretKey less than 32")
	}
	aes_key := []byte(secret)[0:32]
	block, err := aes.NewCipher(aes_key)
	if err != nil {
		return "", err
	}

	iv := bs[:aes.BlockSize]
	mode := cipher.NewCBCDecrypter(block, iv)
	ciphertext := bs[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("input not full blocks")
	}

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext = PKCS7UnPadding(plaintext)

	buf := bytes.NewBuffer(plaintext)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return "", err
	}
	defer r.Close()
	result, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
