package parser

import (
	"fmt"
	"maps"
	"reflect"
	"testing"
)

func TestParseJsonFile(t *testing.T) {
	cases := []struct {
		filepath string
		error    error
		want     map[string]any
	}{
		{"./../../testdata/file1.json", nil, map[string]any{"follow": false, "host": "hexlet.io", "proxy": "123.234.53.22", "timeout": float64(50)}},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s", c.filepath)

		t.Run(name, func(t *testing.T) {

			got, e := ParseJsonFile(c.filepath)

			if c.error == nil && e != nil {
				t.Errorf("Не ожидали ошибку %v", e)
				panic("")
			}

			if c.error != nil && e == nil {
				t.Errorf("Ожидалась ошибка")
				panic("")
			}

			if !maps.Equal(got, c.want) {

				for k, v := range got {
					fmt.Printf("%s %v \n", k, reflect.TypeOf(v))
				}

				for k, v := range c.want {
					fmt.Printf("%s %v \n", k, reflect.TypeOf(v))
				}
				t.Errorf("ReadFile(%s) = %v, хотели %v", c.filepath, got, c.want)
			}
		})
	}
}
