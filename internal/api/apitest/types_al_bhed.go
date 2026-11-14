package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"

type AlBhedResponse struct {
	testGeneral
	servicetypes.AlBhedResponse
}

func (e AlBhedResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e AlBhedResponse) Compare(test test, got servicetypes.AlBhedResponse) {
	test.t.Helper()
	test.compare("translated text", e.TranslatedText, got.TranslatedText)
	test.compare("original text", e.OriginalText, got.OriginalText)
	test.compare("direction", e.Direction, got.Direction)
}
