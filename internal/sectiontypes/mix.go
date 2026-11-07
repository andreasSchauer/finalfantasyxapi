package sectiontypes

type MixSimple struct {
	ID           int32    `json:"id"`
	Name         string   `json:"name"`
	URL          string   `json:"url"`
	Category     string   `json:"category"`
	Combinations []string `json:"combinations"`
}

func (m MixSimple) GetURL() string {
	return m.URL
}
