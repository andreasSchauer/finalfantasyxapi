package seeding

import (
	"crypto/sha256"
	"fmt"
)

func generateDataHash(h Hashable) string {
	fields := h.ToHashFields()
	combined := combineFields(fields)
	return combined
	//hash := sha256.Sum256([]byte(combined))
	//return fmt.Sprintf("%x", hash)
}

func generateJunctionHash(junction Junction, name string) string {
	fields := junction.ToHashFieldsJ(name)
	combined := combineFields(fields)
	hash := sha256.Sum256([]byte(combined))
	return fmt.Sprintf("%x", hash)
}

func (l *Lookup) getHashID(h Hashable) (int32, error) {
	id, ok := l.Hashes[generateDataHash(h)]
	if !ok {
		return 0, fmt.Errorf("no data hash available for %s", combineFields(h.ToHashFields()))
	}

	return id, nil
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