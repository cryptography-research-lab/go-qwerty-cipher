package qwerty_cipher

// KeyboardLayoutTable 表示一个键盘布局表，这个布局可以是任意布局，只要是确定了映射关系就可以了
type KeyboardLayoutTable []rune

// 检查当前的键盘布局映射表是否合法
func (x KeyboardLayoutTable) check() error {

	// 如果长度都不对，则没必要继续下去了
	if len(x) != 26 {
		return ErrKeyboardLayoutTableCharacterInvalid
	}

	// 然后统计每个字符的出现次数
	characterCountSlice := make([]int, 26)
	for _, character := range x {
		character = toUppercaseIfNeed(character)
		mappingToIndex := character - 'A'
		if mappingToIndex >= 26 {
			return ErrKeyboardLayoutTableCharacterInvalid
		}
		characterCountSlice[mappingToIndex]++
	}
	for _, count := range characterCountSlice {
		if count != 1 {
			return ErrKeyboardLayoutTableCharacterInvalid
		}
	}

	return nil
}

// TransformToDecryptTable 将加密表转换为解密表
func (x KeyboardLayoutTable) TransformToDecryptTable() KeyboardLayoutTable {
	newTable := make([]rune, len(x))
	for index, character := range x {
		from := rune('A' + index)
		to := character
		newTable[to-'A'] = from
	}
	return newTable
}

// ------------------------------------------------ ---------------------------------------------------------------------

// QwertyKeyboardLayoutTable qwerty布局键盘的加密映射表
var QwertyKeyboardLayoutTable KeyboardLayoutTable

// QwertyKeyboardLayoutDecryptTable qwerty布局键盘的解密映射表
var QwertyKeyboardLayoutDecryptTable KeyboardLayoutTable

func init() {

	QwertyKeyboardLayoutTable = make([]rune, 26)
	layout := "QWERTYUIOPASDFGHJKLZXCVBNM"
	for index, character := range layout {
		QwertyKeyboardLayoutTable[index] = character
	}

	QwertyKeyboardLayoutDecryptTable = QwertyKeyboardLayoutTable.TransformToDecryptTable()

}

// ------------------------------------------------ ---------------------------------------------------------------------
