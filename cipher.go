package cryptools

/*Cipher is an interface that implements two methods:
Encode and Decode.*/
type Cipher interface {
	Encode([]byte) []byte
	EncodeToString([]byte) string
	Decode([]byte) []byte
	DecodeString(string) []byte
}
