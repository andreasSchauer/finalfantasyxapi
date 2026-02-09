package seeding

import (
	"crypto/sha256"
	"fmt"
)

// ToHashFields simply puts all values of all fields of an entity into an []any slice (including nil)
type Hashable interface {
	ToHashFields() []any
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
