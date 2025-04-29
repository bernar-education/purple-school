package encrypter

import "os"

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("ENCRYPT_KEY")
	if key == "" {
		panic("Can not find encrypt key in env variables")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainStr string) string {
	return ""
}

func (enc *Encrypter) Decrypt(encryptedStr string) string {
	return ""
}
