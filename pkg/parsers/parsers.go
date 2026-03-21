package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
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
		return map[string]any{}, fmt.Errorf("fail to file reading: %w", err)
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
		return map[string]any{}, fmt.Errorf("path %v does not exist: %w", filepath, err)
	}

	if fileInfo.IsDir() {
		return map[string]any{}, fmt.Errorf("need to select congfiguration json file: %w", err)
	}

	var parser Parser

	switch {
	case strings.HasSuffix(fileInfo.Name(), ".json"):
		parser = ParserJson{}
	}

	return parser.Parse(filepath)

}
