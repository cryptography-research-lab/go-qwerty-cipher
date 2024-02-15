package qwerty_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plaintext := "HelloWorld"
	encrypt, err := Encrypt(plaintext)
	assert.Nil(t, err)
	assert.Equal(t, "ItssgVgksr", encrypt)

	decrypt, err := Decrypt(encrypt)
	assert.Nil(t, err)
	assert.Equal(t, plaintext, decrypt)
}
