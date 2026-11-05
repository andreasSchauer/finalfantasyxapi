package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type Primer struct {
	ID            int32                            `json:"id"`
	Name          string                           `json:"name"`
	KeyItem       apiresource.NamedAPIResource     `json:"key_item"`
	Description   string                           `json:"description"`
	AlBhedLetter  string                           `json:"al_bhed_letter"`
	EnglishLetter string                           `json:"english_letter"`
	Areas         []apiresource.AreaAPIResource    `json:"areas"`
	Treasures     []apiresource.UnnamedAPIResource `json:"treasures"`
}
