package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	diffbuilder "github.com/bkoshelev/go-project-244/pkg/diff_builder"
)

type StylishFormatter struct{}

func (fmtr StylishFormatter) formatMap(data map[string]any, depth int) string {
	var builder strings.Builder

	keys := slices.Collect(maps.Keys(data))
	slices.Sort(keys)

	newDepth := depth + 1
	indent := strings.Repeat(" ", 4*(newDepth))

	builder.WriteString("{\n")

	for _, key := range keys {
		value := fmtr.fmtVal(data[key], newDepth)
		fmt.Fprintf(&builder, "%s%s: %v\n", indent, key, value)
	}

	fmt.Fprintf(&builder, "%s}", strings.Repeat(" ", 4*(depth)))

	return builder.String()
}

func (fmtr StylishFormatter) fmtVal(content any, depth int) string {
	switch v := content.(type) {
	case nil:
		return "null"
	case map[string]any:
		return fmtr.formatMap(v, depth)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (fmtr StylishFormatter) fmtDiff(diff []diffbuilder.Node, objKey string, depth int) string {
	var builder strings.Builder

	if objKey != "" {
		builder.WriteString(strings.Repeat(" ", 4*depth) + objKey + ": {\n")
	} else {
		builder.WriteString("{\n")
	}

	newDepth := depth + 1

	for _, node := range diff {

		fmtVal1 := fmtr.fmtVal(node.Value1, newDepth)
		fmtVal2 := fmtr.fmtVal(node.Value2, newDepth)

		switch node.NodeType {
		case diffbuilder.CHILDREN:
			fmtDiff := fmtr.fmtDiff(node.Children, node.Key, newDepth)
			fmt.Fprintf(&builder, "%s\n", fmtDiff)
		case diffbuilder.UNCHANGED:
			fmt.Fprintf(&builder, "%s%s: %v\n", strings.Repeat(" ", 4*newDepth), node.Key, node.Value1)
		case diffbuilder.ADDED:
			fmt.Fprintf(&builder, "%s+ %s: %v\n", strings.Repeat(" ", 4*newDepth-2), node.Key, fmtVal2)
		case diffbuilder.REMOVED:
			fmt.Fprintf(&builder, "%s- %s: %v\n", strings.Repeat(" ", 4*newDepth-2), node.Key, fmtVal1)
		case diffbuilder.CHANGED:
			fmt.Fprintf(&builder, "%s- %s: %v\n", strings.Repeat(" ", 4*newDepth-2), node.Key, fmtVal1)
			fmt.Fprintf(&builder, "%s+ %s: %v\n", strings.Repeat(" ", 4*newDepth-2), node.Key, fmtVal2)
		}
	}

	fmt.Fprintf(&builder, "%s}", strings.Repeat(" ", 4*depth))

	return builder.String()
}

func (fmtr StylishFormatter) Format(diff []diffbuilder.Node) string {
	return fmtr.fmtDiff(diff, "", 0)
}

func CreateStylishFormatter() StylishFormatter {
	return StylishFormatter{}
}
