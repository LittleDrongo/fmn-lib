package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

// Метод шифрует текст используя ключ 16, 24 или 32 байта длина ключа для AES-128, AES-192 и AES-256.
func Encrypt(key, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext), nil
}

// Метод расшифровывает зашифрованный текст используя ключ
func Decrypt(key, cryptoText string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// Генерирует случайный ключ указанной длины
func GetRandomKey(length int) (string, error) {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

// Преобразует ключ произвольной длины в ключ длиной 256 бит (32 байта) с помощью SHA-256
func ConvertToAES256Key(key string) (string, error) {
	keyBytes := []byte(key)

	hasher := sha256.New()
	_, err := hasher.Write(keyBytes)
	if err != nil {
		return "", err
	}

	hashedKey := hasher.Sum(nil)
	hashedKeyStr := hex.EncodeToString(hashedKey)

	return hashedKeyStr[:32], nil
}
