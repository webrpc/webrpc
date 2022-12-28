package gen

import (
	"fmt"
)

// Create new dictionary.
func dict(pairs ...interface{}) map[string]interface{} {
	if len(pairs)%2 == 1 {
		panic("dict must be created with even number of parameters (key:value pairs)")
	}

	m := map[string]interface{}{}
	for i := 0; i < len(pairs); i += 2 {
		key, ok := pairs[i].(string)
		if !ok {
			panic(fmt.Sprintf("dict argument(%v) must be string key", i))
		}
		m[key] = pairs[i+1]
	}

	return m
}

func get(m map[string]interface{}, key interface{}) (interface{}, error) {
	keyStr := toString(key)

	val, ok := m[keyStr]
	if !ok {
		return nil, fmt.Errorf("get(): key %q doesn't exist", keyStr)
	}

	return val, nil
}

func set(m map[string]interface{}, key string, value interface{}) string {
	m[key] = value
	return ""
}

// TODO: Support slices too?
func exists(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}
