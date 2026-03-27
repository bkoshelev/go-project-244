package diffbuilder

import (
	"maps"
	"slices"

	"github.com/samber/lo"
)

const (
	ADDED     = "ADDED"
	REMOVED   = "REMOVED"
	CHANGED   = "CHANGED"
	UNCHANGED = "UNCHANGED"
	CHILDREN  = "CHILDREN"
)

type NodeType string

type Node struct {
	Key      string
	Value1   any
	Value2   any
	NodeType NodeType
	Children []Node
}

func CreateDiff(data1, data2 map[string]any) (diff []Node) {
	// взято отсюда https://stackoverflow.com/a/69889828
	data1Keys, data2Keys := slices.Collect(maps.Keys(data1)), slices.Collect(maps.Keys(data2))

	allKeys := lo.Union(data1Keys, data2Keys)
	slices.Sort(allKeys)

	removedKeys, addedKeys := lo.Difference(data1Keys, data2Keys)

	for _, key := range allKeys {

		value1 := data1[key]
		_, val1IsMap := value1.(map[string]any)
		value2 := data2[key]
		_, val2IsMap := value2.(map[string]any)

		var children []Node
		var status NodeType

		switch {
		case val1IsMap && val2IsMap:
			children = CreateDiff(value1.(map[string]any), value2.(map[string]any))
			status = CHILDREN
		case slices.Contains(addedKeys, key):
			status = ADDED
		case slices.Contains(removedKeys, key):
			status = REMOVED
		case data1[key] != data2[key]:
			status = CHANGED
		default:
			status = UNCHANGED
		}

		diff = append(diff, Node{key, value1, value2, status, children})
	}

	return diff
}
