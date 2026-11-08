package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type KeyItem struct {
	ID              int32                 `json:"id"`
	Name            string                `json:"name"`
	UntypedItem     apiresource.Typed     `json:"untyped_item"`
	Category        string                `json:"category"`
	Description     string                `json:"description"`
	Effect          string                `json:"effect"`
	CelestialWeapon *apiresource.Named    `json:"celestial_weapon,omitempty"`
	Primer          *apiresource.Named    `json:"primer,omitempty"`
	Areas           []apiresource.Area    `json:"areas"`
	Treasures       []apiresource.Unnamed `json:"treasures"`
	Quests          []apiresource.Quest   `json:"quests"`
}
