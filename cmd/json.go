package cmd

import (
	"fmt"

	"github.com/juststeveking/converter/pkg/converter"
	"github.com/juststeveking/converter/pkg/json"
	"github.com/spf13/cobra"
)

var jsonInput string

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "JSON: Convert JSON into PHP Data Object",
	Long:  "JSON: Convert JSON into PHP Data Object",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := json.ParseJSONToMap(jsonInput)
		if err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return
		}
		dto := converter.GeneratePHPDTO(data, className)
		fmt.Println(dto)
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	// Adding flags for our command
	jsonCmd.Flags().StringVarP(&className, "class", "c", "", "Name of the DTO class (required)")
	jsonCmd.Flags().StringVarP(&jsonInput, "json", "j", "", "JSON input (required)")
	jsonCmd.MarkFlagRequired("class")
	jsonCmd.MarkFlagRequired("json")
}
