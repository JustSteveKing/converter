package cmd

import (
	"fmt"

	"github.com/juststeveking/converter/pkg/schema"
	"github.com/spf13/cobra"
)

var jsonForSchema string

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Schema: Turn a JSON payload into a JSON Schema",
	Long:  "Schema: Turn a JSON payload into a JSON Schema",
	Run: func(cmd *cobra.Command, args []string) {
		schema, err := schema.GenerateSchemaFromJSON(jsonForSchema)
		if err != nil {
			fmt.Printf("Error generating schema: %v\n", err)
			return
		}
		fmt.Println(schema)
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)

	schemaCmd.Flags().StringVarP(&jsonForSchema, "json", "j", "", "JSON input to generate schema (required)")
	schemaCmd.MarkFlagRequired("json")
}
