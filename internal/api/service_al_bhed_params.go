package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type AlBhedParams struct {
	Text      string `json:"text,omitempty"`
	Direction string `json:"direction,omitempty"`
}

func getAlBhedParamsDoc(server *Server) ParamsDoc {
	return ParamsDoc{
		Fields: []FieldDoc{
			{
				Field:       ts.PfnText,
				Type:        "string",
				ExampleUses: []string{"cysbma daqd", "c[am]bma da[x]d"},
				Required:    true,
				Description: "The text you want to be translated. You can wrap letters that are already translated (e.g. the red letters in the game's subtitles) into square brackets to keep them unchanged. By default, it is assumed that you want to translate English text into Al Bhed.",
			},
			{
				Field:       ts.PfnDirection,
				Type:        "string (enum: translationDirection)",
				DefaultVal:  string(database.TranslationDirectionToAlBhed),
				EnumValues:  enums.T.TranslationDirection.Strings(),
				Description: "The translation direction.",
			},
		},
	}
}
