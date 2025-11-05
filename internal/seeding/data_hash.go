package seeding

import (
	"crypto/sha256"
	"fmt"
)

type Hashable interface {
	ToHashFields() []any
}

// ToHashFields simply puts all values of all fields of an entity into an []any slice (including nil)
// Use derefOrNil for every field that is a pointer

func derefOrNil[T any](ptr *T) any {
	if ptr == nil {
		return nil
	}
	return *ptr
}

// used to get the id of objects that can be null and are thus pointers
func ObjPtrToID[T HasID](objPtr *T) any {
	if objPtr == nil {
		return nil
	}

	obj := *objPtr
	id := obj.GetID()

	if id == 0 {
		return nil
	}
	return id
}

func generateDataHash(h Hashable) string {
	fields := h.ToHashFields()
	combined := combineFields(fields)
	hash := sha256.Sum256([]byte(combined))
	return fmt.Sprintf("%x", hash)
}

func combineFields(fields []any) string {
	var combined string

	for i, field := range fields {
		if i > 0 {
			combined += "|"
		}

		if field == nil {
			combined += "NULL"
		} else {
			combined += fmt.Sprintf("%v", field)
		}
	}

	return combined
}
