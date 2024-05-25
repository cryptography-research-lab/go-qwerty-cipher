package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "qwerty-cipher",
	Short: "qwerty-cipher",
	Long:  `qwerty cipher`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("qwerty cipher tool")
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
