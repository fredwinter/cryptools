package cryptools

import (
	"log"
	"math/big"
)

// NewBigInt returns a BigInt Cipher.
func NewBigInt() Cipher {
	return BigInt{}
}

// BigInt is a struct that represents a BigInt Cipher.
type BigInt struct{}

// Encode is a BigInt's method that encodes a slice of bytes using the BigInt Cipher and
// returns it in the form of a slice of bytes.
func (b BigInt) Encode(s string) string {
	sBigInt, ok := new(big.Int).SetString(s, 10)
	if !ok {
		log.Fatalf("error encoding %s", s)
		return ""
	}
	return NewHex().Encode(sBigInt.String())
}

/*Decode is a BigInt's method that decodes a slice of bytes using the BigInt Cipher and
returns it in the form of a slice of bytes.*/
func (b BigInt) Decode(s string) string {
	sBigIntHex, ok := new(big.Int).SetString(s, 16)
	if !ok {
		log.Fatalf("error decoding %s", s)
		return ""
	}
	sBigInt := NewHex().Decode(sBigIntHex.Text(16))
	return string(sBigInt)
}
