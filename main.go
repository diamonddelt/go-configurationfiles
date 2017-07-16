package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-ini/ini"

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

// INIConfiguration specifies how to map this project-specific Windows INI file to a struct
type INIConfiguration struct {
	Name     string `ini:"name"`
	Language string `ini:"language"`
	Skills   struct {
		Golang   int `ini:"golang"`
		Python   int `ini:"python"`
		Java     int `ini:"java"`
		CSharp   int `ini:"csharp"`
		Training int `ini:"training"`
	}
	State struct {
		Active bool `ini:"active"`
	}
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

// MapINI maps the current configuration of an .ini file to the config object reference at the specified path
func MapINI(config *INIConfiguration, path string) {
	err := ini.MapTo(&config, path)
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

	configINI := INIConfiguration{}
	MapINI(&configINI, "config.ini")
	fmt.Println(`The value for "active" in the [State] section is:`, configINI.State.Active) // pull out a value from the INI configuration file
}
