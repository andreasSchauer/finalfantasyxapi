package api

import (
	"reflect"
)

func getInt32FromValMap(fieldName FieldName, valueMap map[FieldName]any) (int32, error) {
	raw, ok := valueMap[fieldName]
	if !ok {
		return 0, errNoVal
	}
	intVal := int32(raw.(float64))

	return intVal, nil
}

func hasDefaultVal(doc FieldDoc) bool {
	return doc.DefaultVal != nil
}

func hasVal(fieldName FieldName, valueMap map[FieldName]any) bool {
	jsonVal, ok := valueMap[fieldName]
	return ok && jsonVal != nil
}

func hasValOld(val any, doc FieldDoc) bool {
	if val == nil {
		return false
	}

	if doc.AllowZeroVal {
		return true
	}

	switch t := val.(type) {
		case string: return t != ""
		case int32: return t != 0
		case int: return t != 0
		case float64: return t != 0
		case bool: return t
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Pointer:
		return !v.IsNil()

	case reflect.Slice:
		return !(v.IsNil() || v.Len() == 0)

	case reflect.Struct:
		return !v.IsZero()
	}

	return false
}

func valIsPointer(val any) bool {
	return reflect.ValueOf(val).Kind() == reflect.Pointer
}