package formatters

import (
	"fmt"
	"strings"

	diffbuilder "github.com/bkoshelev/go-project-244/pkg/diff_builder"
)

type PlainFormatter struct{}

func (fmtr PlainFormatter) fmtVal(content any) string {
	switch v := content.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	case nil:
		return "null"
	case map[string]any:
		return "[complex value]"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (fmtr PlainFormatter) fmtDiff(path []string, diff []diffbuilder.Node) string {
	var builder strings.Builder

	for _, node := range diff {

		newPath := append(path, node.Key)
		newPathS := strings.Join(newPath, ".")

		fmtVal1 := fmtr.fmtVal(node.Value1)
		fmtVal2 := fmtr.fmtVal(node.Value2)

		switch node.NodeType {
		case diffbuilder.CHILDREN:
			fmtDiff := fmtr.fmtDiff(newPath, node.Children)
			fmt.Fprintf(&builder, "%s", fmtDiff)
		case diffbuilder.UNCHANGED:
			continue
		case diffbuilder.ADDED:
			fmt.Fprintf(&builder, "Property '%s' was added with value: %s\n", newPathS, fmtVal2)
		case diffbuilder.REMOVED:
			fmt.Fprintf(&builder, "Property '%s' was removed\n", newPathS)
		case diffbuilder.CHANGED:
			fmt.Fprintf(&builder, "Property '%s' was updated. From %s to %s\n", newPathS, fmtVal1, fmtVal2)
		}
	}

	return builder.String()
}

func (fmtr PlainFormatter) Format(diff []diffbuilder.Node) (string, error) {
	return strings.Trim(fmtr.fmtDiff([]string{}, diff), "\n"), nil
}
func CreatePlainFormatter() PlainFormatter {
	return PlainFormatter{}
}
