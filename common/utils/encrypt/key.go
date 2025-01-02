package encrypt

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/redis/go-redis/v9"
)

type KeyManager struct {
	RedisClient *redis.Client
	Context     context.Context
}

func NewKeyManager(redisClient *redis.Client, ctx context.Context) *KeyManager {
	return &KeyManager{
		RedisClient: redisClient,
		Context:     ctx,
	}
}

func (manager *KeyManager) generateRSAKey(bits int) (*rsa.PrivateKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key pair: %w", err)
	}
	return privKey, nil
}

func (manager *KeyManager) encodePrivateKeyToPEM(key *rsa.PrivateKey) string {
	privKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	pemBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyBytes,
	}
	return string(pem.EncodeToMemory(&pemBlock))
}

func (manager *KeyManager) encodePublicKeyToPEM(pubKey *rsa.PublicKey) (string, error) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", fmt.Errorf("failed to encode public key: %w", err)
	}
	pemBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	return string(pem.EncodeToMemory(&pemBlock)), nil
}

func (manager *KeyManager) saveKeyToRedis(keyName, keyValue string) error {
	err := manager.RedisClient.Set(manager.Context, keyName, keyValue, time.Hour*24).Err()
	if err != nil {
		return fmt.Errorf("failed to save key to Redis: %w", err)
	}
	return nil
}

func (manager *KeyManager) GenerateAndSaveKeyPair(keyName string, bits int) (string, error) {
	privKey, err := manager.generateRSAKey(bits)
	if err != nil {
		return "", err
	}

	privKeyPEM := manager.encodePrivateKeyToPEM(privKey)
	pubKeyPEM, err := manager.encodePublicKeyToPEM(&privKey.PublicKey)
	if err != nil {
		return "", err
	}

	if err := manager.saveKeyToRedis("RSA_"+keyName, privKeyPEM); err != nil {
		return "", err
	}

	return pubKeyPEM, nil
}

func (manager *KeyManager) Decrypt(keyName string, ciphertext string) ([]byte, error) {
	privKeyPEM, err := manager.RedisClient.Get(manager.Context, "RSA_"+keyName).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errors.New("private key not found in Redis")
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve private key: %w", err)
	}

	privKey, err := manager.LoadPrivateKeyFromPEM([]byte(privKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	// 使用 OAEP 填充模式进行解密
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, decodedCiphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return plaintext, nil
}

func (manager *KeyManager) SetAESKey(keyName string, ciphertext string) error {
	key, err := manager.Decrypt(keyName, ciphertext)
	if err != nil {
		return err
	}
	encodedKey := base64.StdEncoding.EncodeToString(key)
	if err := manager.saveKeyToRedis("AES_"+keyName, encodedKey); err != nil {
		return fmt.Errorf("failed to save session key to Redis: %w", err)
	}
	return nil
}

func (manager *KeyManager) getKeyFromRedis(keyName string) (string, error) {
	key, err := manager.RedisClient.Get(manager.Context, keyName).Result()
	if errors.Is(err, redis.Nil) {
		return "", errors.New("key not found in Redis")
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve key from Redis: %w", err)
	}
	return key, nil
}

func (manager *KeyManager) GenerateAESKey(keyName string, size int) (string, error) {
	if size != 16 && size != 24 && size != 32 {
		return "", errors.New("AES key size must be 16, 24, or 32 bytes")
	}
	key := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", fmt.Errorf("failed to generate AES key: %w", err)
	}
	encodedKey := base64.StdEncoding.EncodeToString(key)
	//if err := manager.saveKeyToRedis("AES_"+keyName, encodedKey); err != nil {
	//	return "", err
	//}
	return encodedKey, nil
}

func (manager *KeyManager) DecryptAES(keyName string, text string) ([]byte, error) {
	encodedKey, err := manager.getKeyFromRedis("AES_" + keyName)
	if err != nil {
		return nil, err
	}

	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode AES key: %w", err)
	}

	ciphertext, err := base64.URLEncoding.DecodeString(text)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Text: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher block: %w", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

func (manager *KeyManager) EncryptAES(keyName string, plaintext []byte) (string, error) {
	encodedKey, err := manager.getKeyFromRedis("AES_" + keyName)
	if err != nil {
		return "", err
	}

	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode AES key: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher block: %w", err)
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %w", err)
	}

	ciphertext := make([]byte, len(plaintext)+aes.BlockSize)
	copy(ciphertext, iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// 将加密后的数据转为 Base64 字符串
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// HMAC-based signature generation
func (manager *KeyManager) GenerateSignature(data []byte, secretKey string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

// HMAC-based signature verification
func (manager *KeyManager) verifySignature(data []byte, signature, secretKey string) bool {
	expectedSignature := manager.GenerateSignature(data, secretKey)
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

// Query to JSON conversion with AES signature validation
func (manager *KeyManager) QueryToJSONWithAES(query string, userId string) (string, error) {
	parsedQuery, err := url.ParseQuery(query)
	if err != nil {
		return "", fmt.Errorf("error parsing query string: %w", err)
	}

	signature := parsedQuery.Get("signature")
	if signature == "" {
		return "", errors.New("missing signature in query string")
	}
	delete(parsedQuery, "signature")

	queryMap := make(map[string]interface{})
	for key, values := range parsedQuery {
		if len(values) == 1 {
			queryMap[key] = values[0]
		} else {
			queryMap[key] = values
		}
	}

	jsonData, err := json.Marshal(queryMap)
	if err != nil {
		return "", fmt.Errorf("error converting to JSON: %w", err)
	}

	secretKey, err := manager.getKeyFromRedis("AES_" + userId)
	if err != nil {
		return "", err
	}

	if !manager.verifySignature(jsonData, signature, secretKey) {
		return "", errors.New("invalid signature")
	}

	return string(jsonData), nil
}

// LoadPrivateKeyFromPEM 从 PEM 格式文件加载私钥
func (manager *KeyManager) LoadPrivateKeyFromPEM(pemData []byte) (*rsa.PrivateKey, error) {
	block, rest := pem.Decode(pemData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key PEM format: " + string(rest))
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	return privKey, nil
}

// LoadPublicKeyFromPEM 从 PEM 格式文件加载公钥
func (manager *KeyManager) LoadPublicKeyFromPEM(pemData []byte) (*rsa.PublicKey, error) {
	block, rest := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key PEM format" + string(rest))
	}

	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	pubKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid public key type")
	}
	return pubKey, nil
}

// EncryptWithPublicKey 使用公钥加密数据
func (manager *KeyManager) EncryptWithPublicKey(pubKey *rsa.PublicKey, plaintext []byte) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, plaintext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt: %w", err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptWithPrivateKey 使用私钥解密数据
func (manager *KeyManager) DecryptWithPrivateKey(privKey *rsa.PrivateKey, ciphertext string) ([]byte, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, decodedCiphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}
	return plaintext, nil
}
