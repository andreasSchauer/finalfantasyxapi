package api

import (
	"reflect"
)

func hasDefaultVal(doc FieldDoc) bool {
	return doc.DefaultVal != nil
}

func hasVal(fieldName FieldName, userValues map[FieldName]any) bool {
	jsonVal, ok := userValues[fieldName]
	return ok && jsonVal != nil
}

func valIsPointer(val any) bool {
	return reflect.ValueOf(val).Kind() == reflect.Pointer
}
