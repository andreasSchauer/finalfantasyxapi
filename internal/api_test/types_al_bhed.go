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
	compare(test, "translated text", exp.TranslatedText, got.TranslatedText)
	compare(test, "original text", exp.OriginalText, got.OriginalText)
	compare(test, "direction", exp.Direction, got.Direction)
}
