package code

import (
	"fmt"

	"github.com/bkoshelev/go-project-244/pkg/parsers"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	json1, err := parsers.ParseFile(filepath1)

	if err != nil {
		return "", fmt.Errorf("fail to parse first json file, %w", err)
	}

	json2, err := parsers.ParseFile(filepath2)

	if err != nil {
		return "", fmt.Errorf("fail to parse second json file, %w", err)
	}

	return fmt.Sprintf("json1: %v, json2: %v", json1, json2), nil
}
