package main


type NamedApiResourceList struct {
	Count		int					`json:"count"`
	Next		*string				`json:"next"`
	Previous	*string				`json:"previous"`
	Results		[]NamedAPIResource	`json:"results"`
}


type NamedAPIResource struct {
	Name			string		`json:"name"`
	Version			*int32		`json:"version,omitempty"`
	Specification	*string		`json:"specification,omitempty"`
	URL				string		`json:"url"`
}