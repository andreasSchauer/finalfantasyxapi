package sectiontypes

type LocationSimple struct {
	ID        int32               `json:"id"`
	URL       string              `json:"url"`
	Name      string              `json:"name"`
	Shops     []ShopLocSimple     `json:"shops"`
	Treasures *TreasuresLocSimple `json:"treasures"`
	Monsters  []string            `json:"monsters"`
}

func (l LocationSimple) GetURL() string {
	return l.URL
}