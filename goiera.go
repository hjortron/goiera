package hiergo
import "strings"

type Config interface {
	GetValue(field string)interface{}
}

type fieldPath struct {
	path []string
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
	fp.path = fields[0:fieldsLen-1]

	return
}

func (c config) GetValue(field string) interface{} {
	return ""
}


