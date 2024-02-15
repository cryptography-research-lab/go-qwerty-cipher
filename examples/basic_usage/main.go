package main

import (
	"fmt"
	"github.com/cryptography-research-lab/go-qwerty-cipher/pkg/qwerty_cipher"
)

func main() {

	// 对明文加密
	plaintext := "HelloWorld"
	encrypt, err := qwerty_cipher.Encrypt(plaintext)
	if err != nil {
		fmt.Println("加密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("加密结果： " + encrypt) // Output: 加密结果： ItssgVgksr

	// 对密文解密
	decrypt, err := qwerty_cipher.Decrypt(encrypt)
	if err != nil {
		fmt.Println("解密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decrypt) // Output: 解密结果： HelloWorld

}
