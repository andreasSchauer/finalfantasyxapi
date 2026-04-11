package api

type CelestialWeapon struct {
	ID                      int32              		`json:"id"`
	Name					string					`json:"name"`
	Formula					string					`json:"formula"`
	Character				NamedAPIResource		`json:"character"`
	Aeon					*NamedAPIResource		`json:"aeon"`
	Equipment				NamedAPIResource		`json:"equipment"`
	AutoAbilities			[]NamedAPIResource		`json:"auto_abilities"`
	Crest					NamedAPIResource		`json:"crest"`
	Sigil					NamedAPIResource		`json:"sigil"`
	WpnTreasure				UnnamedAPIResource		`json:"wpn_treasure"`
	CrestTreasure			UnnamedAPIResource		`json:"crest_treasure"`
	SigilQuest				QuestAPIResource		`json:"sigil_quest"`
}