package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expAlBhedResponse struct {
	testGeneral
	api.AlBhedResponse
}

func (e expAlBhedResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAlBhedResponses(test test, exp expAlBhedResponse, got api.AlBhedResponse) {
	test.t.Helper()
	test.compare("translated text", exp.TranslatedText, got.TranslatedText)
	test.compare("original text", exp.OriginalText, got.OriginalText)
	test.compare("direction", exp.Direction, got.Direction)
}
