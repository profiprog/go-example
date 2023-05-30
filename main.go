package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello World example for Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello %s!\n", FlagName)
	},
}

var CmdVersion = &cobra.Command{
	Use:   "proc",
	Short: "Print proc name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Proc name: ", os.Args[0])
	},
}

func init() {
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Name to say hello to",
	)
}

func main() {
	Cmd.Execute()
}
