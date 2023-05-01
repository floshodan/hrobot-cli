package output

import (
	"io"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
	"unicode"
)

type FieldFn func(obj interface{}) string

type writerFlusher interface {
	io.Writer
	Flush() error
}

type Table struct {
	w             writerFlusher
	columns       map[string]bool
	fieldMapping  map[string]FieldFn
	fieldAlias    map[string]string
	allowedFields map[string]bool
}

// NewTable creates a new Table.
func NewTable() *Table {
	return &Table{
		w:             tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0),
		columns:       map[string]bool{},
		fieldMapping:  map[string]FieldFn{},
		fieldAlias:    map[string]string{},
		allowedFields: map[string]bool{},
	}
}

// AddFieldFn adds a function which handles the output of the specified field.
func (o *Table) AddFieldFn(field string, fn FieldFn) *Table {
	o.fieldMapping[field] = fn
	o.allowedFields[field] = true
	o.columns[field] = true
	return o
}

// AddAllowedFields reads all first level fieldnames of the struct and allows them to be used.
func (o *Table) AddAllowedFields(obj interface{}) *Table {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		panic("AddAllowedFields input must be a struct")
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		k := t.Field(i).Type.Kind()
		if k != reflect.Bool &&
			k != reflect.Float32 &&
			k != reflect.Float64 &&
			k != reflect.String &&
			k != reflect.Int {
			// only allow simple values
			// complex values need to be mapped via a FieldFn
			continue
		}
		o.allowedFields[strings.ToLower(t.Field(i).Name)] = true
		o.allowedFields[fieldName(t.Field(i).Name)] = true
		o.columns[fieldName(t.Field(i).Name)] = true
	}
	return o
}

func fieldName(name string) string {
	r := []rune(name)
	var out []rune
	for i := range r {
		if i > 0 && (unicode.IsUpper(r[i])) && (i+1 < len(r) && unicode.IsLower(r[i+1]) || unicode.IsLower(r[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(r[i]))
	}
	return string(out)
}
