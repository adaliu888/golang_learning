package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// PKCS7Padding 函数实现PKCS#7填充
func PKCS7Padding(src []byte, blockSize int) []byte {
	// 计算填充的字节数
	padding := blockSize - len(src)%blockSize
	// 创建填充的字节序列
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	// 将填充的字节序列追加到原始数据的末尾
	return append(src, pad...)
}

// 示例：使用AES加密
func main() {
	data := []byte("Hello, World!")

	// 生成密钥
	key := make([]byte, aes.BlockSize)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 填充原始数据
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padText := PKCS7Padding(data, padding)
	data = padText

	// 初始化向量IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(data))
	cbcEncrypter.CryptBlocks(ciphertext, data)

	fmt.Printf("Encrypted: %x\n", ciphertext)

	// 这里可以添加解密逻辑
}
