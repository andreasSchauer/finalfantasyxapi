package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type SampleParams[P ServiceParams] struct {
	Description 	string		`json:"description"`
	Params			P			`json:"params"`
}

func getAlBhedParamsSamples() []SampleParams[AlBhedParams] {
	return []SampleParams[AlBhedParams]{
		{
			Description: "translate from english to al bhed.",
			Params: AlBhedParams{
				Text: "Father",
			},
		},
		{
			Description: "translate from english to al bhed, but keep the original name. You can also explicitly state the translation direction to be 'to-al-bhed', but it will be applied by default.",
			Params: AlBhedParams{
				Text: "My name is [Tidus].",
				Direction: string(database.TranslationDirectionToAlBhed),
			},
		},
		{
			Description: "translate from al bhed to english.",
			Params: AlBhedParams{
				Text: "cysbma daqd",
				Direction: string(database.TranslationDirectionToEnglish),
			},
		},
		{
			Description: "translate from al bhed to english while marking already translated letters (red letters in game).",
			Params: AlBhedParams{
				Text: "am[i]q[ir]",
				Direction: string(database.TranslationDirectionToEnglish),
			},
		},
	}
}