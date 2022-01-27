// Package sha256 对推送数据进行加密和解密的工具包
package sha256

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type EncryptObj struct {
	Encrypt string `json:"encrypt"`
}

// Encrypt 数据加密
func Encrypt(data []byte, key string) (string, error) {
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return "", fmt.Errorf("AESNewCipher Error[%v]", err)
	}
	content := pKCS5Padding(data, aes.BlockSize)
	cipherBytes := make([]byte, len(content))
	iv := RandString(aes.BlockSize)
	aesEncrypt := cipher.NewCBCEncrypter(block, []byte(iv))
	aesEncrypt.CryptBlocks(cipherBytes, content)
	buf := BytesCombine([]byte(iv), cipherBytes)
	return base64.StdEncoding.EncodeToString(buf), nil
}

// Decrypt 解密接收到的推送数据
func Decrypt(encrypt string, key string) (string, error) {
	buf, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", fmt.Errorf("base64StdEncode Error[%v]", err)
	}
	if len(buf) < aes.BlockSize {
		return "", errors.New("cipher  too short")
	}
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return "", fmt.Errorf("AESNewCipher Error[%v]", err)
	}
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(buf)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(buf, buf)
	return string(pKCS5Trimming(buf)), nil
}

// pKCS5Padding 不足位的填充补足
func pKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// pKCS5Trimming 去掉填充
func pKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// RandString 生成随机字符串，用来生成IV
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		a := r.Intn(26) + 65
		b[i] = byte(a)
	}
	return string(b)
}

// BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
