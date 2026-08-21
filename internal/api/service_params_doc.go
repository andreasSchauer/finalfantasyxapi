package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type ParamsDoc struct {
	GeneralRules *string    `json:"general_rules"`
	Fields       []FieldDoc `json:"fields"`
}

type FieldDoc struct {
	Field           FieldName  `json:"field"`
	Type            string     `json:"type"`
	Required        bool       `json:"required"`
	RequiredOr      []FieldName   `json:"required_or"`
	ConflictsWith   []FieldName   `json:"conflicts_with,omitempty"`
	Usage           []string   `json:"usage,omitempty"` // won't be used for slices
	DefaultVal      any        `json:"default_val,omitempty"`
	MinVal          *int32     `json:"min_val,omitempty"`
	MaxVal          *int32     `json:"max_val,omitempty"`
	MaxArrayLen     *int       `json:"max_array_len,omitempty"`
	EnumValues      []string   `json:"enum_values,omitempty"`
	Description     string     `json:"description"`
	ChildProperties []FieldDoc `json:"child_properties,omitempty"`
}

func paramsDocToFieldMap(doc ParamsDoc) map[FieldName]FieldDoc {
	m := make(map[FieldName]FieldDoc, len(doc.Fields))

	for _, fieldEntry := range doc.Fields {
		m[fieldEntry.Field] = fieldEntry
	}

	return m
}

// some fields will be auto-generated later, like with queryParams
// before returning, I can iterate over the []FieldDoc and depending on the type, I will fill in things like usage, if it is empty, or min and max for ids
// so like a completeInit kind of function

func (cfg *Config) getAlBhedParamsDoc() ParamsDoc {
	return ParamsDoc{
		Fields: []FieldDoc{
			{
				Field:       pfnText,
				Type:        "string",
				Required:    true,
				Usage:       []string{"cysbma daqd", "c[am]bma da[x]d"},
				Description: "The text you want to be translated. You can wrap letters that are already translated (e.g. the red letters in the game's subtitles) into square brackets to keep them unchanged. By default, it is assumed that you want to translate English text into Al Bhed.",
			},
			{
				Field:       pfnDirection,
				Type:        "string (enum: translationDirection)",
				Usage:       []string{"to-al-bhed"},
				DefaultVal:  string(database.TranslationDirectionToAlBhed),
				EnumValues:  createEnumStringSlice(cfg.t.TranslationDirection.lookup),
				Description: "The translation direction.",
			},
		},
	}
}
