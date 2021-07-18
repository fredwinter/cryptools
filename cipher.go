package cryptools

/*Cipher is an interface that implements two methods:
Encode and Decode.*/
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
