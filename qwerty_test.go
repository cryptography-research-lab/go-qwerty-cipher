package qwerty_cipher

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plaintext := "helloworld"
	encrypt := Encrypt(plaintext)
	fmt.Println(encrypt)

	decrypt := Decrypt(encrypt)
	fmt.Println(decrypt)
}
