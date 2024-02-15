package qwerty_cipher

// 判断给定的rune是否是英文字母
func isLetter(c rune) bool {
	return isLowercaseLetter(c) || isUppercaseLetter(c)
}

// 判断是否是小写字母
func isLowercaseLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}

// 判断是否是大写字母
func isUppercaseLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

// 把字母转换为大写
func toUppercase(c rune) rune {
	if isLowercaseLetter(c) {
		c -= 32
	}
	return c
}

// 把字母转换为小写
func toLowercase(c rune) rune {
	if isUppercaseLetter(c) {
		c += 32
	}
	return c
}

// b跟随a的大小写特性
func followUppercaseOrLowercase(a, b rune) rune {
	if isUppercaseLetter(a) {
		b = toUppercase(b)
	} else if isLowercaseLetter(a) {
		b = toLowercase(b)
	}
	return b
}

// 把字符转换为键盘布局切片对应的下标，比如a和A转换为0，b和B转换为1，以此类推
func toMappingIndex(c rune) int {
	if isLowercaseLetter(c) {
		return int(c - 'a')
	} else if isUppercaseLetter(c) {
		return int(c - 'A')
	} else {
		return -1
	}
}

