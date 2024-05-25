package cmd

import (
	"fmt"
	"github.com/cryptography-research-lab/go-qwerty-cipher/pkg/qwerty_cipher"
	"github.com/spf13/cobra"
)

// 使用qwerty算法加密
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "",
	Long:  `Encrypt with qwerty cipher`,
	Run: func(cmd *cobra.Command, args []string) {

		if !validateEncryptArgs(args) {
			return
		}

		// 允许一次指定多个字符串加密
		for _, s := range args {
			encrypt, err := qwerty_cipher.Encrypt(s)
			if err != nil {
				fmt.Println(s+" encrypt error: ", err.Error())
				continue
			}
			fmt.Println(s + " --> " + encrypt)
		}

	},
	TraverseChildren: true,
}

// 检查参数
func validateEncryptArgs(args []string) bool {

	if len(args) == 0 {
		fmt.Println("Must input encrypt string, example: ./qwerty-cipher encrypt foo bar")
		return false
	}

	return true
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
