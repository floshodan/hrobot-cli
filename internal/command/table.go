package command

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type TableWriter interface {
	Render()
}

type Table struct {
	Headers []string
	Rows    [][]string
}

func (t *Table) AddRow(values ...string) {
	t.Rows = append(t.Rows, values)
}

func NewTable(data interface{}, fields ...string) *Table {
	headers := make([]string, len(fields))
	for i, f := range fields {
		headers[i] = strings.Title(f)
	}

	var rows [][]string

	return &Table{
		Headers: headers,
		Rows:    rows,
	}
}

func (t *Table) Renderer() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(t.Headers)
	table.AppendBulk(t.Rows)
	table.Render()
}

func getField(v interface{}, field string) (string, error) {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	fieldType := f.Type().Name()
	fmt.Println(fieldType)
	switch fieldType {
	case "int":
		return strconv.FormatInt(f.Int(), 10), nil
	case "string":
		return f.String(), nil
	default:
		return "", fmt.Errorf("Could not convert error")
	}
}
