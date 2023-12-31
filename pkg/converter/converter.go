package converter

import (
	"encoding/json"
	"fmt"
	"strings"
)

type SchemaProperty struct {
	Type string `json:"type"`
}

type JSONSchema struct {
	Properties map[string]SchemaProperty `json:"properties"`
	Type       string                    `json:"type"`
}

func GeneratePHPDTO(data map[string]interface{}, className string) string {
	var properties []string
	for key, value := range data {
		dataType := getType(value)
		property := fmt.Sprintf("private %s $%s", dataType, key)
		properties = append(properties, property)
	}

	return createDTOClass(className, properties)
}

func getType(value interface{}) string {
	switch v := value.(type) {
	case string:
		return "string"
	case float64:
		// Check if the float64 value is actually an integer.
		if float64(int(v)) == v {
			return "int"
		}
		return "float"
	case int, int32, int64:
		return "int"
	case bool:
		return "bool"
	case map[string]interface{}:
		return "array"
	case []interface{}:
		return "array"
	default:
		return "mixed"
	}
}

func createDTOClass(className string, properties []string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("final readonly class %s\n", className))
	builder.WriteString("{\n")
	builder.WriteString("    public function __construct(\n")
	for _, prop := range properties {
		builder.WriteString("        " + prop + ",\n")
	}
	builder.WriteString("    ) {}\n")
	builder.WriteString("}")

	return builder.String()
}

func GeneratePHPFromSchema(schema string, className string) (string, error) {
	var schemaObj JSONSchema
	if err := json.Unmarshal([]byte(schema), &schemaObj); err != nil {
		return "", err
	}

	if schemaObj.Properties == nil {
		return "", fmt.Errorf("no properties found in the schema")
	}

	data := make(map[string]interface{})
	for key, propData := range schemaObj.Properties {
		data[key] = mapJSONSchemaTypeToPlaceholderValue(propData.Type)
	}

	return GeneratePHPDTO(data, className), nil
}

func mapJSONSchemaTypeToPlaceholderValue(dataType string) interface{} {
	switch dataType {
	case "string":
		return ""
	case "number":
		return 0.0
	case "integer":
		return 0
	case "boolean":
		return false
	case "array":
		return []interface{}{}
	case "object":
		return map[string]interface{}{}
	default:
		return nil
	}
}
