package ciphers

/*Cipher is an interface that implements two methods:
Encrypt and Decrypt.*/
type Cipher interface {
	Encrypt([]byte) []byte
	EncryptToString([]byte) string
	Decrypt([]byte) []byte
	DecryptString(string) []byte
}
