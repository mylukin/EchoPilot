package helper

import "encoding/json"

// JSONBody is send message
type JSONBody map[string]any

// Error
func (body JSONBody) Error() string {
	data, err := json.Marshal(body)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

// MergeJSON merge json body
func MergeJSON(map1 JSONBody, maps ...JSONBody) JSONBody {
	for _, mapv := range maps {
		for k, v := range mapv {
			map1[k] = v
		}
	}
	return map1
}
