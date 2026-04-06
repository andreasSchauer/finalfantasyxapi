package api

type MasterItem struct {
	ID                 int32                `json:"id"`
	Name               string               `json:"name"`
	Type               NamedAPIResource    	`json:"type"`
	TypedItem		   NamedAPIResource		`json:"typed_item"`
	Description        string               `json:"description"`
	Effect             string               `json:"effect"`
	ObtainableFrom	   ObtainableFrom		`json:"obtainable_from"`
}

type ObtainableFrom struct {
	Monsters 	bool	`json:"monsters"`
	Treasures 	bool	`json:"treasures"`
	Shops		bool	`json:"shops"`
	Quests		bool	`json:"quests"`
}