package lark

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// AESDecrypt 解密
func AESDecrypt(base64CipherText string, key []byte) (unpadDecrypted []byte, err error) {
	sum := sha256.Sum256(key)

	key = []byte(``)
	for _, b := range sum {
		key = append(key, b)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("len(crypt) < aes.BlockSize")
		return
	}
	cbc := cipher.NewCBCDecrypter(block, ciphertext[:aes.BlockSize])
	ciphertext = ciphertext[aes.BlockSize:]
	decrypted := make([]byte, len(ciphertext))
	cbc.CryptBlocks(decrypted, ciphertext)

	unpadDecrypted = PKCS5Trimming(decrypted)
	return
}

// PKCS5Padding PKCS5Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Trimming PKCS5Trimming
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
