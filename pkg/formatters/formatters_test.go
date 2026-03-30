package formatters

import (
	"fmt"
	"os"
	"testing"

	diffbuilder "github.com/bkoshelev/go-project-244/pkg/diff_builder"
	"github.com/bkoshelev/go-project-244/pkg/parsers"
	"github.com/stretchr/testify/assert"
)

func TestStylishFormatter(t *testing.T) {

	deep, _ := os.ReadFile("../../testdata/fixtures/deep.txt")

	data1, _ := parsers.ParseFile("./../../testdata/deep_1.json")
	data2, _ := parsers.ParseFile("./../../testdata/deep_2.json")
	diff := diffbuilder.CreateDiff(data1, data2)

	cases := []struct {
		diff []diffbuilder.Node
		want string
	}{
		{
			diff,
			string(deep),
		},
	}

	for _, c := range cases {
		name := fmt.Sprintf("StylishFormatter(%s)", c.diff)

		t.Run(name, func(t *testing.T) {
			got, _ := CreateStylishFormatter().Format(c.diff)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestPlainFormatter(t *testing.T) {

	deep_plain, _ := os.ReadFile("../../testdata/fixtures/deep_plain.txt")

	data1, _ := parsers.ParseFile("./../../testdata/deep_1.json")
	data2, _ := parsers.ParseFile("./../../testdata/deep_2.json")
	diff := diffbuilder.CreateDiff(data1, data2)

	cases := []struct {
		diff []diffbuilder.Node
		want string
	}{
		{
			diff,
			string(deep_plain),
		},
	}

	for _, c := range cases {
		name := fmt.Sprintf("PlainFormatter(%s)", c.diff)

		t.Run(name, func(t *testing.T) {
			got, _ := CreatePlainFormatter().Format(c.diff)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestJsonFormatter(t *testing.T) {

	deep_json, _ := os.ReadFile("../../testdata/fixtures/deep_json.txt")

	data1, _ := parsers.ParseFile("./../../testdata/deep_1.json")
	data2, _ := parsers.ParseFile("./../../testdata/deep_2.json")
	diff := diffbuilder.CreateDiff(data1, data2)

	cases := []struct {
		diff []diffbuilder.Node
		want string
	}{
		{
			diff,
			string(deep_json),
		},
	}

	for _, c := range cases {
		name := fmt.Sprintf("JsonFormatter(%s)", c.diff)

		t.Run(name, func(t *testing.T) {
			got, _ := CreateJsonFormatter().Format(c.diff)
			assert.Equal(t, c.want, got)
		})
	}
}
