package qwerty_cipher

// KeyboardLayout 表示一个键盘布局，这个布局可以是任意布局，只要是确定了映射关系就可以了
type KeyboardLayout []rune

// 检查当前的键盘布局映射表是否合法
func (x KeyboardLayout) check() error {

	// 如果长度都不对，则没必要继续下去了
	if len(x) != 26 {
		return ErrKeyboardLayoutCharacterInvalid
	}

	// 然后统计每个字符的出现次数
	characterCountSlice := make([]int, 26)
	for _, character := range x {
		character = toUppercaseIfNeed(character)
		mappingToIndex := character - 'A'
		if mappingToIndex >= 26 {
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
	for index, toCharacter := range x {
		fromCharacter := rune('A' + index)
		toIndex := toCharacter - 'A'
		decryptTable[toIndex] = fromCharacter
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
