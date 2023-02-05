package qwerty_cipher

import "errors"

// ErrKeyboardLayoutCharacterInvalid 传入的键盘布局表中
var ErrKeyboardLayoutCharacterInvalid = errors.New("all A-Z must appear in the keyboard layout table exactly once")
