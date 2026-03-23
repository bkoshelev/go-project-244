package diffbuilder

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestCreateDiff(t *testing.T) {

	cases := []struct {
		data1, data2 map[string]any
		want         []Node
	}{
		{
			map[string]any{"follow": false, "host": "hexlet.io", "proxy": "123.234.53.22", "timeout": 50.0},
			map[string]any{"verbose": false, "host": "hexlet.io", "timeout": 20.0},
			[]Node{
				{"follow", REMOVED, nil},
				{"host", UNCHANGED, nil},
				{"proxy", REMOVED, nil},
				{"timeout", CHANGED, nil},
				{"verbose", ADDED, nil},
			},
		},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s_%s", c.data1, c.data2)

		t.Run(name, func(t *testing.T) {

			got := CreateDiff(c.data1, c.data2)

			assert.Equal(t, c.want, got)
		})
	}
}
