package qwerty_cipher

import "errors"

// ErrKeyboardLayoutTableCharacterInvalid 传入的键盘布局表中
var ErrKeyboardLayoutTableCharacterInvalid = errors.New("all A-Z must appear in the keyboard layout table exactly once")
