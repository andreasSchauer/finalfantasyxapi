package api

type ParamsDoc struct {
	GeneralRules *string    `json:"general_rules"`
	Fields       []FieldDoc `json:"fields"`
}


type FieldDoc struct {
	Field         FieldName   `json:"field"`
	Type          string      `json:"type"`
	ExampleUses   []string    `json:"example_uses,omitempty"`
	Required      bool        `json:"required"`
	RequiredOr    []FieldName `json:"required_or,omitempty"`
	ConflictsWith []FieldName `json:"conflicts_with,omitempty"`
	DefaultVal    any         `json:"default_val,omitempty"`
	MinVal        *int32      `json:"min_val,omitempty"`
	MaxVal        *int32      `json:"max_val,omitempty"`
	MaxArrayLen   *int        `json:"max_array_len,omitempty"`
	EnumValues    []string    `json:"enum_values,omitempty"`
	Description   string      `json:"description"`
	ChildProps    []FieldDoc  `json:"child_properties,omitempty"`
}