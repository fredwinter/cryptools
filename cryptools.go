// Package cryptools has some functions to do basic encryptions using basics ciphers.
// I may not have the best coding skills, but I really enjoyed coding it.
package cryptools

import (
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func NewShift(d int) Cipher {
	if d == 0 || d > 25 || d < -25 {
		return nil
	}
	return Shift{d}
}

type Shift struct {
	distance int
}

// Encode is a Shift method that receives a string as parameter and encodes it using the Shift cipher,
// considering the given distance.
func (c Shift) Encode(s string) string {
	s = strings.ToLower(s)
	var enc string
	for _, r := range s {
		switch {
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

// Decode is a Shift method that receives a string and decodes it using the Shift cipher, considering the given distance
func (c Shift) Decode(s string) string {
	s = strings.ToLower(s)
	var dec string
	for _, r := range s {
		switch {
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

func NewRot13() Cipher {
	return NewShift(13)
}
