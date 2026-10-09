package api

import (
	"fmt"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ParamsDoc struct {
	GeneralRules *string    `json:"general_rules"`
	Fields       []FieldDoc `json:"fields"`
}

type FieldDoc struct {
	Field             FieldName           `json:"field"`
	Type              string              `json:"type"`
	AllowZeroVal      bool                `json:"-"`
	ExampleUses       []string            `json:"example_uses,omitempty"`
	Required          bool                `json:"required"`
	RequiredOr        []FieldName         `json:"required_or,omitempty"`
	RequiresAll       []FieldName         `json:"requires_all,omitempty"`
	RequiresOne       []FieldName         `json:"requires_one,omitempty"`
	ConflictsWith     []FieldName         `json:"conflicts_with,omitempty"`
	DefaultVal        any                 `json:"default_val,omitempty"`
	MinVal            *int32              `json:"min_val,omitempty"`
	MaxVal            *int32              `json:"max_val,omitempty"`
	MaxArrayLen       *int                `json:"max_array_len,omitempty"`
	AllowedIDs        []int32             `json:"allowed_ids,omitempty"`
	AllowedIdValPairs []AllowedIdValPairs `json:"allowed_id_val_pairs"`
	EnumValues        []string            `json:"enum_values,omitempty"`
	Description       string              `json:"description"`
	ChildProps        []FieldDoc          `json:"child_properties,omitempty"`
	ChildMinVal       *int32              `json:"child_min_val,omitempty"`
	ChildMaxVal       *int32              `json:"child_max_val,omitempty"`
}

type AllowedIdValPairs struct {
	ID     int32   `json:"id"`
	Values []int32 `json:"values"`
	MaxVal int32   `json:"-"`
}

func (e AllowedIdValPairs) format() string {
	return fmt.Sprintf("%d - %s", e.ID, h.IntSliceToString(e.Values))
}

func formatAllowedIdValPairs(entries []AllowedIdValPairs) string {
	slice := make([]string, len(entries))

	for i := range entries {
		slice[i] = entries[i].format()
	}

	return strings.Join(slice, " | ")
}
