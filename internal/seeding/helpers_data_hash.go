package seeding

import (
	"fmt"
	"strings"
)

func generateDataHash(h Hashable) string {
	return combineFields(h.ToHashFields())
}

func generateJunctionHash(j Junction, desc string) string {
	return combineFields(j.ToHashFieldsJ(desc))
}

func (l *Lookup) getHashID(h Hashable) (int32, error) {
	id, ok := l.Hashes[generateDataHash(h)]
	if !ok {
		return 0, fmt.Errorf("no data hash available for %s", combineFields(h.ToHashFields()))
	}

	return id, nil
}


func combineFields(fields []any) string {
	var builder strings.Builder

	for i, field := range fields {
		if i > 0 {
			builder.WriteString("|")
		}

		if field == nil {
			builder.WriteString("NULL")
		} else {
			fmt.Fprint(&builder, field)
		}
	}

	return builder.String()
}