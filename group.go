package tlog

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func Group[T any](root string, group *T) T {
	if group == nil {
		panic("tlog: nil group")
	}

	v := reflect.ValueOf(group).Elem()
	if v.Kind() != reflect.Struct {
		panic("tlog: group must point to a struct")
	}

	fillGroup(v, Op(root))
	return *group
}

func kebabCase(s string) string {
	var b strings.Builder
	runes := []rune(s)

	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) {
			prev := runes[i-1]
			nextIsLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if unicode.IsLower(prev) || nextIsLower {
				b.WriteByte('-')
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}

	return b.String()
}

func fillGroup(v reflect.Value, prefix Op) {
	t := v.Type()
	opType := reflect.TypeFor[Op]()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if !value.CanSet() {
			continue
		}

		name := field.Tag.Get("tlog")
		if name == "" {
			name = kebabCase(field.Name)
		}

		switch field.Type {
		case opType:
			value.Set(reflect.ValueOf(prefix.Extend(name)))
		default:
			if value.Kind() == reflect.Struct {
				fillGroup(value, prefix.Extend(name))
				continue
			}

			panic(fmt.Sprintf(
				`tlog: field %q has unsupported type %s; expected tlog.Op or nested struct`,
				field.Name,
				field.Type,
			))
		}
	}
}
