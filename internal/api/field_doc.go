package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"


type FieldDoc struct {
	Field			string		`json:"field"`
	Type			string		`json:"type"`
	Required		bool		`json:"required"`
	ConflictsWith	[]string	`json:"conflicts_with,omitempty"`
	// need something for if at least one of a set of fields must be active
	Usage			[]string	`json:"usage,omitempty"` 		 // won't be used for slices
	DefaultVal		any			`json:"default_val,omitempty"`
	MinVal			*int32		`json:"min_val,omitempty"`
	MaxVal			*int32		`json:"max_val,omitempty"`
	EnumValues		[]string	`json:"enum_values,omitempty"`
	Description		string		`json:"description"`
	ChildProperties	[]FieldDoc	`json:"child_properties,omitempty"`
}


func (cfg *Config) getAlBhedFieldDoc() []FieldDoc {
	return []FieldDoc{
		{
			Field: "text",
			Type: "string",
			Required: true,
			Usage: []string{"cysbma daqd", "c[am]bma da[x]d"},
			Description: "The text you want to be translated. You can wrap letters that are already translated (e.g. the red letters in the game's subtitles) into square brackets to keep them unchanged. By default, it is assumed that you want to translate English text into Al Bhed.",
		},
		{
			Field: "direction",
			Type: "string (enum: translationDirection)",
			Required: false,
			Usage: []string{"to-al-bhed"}, // will be auto-generated later, like with queryParams
			DefaultVal: string(database.TranslationDirectionToAlBhed),
			EnumValues: createEnumStringSlice(cfg.t.TranslationDirection.lookup),
			Description: "The translation direction.",
		},
	}
}

func fieldDocSliceToMap(s []FieldDoc) map[string]FieldDoc {
	m := make(map[string]FieldDoc, len(s))

	for _, fieldEntry := range s {
		m[fieldEntry.Field] = fieldEntry
	}

	return m
}