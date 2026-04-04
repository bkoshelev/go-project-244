package diffbuilder

import (
	"maps"
	"reflect"
	"slices"

	"github.com/samber/lo"
)

type NodeType string

const (
	ADDED     NodeType = "ADDED"
	REMOVED   NodeType = "REMOVED"
	CHANGED   NodeType = "CHANGED"
	UNCHANGED NodeType = "UNCHANGED"
	CHILDREN  NodeType = "CHILDREN"
)

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

	for _, key := range allKeys {

		value1, existInData1 := data1[key]
		_, val1IsMap := value1.(map[string]any)
		value2, existInData2 := data2[key]
		_, val2IsMap := value2.(map[string]any)

		var children []Node
		var status NodeType

		switch {
		case val1IsMap && val2IsMap:
			children = CreateDiff(value1.(map[string]any), value2.(map[string]any))
			status = CHILDREN
		case !existInData1 && existInData2:
			status = ADDED
		case existInData1 && !existInData2:
			status = REMOVED
		case !reflect.DeepEqual(value1, value2):
			status = CHANGED
		default:
			status = UNCHANGED
		}

		diff = append(diff, Node{key, value1, value2, status, children})
	}

	return diff
}
