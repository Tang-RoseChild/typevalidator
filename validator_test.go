package typevalidator

import (
	"testing"
)

type testdata struct {
	a, b string
}

func TestTypeValid(t *testing.T) {

	testData := struct {
		name interface{} "typevalidator.testdata"
		age  interface{}
	}{
		name: testdata{},
		age:  testdata{},
	}

	if TypeValid(testData) == false {
		t.Error("should return  true")
	}

	// fundermental type

	fT := struct {
		a   interface{} "int"
		b   interface{} "bool"
		f   interface{} "float64"
		s   interface{} "string"
		m   interface{} "map[string]string"
		mi  interface{} "map[int]interface {}"
		sl  interface{} "[]string"
		sli interface{} "[]interface{}"
		arr interface{} "[3]int"
	}{
		a:   int(12),
		b:   false,
		f:   float64(32),
		s:   "abc",
		m:   map[string]string{"a": "a"},
		mi:  map[int]interface{}{1: "abc", 2: true},
		sl:  []string{"a"},
		sli: []interface{}{1, "ab", true},
		arr: [3]int{1},
	}

	if TypeValid(fT) == false {
		t.Error("should return true for fT")
	}

	// ptr struct

	type pname struct {
		name interface{} "string"
	}

	if TypeValid(&pname{name: "abc"}) == false {
		t.Error("should return true for fT")
	}

	if TypeValid(&pname{name: 123}) == true {
		t.Error("should return false ")
	}

	// not all interface fields

	niT := struct {
		a interface{} "map[int]string"
		b bool
	}{
		a: map[int]string{1: ""},
		b: true,
	}

	if TypeValid(niT) == false {
		t.Error("should return true for fT")
	}

	// multi type check
	mtT := struct {
		a interface{} "map[int]string, map[int]int"
	}{
		a: map[int]int{1: 2},
	}

	if TypeValid(mtT) == false {
		t.Error("should return true for fT")
	}

	mtT.a = map[int]string{1: "st"}
	if TypeValid(mtT) == false {
		t.Error("should return true for fT")
	}
	mtT.a = "abc"
	if TypeValid(mtT) == true {
		t.Error("should return false for fT")
	}
}
