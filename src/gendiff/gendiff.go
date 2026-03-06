package code

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/bkoshelev/go-project-244/src/parser"
	"github.com/samber/lo"
)

const (
	ADDED = iota
	REMOVED
	CHANGED
	UNCHANGED
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	json1, err := parser.ParseJsonFile(filepath1)

	if err != nil {
		return "", fmt.Errorf("fail to parse first json file, %w", err)
	}

	json2, err := parser.ParseJsonFile(filepath2)

	if err != nil {
		return "", fmt.Errorf("fail to parse second json file, %w", err)
	}

	// https://stackoverflow.com/a/69889828
	json1Keys, json2Keys := slices.Collect(maps.Keys(json1)), slices.Collect(maps.Keys(json2))

	keys := lo.Union(json1Keys, json2Keys)
	slices.Sort(keys)

	info := map[string]int{}

	diff1, diff2 := lo.Difference(json2Keys, json1Keys)

	for _, key := range keys {

		switch {
		case slices.Contains(diff1, key):
			info[key] = ADDED
		case slices.Contains(diff2, key):
			info[key] = REMOVED
		case json1[key] != json2[key]:
			info[key] = CHANGED
		default:
			info[key] = UNCHANGED
		}
	}

	var builder strings.Builder

	builder.WriteString("{\n")

	for _, key := range keys {

		switch info[key] {
		case UNCHANGED:
			builder.WriteString(fmt.Sprintf("    %s: %v\n", key, json1[key]))
		case ADDED:
			builder.WriteString(fmt.Sprintf("  + %s: %v\n", key, json2[key]))
		case REMOVED:
			builder.WriteString(fmt.Sprintf("  - %s: %v\n", key, json1[key]))
		case CHANGED:
			builder.WriteString(fmt.Sprintf("  - %s: %v\n", key, json1[key]))
			builder.WriteString(fmt.Sprintf("  + %s: %v\n", key, json2[key]))
		default:
			return "", errors.New("неизвестный тип отличия между json")
		}
	}

	builder.WriteString("}")

	return builder.String(), nil
}
