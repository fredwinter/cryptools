package cryptools

import (
	"encoding/hex"
)

// NewHex returns a Hex encoding object.
func NewHex() Encoding {
	return Hex{}
}

// Hex is a struct that represents a Hex encoding object.
type Hex struct{}

/* Encode is a Hex's method that encodes a slice of bytes using the hex encoding
and returns it in the form of a slice of bytes. */
func (h Hex) Encode(m []byte) []byte {
	enc := make([]byte, hex.EncodedLen(len(m)))
	hex.Encode(enc, m)
	return enc
}

// EncodeToString returns a string using the Encoding's Encode method.
func (h Hex) EncodeToString(m []byte) string {
	return string(h.Encode(m))
}

/* Decode is a Hex's method that decodes a slice of bytes using the hex encoding
and returns it in the form of a slice of bytes. */
func (h Hex) Decode(m []byte) []byte {
	dec := make([]byte, hex.DecodedLen(len(m)))
	hex.Decode(dec, m)
	return dec
}

// DecodeString decodes a string using the Encoding's Decode method.
func (h Hex) DecodeString(m string) []byte {
	return h.Decode([]byte(m))
}
