package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"

type AlBhedResponse struct {
	TestGeneral
	servicetypes.AlBhedResponse
}

func (e AlBhedResponse) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e AlBhedResponse) Compare(test Test, got servicetypes.AlBhedResponse) {
	test.T.Helper()
	test.compare("translated text", e.TranslatedText, got.TranslatedText)
	test.compare("original text", e.OriginalText, got.OriginalText)
	test.compare("direction", e.Direction, got.Direction)
}
