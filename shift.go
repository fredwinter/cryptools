package cryptools

import (
	"strings"
	"unicode"
)

// NewShift returns a Shift Cipher.
func NewShift(d int) Cipher {
	if d == 0 || d > 25 || d < -25 {
		return nil
	}
	return Shift{d}
}

// Shift is a struct that represents a Shift Cipher.
type Shift struct {
	distance int
}

/* Encode is a Shift's method that encodes a slice of bytes and returns it
using the Shift Cipher, considering the given distance, and returns it as
a slice of bytes. */
func (c Shift) Encode(m []byte) []byte {
	s := strings.ToLower(string(m))
	enc := make([]byte, len(m))
	for _, r := range s {
		switch {
		case !unicode.IsLetter(r):
			enc = append(enc, byte(r))
		case r+rune(c.distance) > 'z':
			enc = append(enc, byte(r+rune(c.distance)-26))
		case r+rune(c.distance) < 'a':
			enc = append(enc, byte(r+rune(c.distance)+26))
		default:
			enc = append(enc, byte(r+rune(c.distance)))
		}
	}
	return enc
}

// EncodeToString returns a string using the Cipher's Encode method.
func (c Shift) EncodeToString(m []byte) string {
	return string(c.Encode(m))
}

/* Decode is a Shift's method that decodes a slice of bytes using the
Shift Cipher, considering the given distance, and returns it as a
a slice of bytes. */
func (c Shift) Decode(m []byte) []byte {
	s := strings.ToLower(string(m))
	dec := make([]byte, len(m))
	for _, r := range s {
		switch {
		case !unicode.IsLetter(r):
			dec = append(dec, byte(r))
		case r-rune(c.distance) < 'a':
			dec = append(dec, byte(r-rune(c.distance)+26))
		case r-rune(c.distance) > 'z':
			dec = append(dec, byte(r-rune(c.distance)-26))
		default:
			dec = append(dec, byte(r-rune(c.distance)))
		}
	}
	return dec
}

// DecodeString decodes a string using the Cipher's Decode method.
func (c Shift) DecodeString(m string) []byte {
	return c.Decode([]byte(m))
}
