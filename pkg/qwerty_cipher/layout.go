package qwerty_cipher

// KeyboardLayout 表示一个键盘布局，这个布局可以是任意布局，只要是确定了映射关系就可以了
// 具体的例子详见 QwertyKeyboardLayout 实现
type KeyboardLayout []rune

// Check 检查当前的键盘布局映射表是否合法
func (x KeyboardLayout) Check() error {

	// 如果长度都不对，则没必要继续下去了
	if len(x) != 26 {
		return ErrKeyboardLayoutCharacterInvalid
	}

	// 然后统计每个字符的出现次数，每个字符都必须出现，并且只能出现一次
	characterCountSlice := make([]int, 26)
	for _, character := range x {
		mappingToIndex := toMappingIndex(character)
		// 给定的字符必须在有效范围
		if mappingToIndex >= 26 || mappingToIndex < 0 {
			return ErrKeyboardLayoutCharacterInvalid
		}
		characterCountSlice[mappingToIndex]++
	}
	for _, count := range characterCountSlice {
		if count != 1 {
			return ErrKeyboardLayoutCharacterInvalid
		}
	}

	return nil
}

// TransformToDecrypt 将加密表转换为解密表
func (x KeyboardLayout) TransformToDecrypt() KeyboardLayout {
	decryptTable := make([]rune, len(x))
	for index, fromCharacter := range x {
		toCharacter := rune('A' + index)
		fromIndex := toMappingIndex(fromCharacter)
		decryptTable[fromIndex] = toCharacter
	}
	return decryptTable
}

// ------------------------------------------------ ---------------------------------------------------------------------

// QwertyKeyboardLayout qwerty布局键盘的加密映射表
var QwertyKeyboardLayout KeyboardLayout

// QwertyKeyboardLayoutDecrypt qwerty布局键盘的解密映射表
var QwertyKeyboardLayoutDecrypt KeyboardLayout

func init() {

	QwertyKeyboardLayout = make([]rune, 26)
	layout := "QWERTYUIOPASDFGHJKLZXCVBNM"
	for index, character := range layout {
		QwertyKeyboardLayout[index] = character
	}

	QwertyKeyboardLayoutDecrypt = QwertyKeyboardLayout.TransformToDecrypt()

}

// ------------------------------------------------ ---------------------------------------------------------------------

// QwertzKeyboardLayout qwertz布局键盘的加密映射表
// 维基百科： https://zh.wikipedia.org/zh-cn/QWERTZ%E9%8D%B5%E7%9B%A4
var QwertzKeyboardLayout KeyboardLayout

// QwertzKeyboardLayoutDecrypt qwertz布局键盘的解密映射表
var QwertzKeyboardLayoutDecrypt KeyboardLayout

func init() {

	QwertzKeyboardLayout = make([]rune, 26)
	layout := "QWERTZUIOPASDFGHJKLYXCVBNM"
	for index, character := range layout {
		QwertzKeyboardLayout[index] = character
	}

	QwertzKeyboardLayoutDecrypt = QwertzKeyboardLayout.TransformToDecrypt()

}

// ------------------------------------------------ ---------------------------------------------------------------------

// AzertyKeyboardLayout Azerty布局键盘的加密映射表
var AzertyKeyboardLayout KeyboardLayout

// AzertyKeyboardLayoutDecrypt qwertz布局键盘的解密映射表
var AzertyKeyboardLayoutDecrypt KeyboardLayout

func init() {

	AzertyKeyboardLayout = make([]rune, 26)
	layout := "AZERTYUIOPQSDFGHJKLMWXCVBN"
	for index, character := range layout {
		AzertyKeyboardLayout[index] = character
	}

	AzertyKeyboardLayoutDecrypt = AzertyKeyboardLayout.TransformToDecrypt()

}

// ------------------------------------------------ ---------------------------------------------------------------------
