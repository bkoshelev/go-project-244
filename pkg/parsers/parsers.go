package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

type Parser interface {
	Parse(data []byte) (map[string]any, error)
}

type ParserJson struct{}

func (parser ParserJson) Parse(data []byte) (map[string]any, error) {

	result := map[string]any{}

	err := json.Unmarshal(data, &result)

	if err != nil {
		return map[string]any{}, errors.New("fail to parse json")
	}
	return result, nil
}

type ParserYml struct{}

func (parser ParserYml) Parse(data []byte) (map[string]any, error) {
	result := map[string]any{}

	err := yaml.Unmarshal(data, &result)

	if err != nil {
		return map[string]any{}, errors.New("fail to parse yaml")
	}
	return result, nil
}

func getParser(filename string) (Parser, error) {
	switch {
	case strings.HasSuffix(filename, ".json"):
		return ParserJson{}, nil
	case strings.HasSuffix(filename, ".yml"):
		return ParserYml{}, nil
	default:
		return nil, fmt.Errorf("unsupported file type: %v", filename)
	}
}

func ParseFile(filepath string) (map[string]any, error) {
	fileInfo, err := os.Lstat(filepath)

	if err != nil {
		return map[string]any{}, fmt.Errorf("path %v does not exist: %w", filepath, err)
	}

	if fileInfo.IsDir() {
		return map[string]any{}, fmt.Errorf("need to select congfiguration json file: %w", err)
	}

	fileData, err := os.ReadFile(filepath)

	if err != nil {
		return map[string]any{}, fmt.Errorf("fail to file reading: %w", err)
	}

	parser, err := getParser(fileInfo.Name())

	if err != nil {
		return map[string]any{}, fmt.Errorf("parser creation fail: %w", err)
	}

	return parser.Parse(fileData)

}
