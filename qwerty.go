package qwerty_cipher

import variable_parameter "github.com/golang-infrastructure/go-variable-parameter"

// Encrypt 对明文进行加密
// 支持指定自定义的键盘布局，如果不指定的话默认使用QWERTY
func Encrypt(plaintext string, table ...Table) string {

	// 未指定键盘布局时使用qwerty布局的键盘
	table = variable_parameter.SetDefaultParam(table, QwertyTable)

	resultSlice := make([]rune, len(plaintext))
	for index, character := range plaintext {
		character = toUppercaseIfNeed(character)
		// 非字母原样保存
		if character < 'A' || character > 'Z' {
			resultSlice[index] = character
		} else {
			// 字母的话做个映射
			mappingTo := table[0][character-'A']
			resultSlice[index] = mappingTo
		}
	}

	return string(resultSlice)
}

// Decrypt 对密文解密
func Decrypt(ciphertext string, table ...Table) string {

	// 未指定键盘布局时使用qwerty布局的键盘
	table = variable_parameter.SetDefaultParam(table, QwertyTable)

	// 将矩阵转换为适合解密使用的形式，然后按照加密流程走一遍就可以了
	table[0] = table[0].TransformToDecrypt()

	return Encrypt(ciphertext, table...)
}

func toUppercaseIfNeed(character rune) rune {
	if character >= 'a' && character <= 'z' {
		character -= 32
	}
	return character
}
