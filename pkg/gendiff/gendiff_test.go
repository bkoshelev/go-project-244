package code

import (
	"errors"
	"fmt"
	"os"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	file, _ := os.ReadFile("../../testdata/fixtures/file.txt")
	empty, _ := os.ReadFile("../../testdata/fixtures/empty.txt")

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
			string(file),
		},
		{
			"./../../testdata/empty1.json",
			"./../../testdata/empty2.json",
			"",
			nil,
			string(empty),
		},
		{
			"./unknown.json",
			"./../../testdata/file2.json",
			"",
			errors.New(""),
			"",
		},
		{
			"./../../testdata/file1.json",
			"./unknown.json",
			"",
			errors.New(""),
			"",
		},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s_%s_%s", c.filepath1, c.filepath2, c.format)

		t.Run(name, func(t *testing.T) {

			got, e := GenDiff(c.filepath1, c.filepath2, c.format)

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
