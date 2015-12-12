package hiergo

import (
	"github.com/ghodss/yaml"
	"reflect"
	"testing"
)

var guardians = []byte(`quotes:
- "All heroes start somewhere."
- "When things get bad, they'll do their worst."
Star-Lord:
    quotes:
    - "Bereet! Okay, I'm gonna be totally honest with you ... I forgot you were here."
    - "Hey, you know what? There's another name you might know me by! ...Star-Lord."
Gamora:
    quotes:
    - "I am going to die surrounded by the biggest idiots in the galaxy."
    - "We're just like Kevin Bacon!"
"Rocket Raccoon":
    quotes:
    - "Ain't no thing like me, 'cept me."
    - "I'm giving you to the count of five. Five, four, three--"
`)

func Test_GetChildValue(t *testing.T) {
	conf := config{}
	yaml.Unmarshal(guardians, &conf)
	quotes := conf.getString("Gamora:quotes")
	if !reflect.DeepEqual(quotes, []string{"I am going to die surrounded by the biggest idiots in the galaxy.", "We're just like Kevin Bacon!"}) {
		t.Error(quotes)
	}
	quotes = conf.getString("Drax:quotes")
	if !reflect.DeepEqual(quotes, []string{"All heroes start somewhere.", "When things get bad, they'll do their worst."}) {
		t.Error(quotes)
	}
}

func Test_GetFieldPath(t *testing.T) {
	testFields := map[string]fieldPath{}
	testFields["child:field"] = fieldPath{fieldName: "field", path: []string{"child"}}
	testFields["field"] = fieldPath{fieldName: "field", path: []string{}}
	testFields["child1:child2:field"] = fieldPath{fieldName: "field", path: []string{"child1", "child2"}}
	testFields[""] = fieldPath{}
	for path, v := range testFields {
		newFP := newFieldPath(path)
		if !reflect.DeepEqual(newFP, v) {
			t.Error(newFP, v)
		}
	}
}
