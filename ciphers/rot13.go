package ciphers

// NewRot13 returns a new Shift Cipher with 13 as distance.
func NewRot13() Cipher {
	return NewShift(13)
}
