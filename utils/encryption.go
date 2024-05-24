package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// padKey 调整密钥长度为 32 字节
func padKey(key string) []byte {
	const keySize = 32
	paddedKey := make([]byte, keySize)
	copy(paddedKey, key)
	return paddedKey
}

// Encrypt 加密字符串
func Encrypt(text, key string) (string, error) {
	paddedKey := padKey(key)
	block, err := aes.NewCipher(paddedKey)
	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, paddedKey[:aes.BlockSize])
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt 解密字符串
func Decrypt(encryptedText, key string) (string, error) {
	paddedKey := padKey(key)
	block, err := aes.NewCipher(paddedKey)
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, paddedKey[:aes.BlockSize])
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

// RunEncryption 运行加密函数
func RunEncryption() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("无法加载 .env 文件:", err)
		return
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		fmt.Println("未找到 SECRET_KEY 环境变量")
		return
	}

	fmt.Println("请输入要加密的文本:")
	var text string
	_, _ = fmt.Scanln(&text)

	encryptedText, err := Encrypt(text, secretKey)
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}

	fmt.Println("加密后的文本:", encryptedText)
}
