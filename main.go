package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSONConfiguration specifies how to unmarshal a given configuration file
type JSONConfiguration struct {
	Name     string                 `json:"name"`
	Language string                 `json:"language"`
	Skills   map[string]interface{} `json:"skills"`
	Active   bool                   `json:"active"`
}

// UnmarshalJSON parses a JSON configuration specified by the 'JSONConfiguration' struct
func UnmarshalJSON(config *JSONConfiguration) {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	config := JSONConfiguration{}
	UnmarshalJSON(&config)
	fmt.Println(config)
	fmt.Println(config.Skills["Golang"]) // pull out a value from the configuration file
}
