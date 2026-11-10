package servicetypes

type AlBhedResponse struct {
	TranslatedText string `json:"translated_text"`
	OriginalText   string `json:"original_text"`
	Direction      string `json:"direction"`
}