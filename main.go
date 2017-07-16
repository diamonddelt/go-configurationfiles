package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// JSONConfiguration specifies how to unmarshal a given JSON configuration file
type JSONConfiguration struct {
	Name     string                 `json:"name"`
	Language string                 `json:"language"`
	Skills   map[string]interface{} `json:"skills"`
	Active   bool                   `json:"active"`
}

// YAMLConfiguration specifies how to unmarshal a given YAML configuration file
type YAMLConfiguration struct {
	Name     string         `yaml:"name"`
	Language string         `yaml:"language"`
	Skills   map[string]int `yaml:"skills"`
	Active   bool           `yaml:"active"`
}

// DecodeJSON parses a JSON configuration specified by the 'JSONConfiguration' struct
func DecodeJSON(config *JSONConfiguration, path string) {
	file, _ := os.Open(path)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// UnmarshalYAML unmarshalls a YAML configuration into a struct for logical key/value access
func UnmarshalYAML(config *YAMLConfiguration, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	configJSON := JSONConfiguration{}
	DecodeJSON(&configJSON, "config.json")
	fmt.Println(`The value for "skills: golang" is:`, configJSON.Skills["golang"]) // pull out a value from the JSON configuration file

	configYAML := YAMLConfiguration{}
	UnmarshalYAML(&configYAML, "config.yaml")
	fmt.Println(`The value for "skills: training" is:`, configYAML.Skills["training"]) // pull out a value from the YAML configuration file
}
