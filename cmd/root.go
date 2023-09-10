package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var className string

var rootCmd = &cobra.Command{
	Use:     "convert",
	Short:   "Convert: Quickly turn JSON into PHP Data Objects",
	Long:    "Convert: Quickly turn JSON into PHP Data Objects",
	Version: "0.0.1",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to run command: %v\n\n", err)
		os.Exit(1)
	}
}
