package helpers

import (
	"fmt"
	"strings"
)

func CombineFields(fields []any) string {
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