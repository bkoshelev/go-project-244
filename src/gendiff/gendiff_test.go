package code

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenDiff(t *testing.T) {
	cases := []struct {
		filepath1 string
		filepath2 string
		format    string
		error     error
		want      string
	}{
		{
			"./../../testdata/file1.json",
			"./../../testdata/file2.json",
			"",
			nil,
			`{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s_%s_%s", c.filepath1, c.filepath2, c.format)

		t.Run(name, func(t *testing.T) {

			got, e := GenDiff(c.filepath1, c.filepath2, c.format)

			if c.error == nil && e != nil {
				t.Errorf("Не ожидали ошибку %v", e)
				panic("")
			}

			if c.error != nil && e == nil {
				t.Errorf("Ожидалась ошибка")
				panic("")
			}

			if !strings.EqualFold(got, c.want) {
				t.Errorf("GenDiff(%s, %s, %s) = \n%v\n, хотели \n%v", c.filepath1, c.filepath2, c.format, got, c.want)
			}
		})
	}
}
