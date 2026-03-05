package parser

import (
	"encoding/json"
	"errors"
	"os"
)

// type Reader interface {
// 	ReadFile(filepath string) (string, error)
// }

// type FileReader struct{}

// func (f FileReader) ReadFile(filepath string) (string, error) {
// 	return "Hello", nil
// }

// func (f FileReader) ReadJson(filepath string) (string, error) {
// 	return "Hello", nil
// }

func ParseJsonFile(filepath string) (map[string]any, error) {

	fileInfo, err := os.Lstat(filepath)

	if err != nil {
		return map[string]any{}, errors.New("path does not exist")
	}

	if fileInfo.IsDir() {
		return map[string]any{}, errors.New("need to select congfiguration json file")
	}

	fileData, err := os.ReadFile(filepath)

	if err != nil {
		return map[string]any{}, errors.New("fail to file reading")
	}

	result := map[string]any{}

	err = json.Unmarshal(fileData, &result)

	if err != nil {
		return map[string]any{}, errors.New("fail to parse json")
	}
	return result, nil
}
