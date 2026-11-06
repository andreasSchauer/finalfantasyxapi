package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type CelestialWeapon struct {
	ID            int32                          `json:"id"`
	Name          string                         `json:"name"`
	Formula       string                         `json:"formula"`
	Character     apiresource.NamedAPIResource   `json:"character"`
	Aeon          *apiresource.NamedAPIResource  `json:"aeon"`
	Equipment     apiresource.NamedAPIResource   `json:"equipment"`
	AutoAbilities []apiresource.NamedAPIResource `json:"auto_abilities"`
	Crest         apiresource.NamedAPIResource   `json:"crest"`
	Sigil         apiresource.NamedAPIResource   `json:"sigil"`
	WpnTreasure   apiresource.UnnamedAPIResource `json:"wpn_treasure"`
	CrestTreasure apiresource.UnnamedAPIResource `json:"crest_treasure"`
	SigilQuest    apiresource.QuestAPIResource   `json:"sigil_quest"`
}
