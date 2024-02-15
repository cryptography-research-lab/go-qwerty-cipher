package qwerty_cipher

import "errors"

// ErrKeyboardLayoutCharacterInvalid 传入的键盘布局表中指定的字母不合法，必须是A-Z每个字符都出现并且恰好出现一次
var ErrKeyboardLayoutCharacterInvalid = errors.New("all A-Z must appear in the keyboard layout table exactly once")


