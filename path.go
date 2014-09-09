package path

import "strings"

// Path walks the dot-delimited `path` to return a nested map value, or nil.
func Path(m map[string]interface{}, path string) interface{} {
	var obj interface{} = m
	var val interface{} = nil

	parts := strings.Split(path, ".")
	for _, p := range parts {
		if v, ok := obj.(map[string]interface{}); ok {
			obj = v[p]
			val = obj
		} else {
			return nil
		}
	}

	return val
}
