package api

type expAlBhedResponse struct {
	testGeneral
	AlBhedResponse
}

func (e expAlBhedResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAlBhedResponses(test test, exp expAlBhedResponse, got AlBhedResponse) {
	test.t.Helper()
	compare(test, "translated text", exp.TranslatedText, got.TranslatedText)
	compare(test, "original text", exp.OriginalText, got.OriginalText)
	compare(test, "direction", exp.Direction, got.Direction)
}