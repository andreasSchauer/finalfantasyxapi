package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)

func compareAlBhedResponses(test Test, exp AlBhedResponse, got servicetypes.AlBhedResponse) {
	test.t.Helper()
	test.compare("translated text", exp.TranslatedText, got.TranslatedText)
	test.compare("original text", exp.OriginalText, got.OriginalText)
	test.compare("direction", exp.Direction, got.Direction)
}
