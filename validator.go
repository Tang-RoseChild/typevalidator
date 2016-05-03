package typevalidator

import (
	"reflect"
	"strings"
)

// maybe a struct or ptr
func TypeValid(iface interface{}) bool {
	rType := reflect.TypeOf(iface)
	rValue := reflect.ValueOf(iface)

	if rType.Kind() == reflect.Struct {
		return structValid(rValue.Interface())
	}

	if rType.Kind() == reflect.Ptr {
		if rValue.Elem().Type().Kind() != reflect.Struct {
			return false
		}
		return structValid(rValue.Elem().Interface())
	}
	return false
}

// params always is struct type .
// p type is struct ,all fields interface{} => all tags
// just used in case of all fields are interface{}
func structValid(iface interface{}) bool {
	rValue := reflect.ValueOf(iface)

	for i := 0; i < rValue.NumField(); i++ {
		rTag := rValue.Type().Field(i).Tag
		if len(string(rTag)) == 0 { // no tag , pass type check
			continue
		}

		if rValue.Field(i).Type().Kind() != reflect.Interface { // not interface, pass type check.this check will be done when compiling
			continue
		}

		elem := rValue.Field(i).Elem()
		typeName := elem.Type().String()
		if containTag(string(rTag), typeName) == false {
			return false
		}

	}

	return true

}

// tags will be splited by ","
func containTag(tag, cont string) bool {
	cont = strings.Replace(cont, " ", "", -1)
	contained := false
	tags := strings.Split(tag, ",")
	for _, t := range tags {
		t := strings.Replace(t, " ", "", -1)
		if t == cont {
			contained = true
			break
		}

	}
	return contained
}
