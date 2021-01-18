package main

import (
	"encoding/json"
)

func jsonEncode(i interface{}) ([]byte, error) {
	b, err := json.Marshal(i)
	return b, err
}

func jsonDecode(b []byte, i interface{}) error {
	err := json.Unmarshal(b, i)
	return err
}
