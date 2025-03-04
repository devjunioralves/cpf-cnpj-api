package infrastructure

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %v", err)
	}
	return key, nil
}

func EncryptData(data string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return "", fmt.Errorf("failed to generate IV: %v", err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	dataBytes := []byte(data)
	dataBytes = padData(dataBytes)

	encrypted := make([]byte, len(dataBytes))
	mode.CryptBlocks(encrypted, dataBytes)

	encryptedData := append(iv, encrypted...)
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func DecryptData(encryptedData string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %v", err)
	}

	iv := data[:aes.BlockSize]
	encrypted := data[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	decrypted = unpadData(decrypted)

	return string(decrypted), nil
}

func padData(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, pad...)
}

func unpadData(data []byte) []byte {
	padding := data[len(data)-1]
	return data[:len(data)-int(padding)]
}
