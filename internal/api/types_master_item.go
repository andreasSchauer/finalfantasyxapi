package api

type MasterItem struct {
	ID                 int32                `json:"id"`
	Name               string               `json:"name"`
	Type               NamedAPIResource    	`json:"type"`
	TypedItem		   NamedAPIResource		`json:"typed_item"`
	Description        string               `json:"description"`
	Effect             string               `json:"effect"`
	Monsters           []NamedAPIResource   `json:"monsters"`
	Treasures          []UnnamedAPIResource `json:"treasures"`
	Shops              []UnnamedAPIResource `json:"shops"`
	Quests             []QuestAPIResource	`json:"quests"`
}