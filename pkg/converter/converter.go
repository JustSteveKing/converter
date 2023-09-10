package converter

import (
	"fmt"
	"strings"
)

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
	switch value.(type) {
	case string:
		return "string"
	case float64:
		return "float"
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
