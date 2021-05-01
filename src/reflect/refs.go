package tateru

import (
	"reflect"
	"unicode"
)

func String(s string) *string { return &s }

func StructKeys(v reflect.Value, others ...string) map[string]bool {
	T := v.Type()
	L, values := v.NumField(), make(map[string]bool)
	for i := 0; i != L; i++ {
		values[ToSnakeCase(T.Field(i).Name)] = true
	}
	for i, L := 0, len(others); i != L; i++ { values[others[i]] = true }
	return values
}

func ToSnakeCase(s string) string {
	var res = make([]rune, 0, len(s))
	var p = '_'
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			res = append(res, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				res = append(res, '_', unicode.ToLower(r))
			} else {
				res = append(res, unicode.ToLower(r))
			}
		} else {
			res = append(res, unicode.ToLower(r))
		}

		p = r
	}
	return string(res)
}