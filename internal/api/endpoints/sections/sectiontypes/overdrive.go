package sectiontypes

type OverdriveSimple struct {
	ID                 int32    `json:"id"`
	URL                string   `json:"url"`
	Name               string   `json:"name"`
	Version            *int32   `json:"version,omitempty"`
	Specification      *string  `json:"specification,omitempty"`
	Rank               *int32   `json:"rank"`
	OverdriveAbilities []string `json:"overdrive_abilities"`
}

func (o OverdriveSimple) GetURL() string {
	return o.URL
}
