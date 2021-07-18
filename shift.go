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
a slice of bytes.*/
func (c Shift) Encode(s string) string {
	s = strings.ToLower(s)
	var enc string
	for _, r := range s {
		switch {
		case !unicode.IsLetter(r):
			enc += string(r)
		case r+rune(c.distance) > 'z':
			enc += string(r + rune(c.distance) - 26)
		case r+rune(c.distance) < 'a':
			enc += string(r + rune(c.distance) + 26)
		default:
			enc += string(r + rune(c.distance))
		}
	}
	return enc
}

/* Decode is a Shift's method that decodes a slice of bytes using the
Shift Cipher, considering the given distance, and returns it as a
a slice of bytes.*/
func (c Shift) Decode(s string) string {
	s = strings.ToLower(s)
	var dec string
	for _, r := range s {
		switch {
		case !unicode.IsLetter(r):
			dec += string(r)
		case r-rune(c.distance) < 'a':
			dec += string(r - rune(c.distance) + 26)
		case r-rune(c.distance) > 'z':
			dec += string(r - rune(c.distance) - 26)
		default:
			dec += string(r - rune(c.distance))
		}
	}
	return dec
}
