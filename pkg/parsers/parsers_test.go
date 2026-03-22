package parsers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJsonFile(t *testing.T) {
	cases := []struct {
		filepath string
		error    error
		want     map[string]any
	}{
		{"./../../testdata/file1.json", nil, map[string]any{"follow": false, "host": "hexlet.io", "proxy": "123.234.53.22", "timeout": 50.0}},
		{"./../../testdata/flat1.jso", errors.New("Файл не существует"), map[string]any{}},
		{"./../../testdata/incorrect.json", errors.New("Содержимое файла не соответствует типу файла"), map[string]any{}},
		{"./../../testdata/empty_file.json", errors.New("Содержимое файла не соответствует типу файла"), map[string]any{}},
		{"./../../testdata", errors.New("Путь должен вести к файлу, а не к папке"), map[string]any{}},
		{"./../../testdata/unsupported_ext.unsupported", errors.New("Тип файла не поддерживается"), map[string]any{}},
	}

	for _, c := range cases {
		name := fmt.Sprintf("ParseJsonFile (%s)", c.filepath)

		t.Run(name, func(t *testing.T) {

			got, e := ParseFile(c.filepath)

			if c.error == nil && e != nil {
				assert.NoError(t, e)
			}

			if c.error != nil && e == nil {
				assert.Error(t, e)
			}

			assert.Equal(t, c.want, got)
		})
	}
}
