package parsers

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type Parser interface {
	Parse(filepath string) (map[string]any, error)
}

type ParserJson struct{}

func (parser ParserJson) Parse(filepath string) (map[string]any, error) {
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

func ParseFile(filepath string) (map[string]any, error) {

	fileInfo, err := os.Lstat(filepath)

	if err != nil {
		return map[string]any{}, errors.New("path does not exist")
	}

	if fileInfo.IsDir() {
		return map[string]any{}, errors.New("need to select congfiguration json file")
	}

	var parser Parser

	switch {
	case strings.HasSuffix(fileInfo.Name(), ".json"):
		parser = ParserJson{}
	}

	return parser.Parse(filepath)

}
