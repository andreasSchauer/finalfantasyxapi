package sectiontypes

type SubquestSimple struct {
	ID        int32    `json:"id"`
	URL       string   `json:"url"`
	Name      string   `json:"name"`
	Condition *string  `json:"condition"`
	Areas     []string `json:"areas"`
	Reward    string   `json:"reward"`
}

func (s SubquestSimple) GetURL() string {
	return s.URL
}
