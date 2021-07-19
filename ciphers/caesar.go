package ciphers

// NewCaesar returns a new Shift Cipher with 3 as distance.
func NewCaesar() Cipher {
	return NewShift(3)
}
