package automapper

import "encoding/json"

func Map(objOrigin interface{}, objDestination interface{}) {
	jsonOrigin := StructToJson(objOrigin)
	err := json.Unmarshal([]byte(jsonOrigin), objDestination)
	if err != nil {
		return
	}
}

func StructToJson(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}

func JsonToStruct(jsonStr string) (interface{}, error) {
	var data interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)

	return data, err
}