package cryptools

/*Cipher is an interface that implements two methods:
Encrypt and Decrypt.*/
type Cipher interface {
	Encrypt([]byte) []byte
	EncryptToString([]byte) string
	Decrypt([]byte) []byte
	DecryptString(string) []byte
}

/*Encoding is an interface that implements two methods:
Encode and Decode.*/
type Encoding interface {
	Encode([]byte) []byte
	EncodeToString([]byte) string
	Decode([]byte) []byte
	DecodeString(string) []byte
}
