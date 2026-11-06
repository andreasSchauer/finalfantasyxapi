package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type KeyItem struct {
	ID              int32                            `json:"id"`
	Name            string                           `json:"name"`
	UntypedItem     apiresource.TypedAPIResource     `json:"untyped_item"`
	Category        string                           `json:"category"`
	Description     string                           `json:"description"`
	Effect          string                           `json:"effect"`
	CelestialWeapon *apiresource.NamedAPIResource    `json:"celestial_weapon,omitempty"`
	Primer          *apiresource.NamedAPIResource    `json:"primer,omitempty"`
	Areas           []apiresource.AreaAPIResource    `json:"areas"`
	Treasures       []apiresource.UnnamedAPIResource `json:"treasures"`
	Quests          []apiresource.QuestAPIResource   `json:"quests"`
}
