package cryptools

/*Cipher is an interface that implements two methods:
Encrypt and Decrypt.*/
type Cipher interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

/*Encoding is an interface that implements two methods:
Encode and Decode.*/
type Encoding interface {
	Encode([]byte) []byte
	Decode([]byte) []byte
}
