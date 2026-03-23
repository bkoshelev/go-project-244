package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	diffbuilder "github.com/bkoshelev/go-project-244/pkg/diff_builder"
)

const INDENT = 4

type Formatter interface {
	Format(key string, data1, data2 map[string]any, diff []diffbuilder.Node) string
}

type StylishFormatter struct {
	depth int
}

func formatMap(data map[string]any, depth int) string {
	var builder strings.Builder

	keys := slices.Collect(maps.Keys(data))
	slices.Sort(keys)

	builder.WriteString("{\n")

	for _, key := range keys {
		switch data[key].(type) {
		case map[string]any:
			fmt.Fprintf(&builder, "%s%s: %v\n", strings.Repeat(" ", 4*(depth+1)), key, formatMap(data[key].(map[string]any), depth+1))
		default:
			fmt.Fprintf(&builder, "%s%s: %v\n", strings.Repeat(" ", 4*(depth+1)), key, data[key])
		}
	}

	fmt.Fprintf(&builder, "%s}", strings.Repeat(" ", 4*(depth)))

	return builder.String()
}

func formatByType(content any, depth int) string {
	switch v := content.(type) {
	case nil:
		return "null"
	case map[string]any:
		return formatMap(v, depth)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (formatter StylishFormatter) Format(key string, data1, data2 map[string]any, diff []diffbuilder.Node) string {
	var builder strings.Builder

	if key != "" {
		builder.WriteString(strings.Repeat(" ", 4*(formatter.depth)) + key + ": {\n")

	} else {
		builder.WriteString("{\n")
	}

	for _, node := range diff {

		switch node.NodeType {
		case diffbuilder.CHILDREN:
			formattedChilden := StylishFormatter{formatter.depth + 1}.
				Format(node.Key,
					data1[node.Key].(map[string]any),
					data2[node.Key].(map[string]any),
					node.Children,
				)

			fmt.Fprintf(&builder, "%s\n", formattedChilden)
		case diffbuilder.UNCHANGED:
			fmt.Fprintf(&builder, "%s%s: %v\n", strings.Repeat(" ", 4*(formatter.depth+1)), node.Key, data1[node.Key])
		case diffbuilder.ADDED:
			fmt.Fprintf(&builder, "%s+ %s: %v\n", strings.Repeat(" ", 4*(formatter.depth+1)-2), node.Key, formatByType(data2[node.Key], formatter.depth+1))
		case diffbuilder.REMOVED:
			fmt.Fprintf(&builder, "%s- %s: %v\n", strings.Repeat(" ", 4*(formatter.depth+1)-2), node.Key, formatByType(data1[node.Key], formatter.depth+1))
		case diffbuilder.CHANGED:
			fmt.Fprintf(&builder, "%s- %s: %v\n", strings.Repeat(" ", 4*(formatter.depth+1)-2), node.Key, formatByType(data1[node.Key], formatter.depth+1))
			fmt.Fprintf(&builder, "%s+ %s: %v\n", strings.Repeat(" ", 4*(formatter.depth+1)-2), node.Key, formatByType(data2[node.Key], formatter.depth+1))
		}
	}

	fmt.Fprintf(&builder, "%s}", strings.Repeat(" ", 4*(formatter.depth)))

	return builder.String()
}

func CreateStylishFormatter() StylishFormatter {
	return StylishFormatter{0}
}

func GetFormatter(formatterName string) Formatter {
	switch formatterName {
	case "stylish":
		return CreateStylishFormatter()
	default:
		return CreateStylishFormatter()
	}
}
