package api

import "fmt"

type Folder struct {
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Contains    []Folder `json:"contains,omitempty"`
}

func (f Folder) GetURL() string {
	return f.URL
}

func getBody(cfg *Config, endpoint EndpointName) Folder {
	bodyURL := getSectionURL(cfg, endpoint, snBody)
	return Folder{
		URL:         bodyURL,
		Description: fmt.Sprintf("holds resources that explain how the json request body for the '/%s' endpoint is formed.", endpoint),
		Contains: []Folder{
			{
				URL:         appendToURL(bodyURL, string(snDocs)),
				Description: fmt.Sprintf("contains explanations, type definitions, value limits, default values, requirements, dependencies, and rules for every field that can be used to form the json request body for the '%s' endpoint.", endpoint),
			},
			{
				URL:         appendToURL(bodyURL, string(snSamples)),
				Description: fmt.Sprintf("contains a list of example json bodies along with explanations that can be used as a starting point to learn how to operate the '/%s' endpoint.", endpoint),
			},
		},
	}
}
