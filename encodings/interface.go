package encodings

/*Encoding is an interface that implements two methods:
Encode and Decode.*/
type Encoding interface {
	Encode([]byte) []byte
	EncodeToString([]byte) string
	Decode([]byte) []byte
	DecodeString(string) []byte
}
