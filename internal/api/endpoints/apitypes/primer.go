package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type Primer struct {
	ID            int32                 `json:"id"`
	Name          string                `json:"name"`
	KeyItem       apiresource.Named     `json:"key_item"`
	Description   string                `json:"description"`
	AlBhedLetter  string                `json:"al_bhed_letter"`
	EnglishLetter string                `json:"english_letter"`
	Areas         []apiresource.Area    `json:"areas"`
	Treasures     []apiresource.Unnamed `json:"treasures"`
}
