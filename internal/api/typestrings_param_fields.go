package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type FieldName string

const (
	pfnDirection FieldName = "direction"
	pfnText      FieldName = "text"
)

func formatPfnSlice(pfns []FieldName) string {
	if pfns == nil {
		return ""
	}

	strings := []string{}

	for _, qpn := range pfns {
		strings = append(strings, string(qpn))
	}

	return h.FormatStringSlice(strings)
}