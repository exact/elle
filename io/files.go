package io

import (
	"encoding/json"
	"os"
)

// LoadJSON simply attempts to load JSON data from a file into a data structure.
// It returns any errors encountered.
func LoadJSON(f string, v any) error {
	data, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// SaveJSON simply attempts to save JSON data to a file from a given data structure.
// It returns any errors encountered.
func SaveJSON(f string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(f, data, 0o644)
}
