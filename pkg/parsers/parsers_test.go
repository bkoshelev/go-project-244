package parsers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	cases := []struct {
		filepath string
		error    error
		want     map[string]any
	}{
		{"./../../testdata/file1.json", nil, map[string]any{"follow": false, "host": "hexlet.io", "proxy": "123.234.53.22", "timeout": 50.0}},
		{"./../../testdata/flat1.jso", errors.New("Файл не существует"), nil},
		{"./../../testdata/incorrect.json", errors.New("Содержимое файла не соответствует типу файла"), nil},
		{"./../../testdata/empty_file.json", errors.New("Содержимое файла не соответствует типу файла"), nil},
		{"./../../testdata", errors.New("Путь должен вести к файлу, а не к папке"), nil},
		{"./../../testdata/unsupported_ext.unsupported", errors.New("Тип файла не поддерживается"), nil},
	}

	for _, c := range cases {
		name := fmt.Sprintf("ParseFile (%s)", c.filepath)

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
