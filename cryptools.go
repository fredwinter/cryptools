// Package cryptools has some functions to do basic encryptions using basics ciphers.
// I may not have the best coding skills, but I really enjoyed coding it.
package cryptools

import (
	"encoding/base64"
	"log"
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// NewShift returns a Shift Cipher.
func NewShift(d int) Cipher {
	if d == 0 || d > 25 || d < -25 {
		return nil
	}
	return Shift{d}
}

type Shift struct {
	distance int
}

// Encode is a Shift method that receives a string as parameter and encodes it using the Shift Cipher,
// considering the given distance.
func (c Shift) Encode(s string) string {
	s = strings.ToLower(s)
	var enc string
	for _, r := range s {
		switch {
		case !unicode.IsLetter(r):
			continue
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

// Decode is a Shift method that receives a string as parameter and decodes it using the Shift Cipher,
// considering the given distance.
func (c Shift) Decode(s string) string {
	s = strings.ToLower(s)
	var dec string
	for _, r := range s {
		switch {
		case !unicode.IsLetter(r):
			continue
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

// NewRot13 returns a new Shift Cipher with 13 as distance.
func NewRot13() Cipher {
	return NewShift(13)
}

// NewCaesar returns a new Shift Cipher with 3 as distance.
func NewCaesar() Cipher {
	return NewShift(3)
}

// Base64 encoding tools
func NewB64() Cipher {
	return B64{}
}

type B64 struct{}

func (b B64) Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (b B64) Decode(s string) string {
	dec, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(dec)
}
