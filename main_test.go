package main

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

const (
	mapData = `Test:
-   TestID:
    -   ID: 1
        Name: Test1
    -   ID: 2
        Name: Test2
`
	mapDataSpecialCharacter = `Test:
-   TestID:
    -   %z: 1
        Name: Test1
    -   ID: 2
        Name: Test2
`
	mapWithIntData = `Test:
-   TestID:
    -   1: test1       
`
	mapWithFloatData = `Test:
-   TestID:
    -   1.1: test1       
`
)

func Test_getJSONData(t *testing.T) {
	tests := []struct {
		testCase string
		data     string
		wantErr  bool
	}{
		{
			testCase: "Test Failure - With empty data",
			data:     "",
			wantErr:  true,
		},
		{
			testCase: "Test Failure - invalid yaml data",
			data: `Test:
			-   TestIDS:
				-   ID: 1
					Name: aaa
				-   ID: 2
					Name: bbb`,
			wantErr: true,
		},
		{
			testCase: "Test failure - with special character",
			data:     mapDataSpecialCharacter,
			wantErr:  true,
		},
		{
			testCase: "Test Success - valid yaml data",
			data:     mapData,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			if _, err := getJSONData(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("getJSONData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestFileExist(t *testing.T) {
	tests := []struct {
		testCase string
		fileName string
		wantErr  bool
	}{
		{
			testCase: "Test Failure - With non existing filename",
			fileName: "notexistfile.yaml",
			wantErr:  true,
		},
		{
			testCase: "Test Success -valid file name",
			fileName: "config.yaml",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			if _, err := CheckFileExist(tt.fileName); (err != nil) != tt.wantErr {
				t.Errorf("CheckFileExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getConfiguration(t *testing.T) {
	f, err := os.OpenFile("invalid.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error", err)
	}
	_, err = f.Write([]byte(`{"name:invalid}`))
	if err != nil {
		fmt.Println("Error", err)
	}

	tests := []struct {
		testCase string
		fileName string
		wantErr  bool
	}{
		{
			testCase: "Test Failure - With non existing filename",
			fileName: "notexistfile.json",
			wantErr:  true,
		},
		{
			testCase: "Test Failure - With invalid JSON",
			fileName: "invalid.json",
			wantErr:  true,
		},
		{
			testCase: "Test Success -valid file name",
			fileName: "config.json",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {

			if _, err := getConfiguration(tt.fileName); (err != nil) != tt.wantErr {
				t.Errorf("getConfiguration() error = %v, wantErr %v", err, tt.wantErr)

			}

		})
	}

	f.Close()
	if err := os.Remove("invalid.json"); err != nil {
		fmt.Println("unable to delete ", err)
	}
}
func Test_convertYamltoJSON(t *testing.T) {
	tests := []struct {
		testCase string
		data     string
		wantErr  bool
	}{
		{
			testCase: "Test Failure - With empty data",
			data:     "",
			wantErr:  false,
		},
		{
			testCase: "Test Success - With string",
			data:     "hello",
			wantErr:  true,
		},
		{
			testCase: "Test Success - With struct",
			data:     mapData,
			wantErr:  true,
		},
		{
			testCase: "Test Success - With int",
			data:     mapWithIntData,
			wantErr:  true,
		},
		{
			testCase: "Test Success - With float",
			data:     mapWithFloatData,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			var body interface{}
			yaml.Unmarshal([]byte(tt.data), &body)
			if i := convertYamltoJSON(body); (i != nil) != tt.wantErr {
				t.Errorf("convertYamltoJSON() got = %v, wantErr %v", i, tt.wantErr)
				return
			}
		})
	}

}
