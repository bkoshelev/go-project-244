package formatters

import (
	"encoding/json"
	"fmt"

	diffbuilder "github.com/bkoshelev/go-project-244/pkg/diff_builder"
)

type JSONFormatter struct{}

func (fmtr JSONFormatter) Format(diff []diffbuilder.Node) (string, error) {

	if len(diff) == 0 {
		result, _ := json.Marshal(map[string]any{})
		return string(result), nil
	}

	result, e := json.Marshal(diffbuilder.Node{
		Key:      "",
		Value1:   nil,
		Value2:   nil,
		NodeType: "ROOT",
		Children: diff,
	})

	if e != nil {
		return "", fmt.Errorf("creation json output fail: %w", e)
	}
	return string(result), nil
}

func CreateJSONFormatter() JSONFormatter {
	return JSONFormatter{}
}
