package ciphers

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

/* Encrypt is a Shift's method that encrypts a slice of bytes and returns it
using the Shift Cipher, considering the given distance, and returns it as
a slice of bytes. */
func (c Shift) Encrypt(m []byte) []byte {
	s := strings.ToLower(string(m))
	var enc []byte
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

// EncryptToString returns a string using the Cipher's Encrypt method.
func (c Shift) EncryptToString(m []byte) string {
	return string(c.Encrypt(m))
}

/* Decrypt is a Shift's method that decrypts a slice of bytes using the
Shift Cipher, considering the given distance, and returns it as a
a slice of bytes. */
func (c Shift) Decrypt(m []byte) []byte {
	s := strings.ToLower(string(m))
	var dec []byte
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

// DecryptString decrypts a string using the Cipher's Decrypt method.
func (c Shift) DecryptString(m string) []byte {
	return c.Decrypt([]byte(m))
}
