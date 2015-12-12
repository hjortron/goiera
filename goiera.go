package hiergo

import (
	"strings"
	"encoding/json"
	"log"
)

type Number string
type Bool string

type Config interface {
	GetValue(field string) interface{}
}

type Value interface{}

type fieldPath struct {
	path      []string
	fieldName string
}

type config map[string]interface{}

func newFieldPath(field string) (fp fieldPath) {
	fields := strings.Split(field, SEPARATOR)
	fieldsLen := len(fields)
	if fieldsLen == 0 || fields[0] == "" {
		return
	}
	fp.fieldName = fields[fieldsLen-1]
	fp.path = fields[0 : fieldsLen-1]

	return
}

func (c config) GetValue(field string) (value Value) {
	fp := newFieldPath(field)
	t := c
	f, ok := t[fp.fieldName]
	if ok {
		value = f
	}
	for _, node := range fp.path {
		n, ok := t[node]
		if ok {
			if nmap, ok := n.(map[string]interface{}); ok {
				v, ok := nmap[fp.fieldName]
				if ok {
					t = nmap
					value = v
				}
			}
		} else {
			return
		}
	}
	return
}

func (c config) getString(field string) []string {
	f := c.GetValue(field)
	switch f.(type) {
	case []interface{}:
		b := make([]string, len(f.([]interface{})))
		for i := range f.([]interface{}) {
			b[i] = f.([]interface{})[i].(string)
		}
		log.Println(b)
		return b
	case interface{}:
		return []string{f.(string)}
	default:
		return []string{}
	}
}