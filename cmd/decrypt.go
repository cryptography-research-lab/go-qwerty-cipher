package cmd

import (
	"fmt"
	"github.com/cryptography-research-lab/go-qwerty-cipher/pkg/qwerty_cipher"
	"github.com/spf13/cobra"
)

// 使用qwerty算法解密
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "",
	Long:  `Decrypt with qwerty cipher`,
	Run: func(cmd *cobra.Command, args []string) {

		if !validateDecryptArgs(args) {
			return
		}

		// 允许一次指定多个字符串解密
		for _, s := range args {
			encrypt, err := qwerty_cipher.Decrypt(s)
			if err != nil {
				fmt.Println(s+" decrypt error: ", err.Error())
				continue
			}
			fmt.Println(s + " --> " + encrypt)
		}

	},
	TraverseChildren: true,
}

// 检查参数
func validateDecryptArgs(args []string) bool {

	if len(args) == 0 {
		fmt.Println("Must input decrypt string, example: ./qwerty-cipher decrypt foo bar")
		return false
	}

	return true
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
