package api

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type Folder struct {
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Contains    []Folder `json:"contains,omitempty"`
}

func (f Folder) GetURL() string {
	return f.URL
}

func getBody(cfg *Config, endpoint typestrings.EndpointName) Folder {
	bodyURL := getSectionURL(cfg, endpoint, typestrings.SnBody)
	return Folder{
		URL:         bodyURL,
		Description: fmt.Sprintf("holds resources that explain how the json request body for the '/%s' endpoint is formed.", endpoint),
		Contains: []Folder{
			{
				URL:         appendToURL(bodyURL, string(typestrings.SnDocs)),
				Description: fmt.Sprintf("contains explanations, type definitions, value limits, default values, requirements, dependencies, and rules for every field that can be used to form the json request body for the '%s' endpoint.", endpoint),
			},
			{
				URL:         appendToURL(bodyURL, string(typestrings.SnSamples)),
				Description: fmt.Sprintf("contains a list of example json bodies along with explanations that can be used as a starting point to learn how to operate the '/%s' endpoint.", endpoint),
			},
		},
	}
}
