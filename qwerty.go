package qwerty_cipher

import variable_parameter "github.com/golang-infrastructure/go-variable-parameter"

// Encrypt 对明文进行加密
// 支持指定自定义的键盘布局，如果不指定的话默认使用QWERTY
func Encrypt(plaintext string, keyboardLayout ...KeyboardLayout) (string, error) {

	// 未指定键盘布局时使用qwerty布局的键盘
	keyboardLayout = variable_parameter.SetDefaultParam(keyboardLayout, QwertyKeyboardLayout)

	// 参数校验
	if err := keyboardLayout[0].check(); err != nil {
		return "", err
	}

	resultSlice := make([]rune, len(plaintext))
	for index, character := range plaintext {
		character = toUppercaseIfNeed(character)
		// 非字母原样保存
		if character < 'A' || character > 'Z' {
			resultSlice[index] = character
		} else {
			// 字母的话做个映射
			mappingTo := keyboardLayout[0][character-'A']
			resultSlice[index] = mappingTo
		}
	}

	return string(resultSlice), nil
}

// Decrypt 对密文解密，需要传入解密使用的表
func Decrypt(ciphertext string, keyboardLayoutDecryptTable ...KeyboardLayout) (string, error) {

	// 未指定键盘布局时使用qwerty布局的键盘
	keyboardLayoutDecryptTable = variable_parameter.SetDefaultParam(keyboardLayoutDecryptTable, QwertyKeyboardLayoutDecrypt)

	return Encrypt(ciphertext, keyboardLayoutDecryptTable...)
}

// 如果有必要的话则将其转为大写，否则原样返回
func toUppercaseIfNeed(character rune) rune {
	if character >= 'a' && character <= 'z' {
		character -= 32
	}
	return character
}
