package cmd

import (
	"fmt"
	"github.com/cryptography-research-lab/go-qwerty-cipher/pkg/qwerty_cipher"
	"github.com/spf13/cobra"
	"strings"
)

// 可以指定键盘布局
var keyboard string

// DefaultKeyboardLayout 默认的键盘布局为qwerty布局
var DefaultKeyboardLayout = qwerty_cipher.QwertyKeyboardLayout

// 要加密的明文
var plaintext string

func init() {

	encryptCmd.Flags().StringVarP(&keyboard, "keyboard", "k", "", "custom keyboard")
	encryptCmd.Flags().StringVarP(&plaintext, "plaintext", "t", "", "plaintext to encrypt")

	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		// 解析键盘布局
		keyboardLayout := parseKeyboardLayout(keyboard)

		// 参数检查
		if err := checkEncryptParams(keyboardLayout, plaintext); err != nil {
			return err
		}

		plaintext, err := qwerty_cipher.Encrypt(plaintext, keyboardLayout)
		if err != nil {
			return err
		}
		fmt.Println("Encrypt Result: " + plaintext)
		return nil
	},
}

// 解析传入的键盘布局参数，如果没有传的话则使用默认的
func parseKeyboardLayout(keyboard string) qwerty_cipher.KeyboardLayout {

	// 如果没有指定的话，则使用默认的键盘布局
	if keyboard == "" {
		return DefaultKeyboardLayout
	}

	// 看看是不是预置的几个布局
	switch strings.ToLower(keyboard) {
	case "qwerty":
		return qwerty_cipher.QwertyKeyboardLayout
	case "qwertz":
		return qwerty_cipher.QwertzKeyboardLayout
	case "azerty":
		return qwerty_cipher.AzertyKeyboardLayout
	default:
		// 指定的不是预置的键盘布局，说明是自定义的键盘布局，则直接转为键盘布局使用
		return qwerty_cipher.KeyboardLayout(keyboard)
	}

}

// 检查加密所需的参数是否合法
func checkEncryptParams(keyboardLayout qwerty_cipher.KeyboardLayout, plaintext string) error {

	// 要加密的明文不允许为空
	if plaintext == "" {
		return fmt.Errorf("plaintext must not empty")
	}

	// 键盘布局可以不指定，仅当指定的时候才检查布局是否合法
	if err := keyboardLayout.Check(); err != nil {
		return fmt.Errorf("keyboard layout %s is not ok: %s", string(keyboardLayout), err.Error())
	}

	return nil
}
