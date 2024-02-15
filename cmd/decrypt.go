package cmd

import (
	"fmt"
	"github.com/cryptography-research-lab/go-qwerty-cipher/pkg/qwerty_cipher"
	"github.com/spf13/cobra"
)

var ciphertext string

func init() {

	decryptCmd.Flags().StringVarP(&keyboard, "keyboard", "k", "", "custom keyboard")
	decryptCmd.Flags().StringVarP(&ciphertext, "ciphertext", "c", "", "ciphertext to decrypt")

	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		keyboardLayout := parseKeyboardLayout(keyboard)

		if err := checkDecryptParams(keyboardLayout, ciphertext); err != nil {
			return err
		}

		plaintext, err := qwerty_cipher.Decrypt(ciphertext, keyboardLayout.TransformToDecrypt())
		if err != nil {
			return err
		}
		fmt.Println("Decrypt Result: " + plaintext)
		return nil
	},
}

// 检查加密所需的参数是否合法
func checkDecryptParams(keyboardLayout qwerty_cipher.KeyboardLayout, ciphertext string) error {

	// 密文不允许为空
	if ciphertext == "" {
		return fmt.Errorf("ciphertext must not empty")
	}

	// 键盘布局可以不指定，仅当指定的时候才检查布局是否合法
	if err := keyboardLayout.Check(); err != nil {
		return fmt.Errorf("keyboard layout %s is not ok: %s", string(keyboardLayout), err.Error())
	}

	return nil
}
