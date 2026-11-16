package initlookup

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

var AlBhedParamsSamples = []handlerinput.SampleParams[servicetypes.AlBhedParams]{
	{
		Description: "translate from english to al bhed.",
		Params: servicetypes.AlBhedParams{
			Text: "Father",
		},
	},
	{
		Description: "translate from english to al bhed, but keep the original name. You can also explicitly state the translation direction to be 'to-al-bhed', but it will be applied by default.",
		Params: servicetypes.AlBhedParams{
			Text:      "My name is [Tidus].",
			Direction: string(database.TranslationDirectionToAlBhed),
		},
	},
	{
		Description: "translate from al bhed to english.",
		Params: servicetypes.AlBhedParams{
			Text:      "cysbma daqd",
			Direction: string(database.TranslationDirectionToEnglish),
		},
	},
	{
		Description: "translate from al bhed to english while marking already translated letters (red letters in game).",
		Params: servicetypes.AlBhedParams{
			Text:      "am[i]q[ir]",
			Direction: string(database.TranslationDirectionToEnglish),
		},
	},
}
