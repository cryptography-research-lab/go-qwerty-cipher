package qwerty_cipher

type Table []rune

// TransformToDecrypt 转为适合解密使用的样子
func (x Table) TransformToDecrypt() Table {
	newTable := make([]rune, len(x))
	for index, character := range x {
		from := rune('A' + index)
		to := character
		newTable[to-'A'] = from
	}
	return newTable
}

// ------------------------------------------------ ---------------------------------------------------------------------

// QwertyTable qwerty布局键盘的加密映射表
var QwertyTable Table

func init() {
	QwertyTable = make([]rune, 26)
	layout := "QWERTYUIOPASDFGHJKLZXCVBNM"
	for index, character := range layout {
		QwertyTable[index] = character
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------
