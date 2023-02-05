package main

import (
	"fmt"
	qwerty_cipher "github.com/cryptography-research-lab/go-qwerty-cipher"
)

func main() {

	keywordLayout := qwerty_cipher.QwertzKeyboardLayout

	// 对明文加密
	plaintext := "helloworld"
	encrypt, err := qwerty_cipher.Encrypt(plaintext, keywordLayout)
	if err != nil {
		fmt.Println("加密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("加密结果： " + encrypt) // Output: 加密结果： ITSSGVGKSR

	// 对密文解密，解密的时候要将布局转换为解密使用的形式，如果是多次调用，则应该保存转换结果避免重复转换
	decrypt, err := qwerty_cipher.Decrypt(encrypt, keywordLayout.TransformToDecrypt())
	if err != nil {
		fmt.Println("解密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decrypt) // Output: 解密结果： HELLOWORLD

}
