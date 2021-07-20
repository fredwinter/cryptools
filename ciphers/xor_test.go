package ciphers

import (
	"bytes"
	"testing"
)

type testCase struct {
	dec, key, enc []byte
}

var encryptDecryptCases = []testCase{
	{[]byte("label"), []byte{13}, []byte("aloha")},
	{[]byte{69, 108, 118, 105, 115}, []byte{169, 207, 195, 134, 172}, []byte{236, 163, 181, 239, 223}},
	{[]byte("longplaintext"), []byte("key"), []byte{7, 10, 23, 12, 21, 21, 10, 12, 23, 31, 0, 1, 31}},
}

func TestNewXOR(t *testing.T) {
	k := []byte("testing the NewXOR Function")
	got := NewXOR(k)
	expected := XOR{k}
	if bytes.Compare(got.(XOR).key, expected.key) != 0 {
		t.Errorf("error. NewXOR(%v) got: %v, expected: %s", k, got.(XOR).key, expected.key)
	}
}

func TestEncrypt(t *testing.T) {
	for _, tc := range encryptDecryptCases {
		got := NewXOR(tc.key).Encrypt(tc.dec)
		if bytes.Compare(got, tc.enc) != 0 {
			t.Errorf("error. NewXOR(%v).Encrypt(%s) got %s, expected %s", tc.key, tc.dec, got, tc.enc)
		}
	}
}

func TestEncryptToString(t *testing.T) {
	for _, tc := range encryptDecryptCases {
		got := NewXOR(tc.key).EncryptToString(tc.dec)
		if got != string(tc.enc) {
			t.Errorf("error. NewXOR(%v).EncryptToString(%s) got %s, expected %s", tc.key, tc.dec, got, tc.enc)
		}
	}
}

func TestDecrypt(t *testing.T) {
	for _, tc := range encryptDecryptCases {
		got := NewXOR(tc.key).Decrypt(tc.enc)
		if bytes.Compare(tc.dec, got) != 0 {
			t.Errorf("error. NewXOR(%v).Decrypt(%s) got %s, expected %s", tc.key, tc.enc, got, tc.dec)
		}
	}
}

func TestDecryptString(t *testing.T) {
	for _, tc := range encryptDecryptCases {
		got := NewXOR(tc.key).DecryptString(string(tc.enc))
		if bytes.Compare(got, tc.dec) != 0 {
			t.Errorf("error. NewXOR(%v).Decrypt(%s) got %s, expected %s", tc.key, tc.enc, got, tc.dec)
		}
	}
}
