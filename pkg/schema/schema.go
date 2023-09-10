package schema

import "encoding/json"

func GenerateSchemaFromJSON(jsonInput string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonInput), &data); err != nil {
		return "", err
	}

	schema := generateSchema(data)
	result, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func generateSchema(data interface{}) map[string]interface{} {
	schema := map[string]interface{}{}

	switch v := data.(type) {
	case string:
		schema["type"] = "string"
	case bool:
		schema["type"] = "boolean"
	case float64:
		if float64(int(v)) == v {
			schema["type"] = "integer"
		} else {
			schema["type"] = "number"
		}
	case map[string]interface{}:
		schema["type"] = "object"
		properties := map[string]interface{}{}
		for key, value := range v {
			properties[key] = generateSchema(value)
		}
		schema["properties"] = properties
	case []interface{}:
		schema["type"] = "array"
		if len(v) > 0 {
			schema["items"] = generateSchema(v[0])
		}
	}

	return schema
}
