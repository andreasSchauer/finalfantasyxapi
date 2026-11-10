package sectiontypes

type SublocationSimple struct {
	ID             int32               `json:"id"`
	URL            string              `json:"url"`
	Name           string              `json:"name"`
	ParentLocation SimpleRef           `json:"parent_location"`
	Shops          []ShopLocSimple     `json:"shops"`
	Treasures      *TreasuresLocSimple `json:"treasures"`
	Monsters       []string            `json:"monsters"`
}

func (s SublocationSimple) GetURL() string {
	return s.URL
}
