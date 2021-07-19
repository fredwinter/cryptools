package cryptools

import (
	"math/big"
)

// NewBigInt returns a BigInt encoding object.
func NewBigInt() Encoding {
	return BigInt{}
}

// BigInt is a struct that represents a BigInt encoding object.
type BigInt struct{}

/* Encode is a BigInt's method that encodes a slice of bytes using the BigInt encoding and
returns it in the form of a slice of bytes. */
func (b BigInt) Encode(m []byte) []byte {
	sBigInt := new(big.Int).SetBytes(m)
	return NewHex().Encode(sBigInt.Bytes())
}

// EncodeToString returns a string using the Encoding's Encode method.
func (b BigInt) EncodeToString(m []byte) string {
	return string(b.Encode(m))
}

/* Decode is a BigInt's method that decodes a slice of bytes using the BigInt encoding and
returns it in the form of a slice of bytes.*/
func (b BigInt) Decode(m []byte) []byte {
	BigInt := new(big.Int).SetBytes(m)
	return NewHex().Decode(BigInt.Bytes())
}

// DecodeString decodes a string using the Encoding's Decode method.
func (b BigInt) DecodeString(m string) []byte {
	return b.Decode([]byte(m))
}
