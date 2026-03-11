package api

type RonsoRage struct {
	ID              int32					`json:"id"`
	Name            string					`json:"name"`
	Description    	string					`json:"description"`
	Effect          string					`json:"effect"`
	Overdrive		NamedAPIResource		`json:"overdrive"`
	Monsters		[]NamedAPIResource		`json:"monsters"`
}