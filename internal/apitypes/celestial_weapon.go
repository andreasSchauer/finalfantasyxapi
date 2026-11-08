package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type CelestialWeapon struct {
	ID            int32               `json:"id"`
	Name          string              `json:"name"`
	Formula       string              `json:"formula"`
	Character     apiresource.Named   `json:"character"`
	Aeon          *apiresource.Named  `json:"aeon"`
	Equipment     apiresource.Named   `json:"equipment"`
	AutoAbilities []apiresource.Named `json:"auto_abilities"`
	Crest         apiresource.Named   `json:"crest"`
	Sigil         apiresource.Named   `json:"sigil"`
	WpnTreasure   apiresource.Unnamed `json:"wpn_treasure"`
	CrestTreasure apiresource.Unnamed `json:"crest_treasure"`
	SigilQuest    apiresource.Quest   `json:"sigil_quest"`
}
