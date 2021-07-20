package ciphers

//NewXOR returns a XOR Cipher.
func NewXOR(k []byte) Cipher {
	return XOR{k}
}

//XOR is a struct that represents a XOR Cipher.
type XOR struct {
	key []byte
}

/*Encrypt is a XOR's method that encrypts a slice of bytes using the XOR Cipher
and returns it as a slice of bytes.*/
func (x XOR) Encrypt(m []byte) []byte {
	var enc []byte
	for i, b := range m {
		enc = append(enc, b^x.key[i%len(x.key)])
	}
	return enc
}

//EncryptToString returns a string using the Cipher's Encrypt method.
func (x XOR) EncryptToString(m []byte) string {
	return string(x.Encrypt(m))
}

/*Decrypt is a XOR's method that decrypts a slice of bytes using the XOR Cipher
and returns it as a slice of bytes.*/
func (x XOR) Decrypt(m []byte) []byte {
	return x.Encrypt(m)
}

//DecryptString decrypts a string using the Cipher's Decrypt method.
func (x XOR) DecryptString(m string) []byte {
	return x.Decrypt([]byte(m))
}
