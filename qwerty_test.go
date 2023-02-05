package qwerty_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plaintext := "helloworld"
	encrypt, err := Encrypt(plaintext)
	assert.Nil(t, err)
	assert.Equal(t, "ITSSGVGKSR", encrypt)

	decrypt, err := Decrypt(encrypt)
	assert.Nil(t, err)
	assert.Equal(t, "HELLOWORLD", decrypt)
}
