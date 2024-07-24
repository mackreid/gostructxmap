package gostructxmap

import (
	"errors"
	"reflect"
	"strings"
)

const (
	PK_TAG = "sxm"
)

type FieldRule func(value any, key string, opts ...string) bool

func WithOmit() FieldRule {
	return func(value any, key string, opts ...string) bool {
		if doesContain("omitempty", opts...) {
			if reflect.ValueOf(value).IsZero() || reflect.ValueOf(value).IsNil() {
				return false
			}
		}
		return true
	}
}

func WithMask(mask ...string) FieldRule {
	return func(value any, key string, opts ...string) bool {
		if doesContain("mask", opts...) {
			if doesContain(key, mask...) {
				return true
			}
		}
		return false
	}
}

func Make(structToMap any, rules ...FieldRule) (map[string]any, error) {
	mappedStruct := make(map[string]any)
	structToMapValue := reflect.ValueOf(structToMap)
	if structToMapValue.Kind() == reflect.Pointer {
		structToMapValue = structToMapValue.Elem()
	}
	if structToMapValue.Kind() != reflect.Struct {
		return nil, errors.New("structToMap object must be a struct or a pointer to a struct")
	}
	structToMapType := structToMapValue.Type()
fields:
	for i := 0; i < structToMapType.NumField(); i++ {
		fieldValue := structToMapValue.Field(i)
		fieldType := structToMapType.Field(i)
		tag := fieldType.Tag.Get(PK_TAG)
		if doesSkip(tag) {
			continue fields
		}
		t_opts := strings.Split(tag, ",")
		key := t_opts[0]
		value := fieldValue.Interface()
		for _, rule := range rules {
			if !rule(value, t_opts[0], t_opts[1:]...) {
				continue fields
			}
		}
		mappedStruct[key] = value
	}
	return mappedStruct, nil
}

func doesSkip(v string) bool {
	if v == "-" || v == "" {
		return true
	}
	return false
}

func doesContain(key string, opts ...string) bool {
	for _, s := range opts {
		if s == key {
			return true
		}
	}
	return false
}
