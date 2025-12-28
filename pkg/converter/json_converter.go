package converter

import "encoding/json"

// JsonMarshal JSON序列化
func JsonMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JsonUnmarshal JSON反序列化
func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
