package cryptools

import (
	"encoding/hex"
	"log"
)

// NewHex returns a Hex encoding object.
func NewHex() Cipher {
	return Hex{}
}

// Hex is a struct that represents a Hex encoding object.
type Hex struct{}

// Encode is a Hex's method that encodes a slice of bytes using the hex encoding
// and returns it in the form of a slice of bytes.
func (h Hex) Encode(s string) string {
	return hex.EncodeToString([]byte(s))
}

// Decode is a Hex's method that decodes a slice of bytes using the hex encoding
// and returns it in the form of a slice of bytes.
func (h Hex) Decode(s string) string {
	dec, err := hex.DecodeString(s)
	if err != nil {
		log.Fatalln(err)
	}
	return string(dec)
}
