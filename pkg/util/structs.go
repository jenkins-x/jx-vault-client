package util

import (
	"encoding/json"
)

// ToMapStringInterfaceFromStruct marshals a struct to a generic map[string]interface{} by marshalling it to json and back
// Use JSON for the marshalling instead of YAML because sub-structs will get marshalled into map[interface{}]interface{}
// when using YAML, but map[string]interface{} when using JSON and vault libraries can't handle map[interface{}]interface{}
func ToMapStringInterfaceFromStruct(obj interface{}) (map[string]interface{}, error) {
	y, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}
	out := make(map[string]interface{})
	err = json.Unmarshal(y, &out)
	return out, err
}

// ToStructFromMapStringInterface marshals a generic map[string]interface{} to a struct by marshalling to json and back
// Use JSON for the marshalling instead of YAML because sub-structs will get marshalled into map[interface{}]interface{}
// when using YAML, but map[string]interface{} when using JSON and vault libraries can't handle map[interface{}]interface{}
func ToStructFromMapStringInterface(m map[string]interface{}, str interface{}) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(j, str)
}
