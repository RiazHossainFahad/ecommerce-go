package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	b64Header := convertToBase64(byteArrHeader)

	byteArrPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	b64Payload := convertToBase64(byteArrPayload)

	byteMessage := []byte(b64Header + "." + b64Payload)
	byteSecret := []byte(secret)

	hash := hmac.New(sha256.New, byteSecret)
	hash.Write(byteMessage)
	signature := hash.Sum(nil)
	b64Signature := convertToBase64(signature)

	jwt := b64Header + "." + b64Payload + "." + b64Signature

	return jwt, nil
}

func convertToBase64(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
