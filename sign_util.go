package cex

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

type Signer func(payload, key string) []byte

func SignByHmacSHA256(payload, key string) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(payload))
	return h.Sum(nil)
}

func SignByHmacSHA512(payload, key string) []byte {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(payload))
	return h.Sum(nil)
}

func SignByHmacSHA256ToHex(payload, key string) string {
	return hex.EncodeToString(SignByHmacSHA256(payload, key))
}

func SignByHmacSHA512ToHex(payload, key string) string {
	return hex.EncodeToString(SignByHmacSHA512(payload, key))
}

func SignByHmacSHA256ToBase64(payload, key string) string {
	return base64.StdEncoding.EncodeToString(SignByHmacSHA256(payload, key))
}

func Sha512ToHex(payload string) string {
	res := sha512.Sum512([]byte(payload))
	return hex.EncodeToString(res[:])
}
