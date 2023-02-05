package main

import (
	"fmt"
	qwerty_cipher "github.com/cryptography-research-lab/go-qwerty-cipher"
)

func main() {

	// 对明文加密
	plaintext := "helloworld"
	encrypt, err := qwerty_cipher.Encrypt(plaintext)
	if err != nil {
		fmt.Println("加密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("加密结果： " + encrypt) // Output: 加密结果： ITSSGVGKSR

	// 对密文解密
	decrypt, err := qwerty_cipher.Decrypt(encrypt)
	if err != nil {
		fmt.Println("解密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decrypt) // Output: 解密结果： HELLOWORLD

}
