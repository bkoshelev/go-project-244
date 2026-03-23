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
	NodeType NodeType
	Children []Node
}

func CreateDiff(data1, data2 map[string]any) (diff []Node) {
	// взято отсюда https://stackoverflow.com/a/69889828
	data1Keys, data2Keys := slices.Collect(maps.Keys(data1)), slices.Collect(maps.Keys(data2))

	allKeys := lo.Union(data1Keys, data2Keys)
	slices.Sort(allKeys)

	diff1, diff2 := lo.Difference(data1Keys, data2Keys)

	for _, key := range allKeys {

		value1 := data1[key]
		_, ok := value1.(map[string]any)
		value2 := data2[key]
		_, ok2 := value2.(map[string]any)

		if ok && ok2 {
			children := CreateDiff(value1.(map[string]any), value2.(map[string]any))
			diff = append(diff, Node{key, CHILDREN, children})
			continue
		}

		switch {
		case slices.Contains(diff2, key):
			diff = append(diff, Node{key, ADDED, nil})
		case slices.Contains(diff1, key):
			diff = append(diff, Node{key, REMOVED, nil})
		case data1[key] != data2[key]:
			diff = append(diff, Node{key, CHANGED, nil})

		default:
			diff = append(diff, Node{key, UNCHANGED, nil})
		}
	}

	return diff
}
