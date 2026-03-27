package code

import (
	"fmt"

	diffbuilder "github.com/bkoshelev/go-project-244/pkg/diff_builder"
	"github.com/bkoshelev/go-project-244/pkg/formatters"
	"github.com/bkoshelev/go-project-244/pkg/parsers"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := parsers.ParseFile(filepath1)

	if err != nil {
		return "", fmt.Errorf("fail to parse first json file, %w", err)
	}

	data2, err := parsers.ParseFile(filepath2)

	if err != nil {
		return "", fmt.Errorf("fail to parse second json file, %w", err)
	}

	diff := diffbuilder.CreateDiff(data1, data2)

	formatter := formatters.GetFormatter(format)

	return formatter.Format(diff), nil
}
