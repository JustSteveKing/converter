package cmd

import (
	"fmt"

	"github.com/juststeveking/converter/pkg/converter"
	"github.com/juststeveking/converter/pkg/yaml"
	"github.com/spf13/cobra"
)

var yamlInput string

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "YAML: Convert YAML into PHP Data Object",
	Long:  "YAML: Convert YAML into PHP Data Object",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := yaml.ParseYAMLToMap(yamlInput)
		if err != nil {
			fmt.Printf("Error parsing YAML: %v\n", err)
			return
		}
		dto := converter.GeneratePHPDTO(data, className)
		fmt.Println(dto)
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	// Adding flags for our command
	yamlCmd.Flags().StringVarP(&className, "class", "c", "", "Name of the DTO class (required)")
	yamlCmd.Flags().StringVarP(&yamlInput, "yaml", "y", "", "YAML input (required)")
	yamlCmd.MarkFlagRequired("class")
	yamlCmd.MarkFlagRequired("yaml")
}
