package json

import "encoding/json"

func ParseJSONToMap(jsonInput string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonInput), &data)
	return data, err
}
