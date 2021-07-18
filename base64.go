package cryptools

import (
	"encoding/base64"
)

// NewB64 returns a B64 encoding object.
func NewB64() Cipher {
	return B64{}
}

// B64 is a struct that represents a Base64 encoding object.
type B64 struct{}

// Encode is a B64's method that encodes a slice of bytes using the base64 encoding and
// returns it in the form of a slice of bytes.
func (b B64) Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Decode is a B64's method that decodes a slice of bytes using the base64 encoding and
// returns it in the form of a slice of bytes.
func (b B64) Decode(s string) string {
	dec, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(dec)
}
