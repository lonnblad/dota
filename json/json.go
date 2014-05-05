package json

import (
	"encoding/json"
	"log"
)

func CreateJSON(data interface{}) string {
	json, err := json.Marshal(data)
	if err != nil {
		log.Printf("=== marshal error ===\n%+v\n", err)
	}
	return string(json)
}

func ParseJSON(bs []byte) (data interface{}) {
	err := json.Unmarshal(bs, &data)
	if err != nil {
		log.Println("error:", err)
	}
	return
}
