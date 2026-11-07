package sectiontypes

type AreaSimple struct {
	ID                int32               `json:"id"`
	URL               string              `json:"url"`
	Name              string              `json:"name"`
	ParentLocation    SimpleRef           `json:"parent_location"`
	ParentSublocation SimpleRef           `json:"parent_sublocation"`
	Version           *int32              `json:"version,omitempty"`
	Specification     *string             `json:"specification,omitempty"`
	HasSaveSphere     bool                `json:"has_save_sphere"`
	Availability      string              `json:"availability"`
	Shops             []ShopLocSimple     `json:"shops"`
	Treasures         *TreasuresLocSimple `json:"treasures"`
	Monsters          []string            `json:"monsters"`
}

func (a AreaSimple) GetURL() string {
	return a.URL
}
