package yaml

import "gopkg.in/yaml.v2"

func ParseYAMLToMap(yamlInput string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlInput), &data)
	return data, err
}
