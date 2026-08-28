package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AlBhedParams struct {
	Text      string `json:"text"`
	Direction string `json:"direction,omitempty"`
}

func (p AlBhedParams) GetDoc(cfg *Config) ParamsDoc {
	return cfg.getAlBhedParamsDoc()
}

func (cfg *Config) getAlBhedParamsDoc() ParamsDoc {
	return ParamsDoc{
		Fields: []FieldDoc{
			{
				Field:       pfnText,
				Type:        "string",
				ExampleUses: []string{"cysbma daqd", "c[am]bma da[x]d"},
				Required:    true,
				Description: "The text you want to be translated. You can wrap letters that are already translated (e.g. the red letters in the game's subtitles) into square brackets to keep them unchanged. By default, it is assumed that you want to translate English text into Al Bhed.",
			},
			{
				Field:       pfnDirection,
				Type:        "string (enum: translationDirection)",
				DefaultVal:  string(database.TranslationDirectionToAlBhed),
				EnumValues:  createEnumStringSlice(cfg.t.TranslationDirection.lookup),
				Description: "The translation direction.",
			},
		},
	}
}