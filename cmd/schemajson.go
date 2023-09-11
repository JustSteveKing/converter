package cmd

import (
	"fmt"

	"github.com/juststeveking/converter/pkg/converter"
	"github.com/spf13/cobra"
)

var schemaInput string

var schemaJsonCmd = &cobra.Command{
	Use:   "schema-json",
	Short: "Schema JSON: Convert JSON Schema to PHP Data Object",
	Long:  "Schema JSON: Convert JSON Schema to PHP Data Object",
	Run: func(cmd *cobra.Command, args []string) {
		phpClass, err := converter.GeneratePHPFromSchema(schemaInput, className)
		if err != nil {
			fmt.Printf("Error generating PHP class: %v\n", err)
			return
		}
		fmt.Println(phpClass)
	},
}

func init() {
	rootCmd.AddCommand(schemaJsonCmd)

	schemaJsonCmd.Flags().StringVarP(&schemaInput, "schema", "s", "", "JSON Schema input (required)")
	schemaJsonCmd.Flags().StringVarP(&className, "class", "c", "", "Name of the PHP class (required)")
	schemaJsonCmd.MarkFlagRequired("schema")
	schemaJsonCmd.MarkFlagRequired("class")
}
