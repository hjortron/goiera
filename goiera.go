package goiera

import (
	"strings"

	"github.com/ghodss/yaml"
)

type Number string
type Bool string

type Config interface {
	GetValue(field string) interface{}
	GetString(field string) []string
}

type Value interface{}

type fieldPath struct {
	path      []string
	fieldName string
}

type config map[string]interface{}

func UnmarshalConfig(cbytes []byte) Config {
	conf := config{}
	yaml.Unmarshal(cbytes, &conf)
	return conf
}

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

func (c config) GetValue(field string) (value interface{}) {
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
				t = nmap
				if ok {
					value = v
				}
			}
		} else {
			return
		}
	}
	return
}

func (c config) GetString(field string) []string {
	f := c.GetValue(field)
	switch f.(type) {
	case []interface{}:
		b := make([]string, len(f.([]interface{})))
		for i := range f.([]interface{}) {
			b[i] = f.([]interface{})[i].(string)
		}
		return b
	case interface{}:
		return []string{f.(string)}
	default:
		return []string{}
	}
}
