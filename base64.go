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

/* Encode is a B64's method that encodes a slice of bytes using the base64 encoding and
returns it in the form of a slice of bytes. */
func (b B64) Encode(m []byte) []byte {
	enc := make([]byte, len(m))
	base64.StdEncoding.Encode(enc, m)
	return enc
}

// EncodeToString returns a string using the Cipher's Encode method.
func (b B64) EncodeToString(m []byte) string {
	return string(b.Encode(m))
}

/* Decode is a B64's method that decodes a slice of bytes using the base64 encoding and
returns it in the form of a slice of bytes. */
func (b B64) Decode(m []byte) []byte {
	dec := make([]byte, len(m))
	base64.StdEncoding.Decode(dec, m)
	return dec
}

// DecodeString decodes a string using the Cipher's Decode method.
func (b B64) DecodeString(m string) []byte {
	return b.Decode([]byte(m))
}
