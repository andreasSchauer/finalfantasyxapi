// Package sectiontypes ajsidasdo
package sectiontypes

type SimpleRef struct {
	ID            int32   `json:"id,omitempty"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
}
