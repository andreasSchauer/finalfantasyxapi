package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"

type expAlBhedResponse struct {
	testGeneral
	servicetypes.AlBhedResponse
}

func (e expAlBhedResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAlBhedResponses(test test, exp expAlBhedResponse, got servicetypes.AlBhedResponse) {
	test.t.Helper()
	test.compare("translated text", exp.TranslatedText, got.TranslatedText)
	test.compare("original text", exp.OriginalText, got.OriginalText)
	test.compare("direction", exp.Direction, got.Direction)
}
