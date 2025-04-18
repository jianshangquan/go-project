package utils

import "encoding/json"


func ParseJsonFromBytes[T any](data []byte) (T, error){
	var js T
	err := json.Unmarshal(data, &js)
	return js, err
}