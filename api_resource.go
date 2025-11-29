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


func (cfg *apiConfig) newNamedAPIResource(endpoint string, id int32, name string, version *int32, spec *string) NamedAPIResource {
	return NamedAPIResource{
		Name: 			name,
		Version: 		version,
		Specification: 	spec,
		URL: 			cfg.createURL(endpoint, id),
	}
}

func createNamedAPIResources[T any](
	cfg *apiConfig,
	items []T,
	endpoint string,
	mapper func(T) (id int32, name string, version *int32, spec *string),
) []NamedAPIResource {
	var resources []NamedAPIResource

	for _, item := range items {
		id, name, version, spec := mapper(item)
		resource := cfg.newNamedAPIResource(endpoint, id, name, version, spec)

		resources = append(resources, resource)
	}

	return resources
}


func (cfg *apiConfig) newNamedAPIResourceSimple(endpoint string, id int32, name string) NamedAPIResource {
	return NamedAPIResource{
		Name: 			name,
		URL: 			cfg.createURL(endpoint, id),
	}
}

func createNamedAPIResourcesSimple[T any](
	cfg *apiConfig,
	items []T,
	endpoint string,
	mapper func(T) (id int32, name string),
) []NamedAPIResource {
	var resources []NamedAPIResource

	for _, item := range items {
		id, name := mapper(item)
		resource := cfg.newNamedAPIResourceSimple(endpoint, id, name)

		resources = append(resources, resource)
	}

	return resources
}