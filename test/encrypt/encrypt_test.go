package encrypt

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"gomall/common/utils/encrypt"
	"io"
	"testing"
)

func TestRSA(t *testing.T) {
	manager := encrypt.NewKeyManager(nil, context.Background())
	pub, err := manager.LoadPublicKeyFromPEM([]byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxHAbmNatnfGk0S7xlXJQ
0WI/sY+ZZxJlD6yWfV0zPg6H46mWvJZj6+vJCjTFkNmETmbbmGPemDJ3E6lgCi+e
8JEk9461wYGbSFy6R71WQGbCYUv0rXZt4fnk9vTK1Aedjwckiwc42f+aA34Wwhhq
TTY++S/7nPo4Yfv9OkWzDtULs9d+HoIqF29KBykCKfBuGB8csGPl/ulm/1UTDrjT
N6cyf3c+RDEAe2iWSOFWahIgMd27cY7+nwQwMm44UFs+27KKz/khcLGz4KjN+Zow
PCjOg6kkXIsufjag702Fhe5Bi+TrhrPfYI/QvKKIlmPCum8PKnrJAWQ1zvZlYcBC
bwIDAQAB
-----END PUBLIC KEY-----
`))
	if err != nil {
		t.Error(err)
		return
	}

	encodedData, err := manager.EncryptWithPublicKey(pub, []byte(`ugQiO/004Gt6/7rwf2mg8A==`))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(encodedData))

	//	pri, err := manager.LoadPrivateKeyFromPEM([]byte(`-----BEGIN RSA PRIVATE KEY-----
	//MIIEpQIBAAKCAQEAxHAbmNatnfGk0S7xlXJQ0WI/sY+ZZxJlD6yWfV0zPg6H46mW
	//vJZj6+vJCjTFkNmETmbbmGPemDJ3E6lgCi+e8JEk9461wYGbSFy6R71WQGbCYUv0
	//rXZt4fnk9vTK1Aedjwckiwc42f+aA34WwhhqTTY++S/7nPo4Yfv9OkWzDtULs9d+
	//HoIqF29KBykCKfBuGB8csGPl/ulm/1UTDrjTN6cyf3c+RDEAe2iWSOFWahIgMd27
	//cY7+nwQwMm44UFs+27KKz/khcLGz4KjN+ZowPCjOg6kkXIsufjag702Fhe5Bi+Tr
	//hrPfYI/QvKKIlmPCum8PKnrJAWQ1zvZlYcBCbwIDAQABAoIBAQCQam8PhTxssrte
	//AmofWcSqutViv9SinzZnOJYGol9Kzzn+GK62BMZ/KoBJnZRlslR/o0TsGvgJ4ogC
	//j3II6wuphruruGJNWfCEWY+lsD/Z5vIev82pPTj5elNnb34yNUsTXMfz4iJcunpK
	//+QbYOUTlcO0JG8qalKag+rYnghhq3NHC0t2CKd6lFt8RbIufmLmHqBYhzZ5o7AiX
	//SVOTqtmiZ3d3fEFtgdC1Dku8kaPWpWhS21Y4nGT6UgpmH2fuPbzqcevcWyzi7Czs
	//H1wCqSFYnVv/d9MZsvwSj/GHhPB8zL29gaS3AkCEvAXu3P9fnaiS1kJ+xs5iEV1O
	//Qj4Z/vOxAoGBANYwM+69JA4A1FC7vdHlVmcpwe/4ZDCsDoUNMfKjishhnDLX7s7T
	//pN5G/otRPTZUg8F63V+eeFcbVJn5HyZt/3VYwn1HTz7P7c/GHVSB53XKtl8D3Zeq
	//6HKK2vfmVVQFp8r2hXykbHZ37swh0ESWGmCQyNKOhBHjfieeWeEsWFtpAoGBAOrI
	//2ha0XrSxFjWoTBWRYd1mfTGY0QIy6heDR/AZfa6G3nU86vR28X0rSTUwAgUPKxHm
	//K1p9v5Ovyp3254C1+nO3eJ0WSLg+SGvjhNDzbvkcCunOki3LrPq8JcGjciItzQ/n
	//Ftoxb8cF8eNdicOxPZTkNJgb/clU2aGGOhoECiwXAoGAQDNanajb8caV3U0o7I1N
	//hMajdwaBIYWxJHh8DDqxErcPVr4auqv9sxKcoa3MJ0jV5Vyqlkqtz45FoZFmoOI/
	//vDDKuzpwqmcw5SKBEB+P/WKxn3FNLnTwD8VHNR85XGIFlXSnNmEikbAJR+6quqQz
	//a9Z6G4LUW3hRDBcO4culAGECgYEA4agMHNhdUhQGQaow/mXOBuqzl1DGSfO/lLvE
	//D5ugdXcBJvNW64HKlsBcy3cJ6ezrO3fa4U2NLRg/iNW/KbE+N6v2jBzX5eVO3AtA
	//I0hlt53hS1kUnFlvN0pQi61ZTEpzFj7Icwwi38nx89J6T5DxnEI93pjAspoP1jRZ
	//cZnCYR8CgYEAjoloGSgMxfmzsPP8cQYOl4YWlaJ+YPlkTec5x+o/hZz2StflP5Gn
	//R0hCFKZw5d4O3euZsj5VvLDfAbbeM60k+X/CvMSBrszxeshL1l/TE5q2Np0OalME
	//feqgj1KjEEDruitZvCIS0zk9d4eeDmNyRaeylxmfCunRxJ9z8PBXS1I=
	//-----END RSA PRIVATE KEY-----
	//`))
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}

	//key, err := manager.DecryptWithPrivateKey(pri, encodedData)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}

}

func TestEncryptAES(t *testing.T) {
	manager := encrypt.NewKeyManager(nil, context.Background())
	key := base64.StdEncoding.EncodeToString([]byte("ugQiO/004Gt6/7rwf2mg8A=="))
	//key, err := manager.GenerateAESKey("78138335716970496", 16)
	//if err != nil {
	//	return
	//}
	fmt.Println(key)
	url := []byte("?text = 666")
	data, err := EncrypAES(url, key)
	if err != nil {
		t.Error(err)
		return
	}

	signature := manager.GenerateSignature([]byte(data), key)

	url = []byte(string(url) + "&signature=" + signature)

	fmt.Println(string(url))
	data, err = EncrypAES(url, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)

	decryptAES, err := DecryptAES(key, data)
	if err != nil {
		return
	}
	fmt.Println(string(decryptAES))
}

func EncrypAES(plaintext []byte, keys string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(keys)
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
	return base64.URLEncoding.EncodeToString(ciphertext), nil

}

func DecryptAES(encodedKey string, text string) ([]byte, error) {
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
