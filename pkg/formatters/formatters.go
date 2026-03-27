package formatters

import diffbuilder "github.com/bkoshelev/go-project-244/pkg/diff_builder"

const INDENT = 4

type Formatter interface {
	Format(diff []diffbuilder.Node) string
}

func GetFormatter(formatterName string) Formatter {
	switch formatterName {
	case "stylish":
		return CreateStylishFormatter()
	case "plain":
		return CreatePlainFormatter()
	default:
		return CreateStylishFormatter()
	}
}
