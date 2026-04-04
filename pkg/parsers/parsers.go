package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

type Parser interface {
	Parse(data []byte) (map[string]any, error)
}

type JSONParser struct{}

func (parser JSONParser) Parse(data []byte) (map[string]any, error) {

	result := map[string]any{}

	err := json.Unmarshal(data, &result)

	if err != nil {
		return nil, fmt.Errorf("fail to parse json file: %w", err)
	}
	return result, nil
}

type YAMLParser struct{}

func (parser YAMLParser) Parse(data []byte) (map[string]any, error) {
	result := map[string]any{}

	err := yaml.Unmarshal(data, &result)

	if err != nil {
		return nil, fmt.Errorf("fail to parse yaml file: %w", err)
	}
	return result, nil
}

func getParser(filename string) (Parser, error) {
	switch {
	case strings.HasSuffix(filename, ".json"):
		return JSONParser{}, nil
	case strings.HasSuffix(filename, ".yml") || strings.HasSuffix(filename, ".yaml"):
		return YAMLParser{}, nil
	default:
		return nil, fmt.Errorf("unsupported file type: %v", filename)
	}
}

func ParseFile(filepath string) (map[string]any, error) {
	fileInfo, err := os.Lstat(filepath)

	if err != nil {
		return nil, fmt.Errorf("path %v does not exist: %w", filepath, err)
	}

	if fileInfo.IsDir() {
		return nil, fmt.Errorf("need to select valid congfiguration file instead of %s", filepath)
	}

	fileData, err := os.ReadFile(filepath)

	if err != nil {
		return nil, fmt.Errorf("fail to file reading: %w", err)
	}

	parser, err := getParser(fileInfo.Name())

	if err != nil {
		return nil, fmt.Errorf("parser creation fail: %w", err)
	}

	return parser.Parse(fileData)

}
