package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Config struct {
	FileName string `json:"filename"`
}

func getJSONData(yamlData string) (string, error) {
	var data interface{}
	if yamlData == "" {
		return "", errors.New("File should not be empty")
	}
	//Unmarshaling the yaml file
	if err := yaml.Unmarshal([]byte(yamlData), &data); err != nil {
		return "", err
	}
	//convert yaml to json
	data = convertYamltoJSON(data)
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func convertYamltoJSON(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			if w, ok := k.(int); ok {
				k = strconv.Itoa(w)
			}
			if w, ok := k.(float64); ok {
				k = fmt.Sprintf("%f", w)
			}
			m2[k.(string)] = convertYamltoJSON(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convertYamltoJSON(v)
		}
	}
	return i
}

//Check the file exist in the current path
func CheckFileExist(filename string) (string, error) {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(yamlFile), nil
}

func getConfiguration(filename string) (string, error) {
	var config = Config{}
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println("Error during opening configuration file: ", err)
		return "", err
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Error during decoding configuration file: ", err)
		return "", err
	}

	return config.FileName, nil
}

func main() {
	file, err := getConfiguration("config.json")
	if err != nil {
		fmt.Println("Error getting on", err)
		os.Exit(1)
	}

	fileData, err := CheckFileExist(file)
	if err != nil {
		fmt.Println("Error getting on", err)
		os.Exit(1)
	}

	JSONstring, err := getJSONData(fileData)
	if err != nil {
		fmt.Println("Error getting on", err)
		os.Exit(1)
	}
	fmt.Println(JSONstring)
}
