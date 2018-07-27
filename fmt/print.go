package fmt

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/aktoro-lang/container/dict"
	"github.com/aktoro-lang/types"
)

var variantType = reflect.TypeOf((*types.AkVariant)(nil)).Elem()

var dictType = reflect.TypeOf((*dict.Dict)(nil)).Elem()

func Println(args ...interface{}) interface{} {
	fmt.Println(Sprint(args...))
	return nil
}

func Sprint(args ...interface{}) string {
	buffer := []string{}

	for _, arg := range args {
		val := reflect.ValueOf(arg)
		if val.Type() == dictType {
			buffer = append(buffer, fmt.Sprint(arg))
		} else if val.Type().Implements(variantType) {
			buffer = append(buffer, sprintVariantConstructor(val))
		} else if val.Kind() == reflect.Kind(25) { // typ.Kind == Struct
			buffer = append(buffer, sprintStruct(val))
		} else {
			buffer = append(buffer, fmt.Sprint(arg))
		}
	}
	return strings.Join(buffer, " ")
}

func sprintStruct(val reflect.Value) string {
	buffer := []string{}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		fmtString := toSnake(typ.Field(i).Name) + ": " + Sprint(f.Interface())
		buffer = append(buffer, fmtString)
	}
	str := strings.Join(buffer, ", ")
	return "{" + str + "}"
}

func sprintVariantConstructor(val reflect.Value) string {
	buffer := []string{}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		buffer = append(buffer, Sprint(f.Interface()))
	}
	str := strings.Join(buffer, " ")
	constructor := typ.String()
	split := strings.Split(constructor, ".")
	name := split[len(split)-1]
	return name + " " + str
}

func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
