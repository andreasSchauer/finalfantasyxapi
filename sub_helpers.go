package main

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SubRef struct {
	ID            int32   `json:"id,omitempty"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
}

func createSubReference(id int32, name string, version *int32, spec *string) SubRef {
	return SubRef{
		ID:   			id,
		Name: 			name,
		Version: 		version,
		Specification: 	spec,
	}
}

func sortNamesByID[T h.HasID](s []string, lookup map[string]T) []string {
	slices.SortStableFunc(s, func(a, b string) int {
		A, _ := seeding.GetResource(a, lookup)
		B, _ := seeding.GetResource(b, lookup)

		if A.GetID() < B.GetID() {
			return -1
		}

		if A.GetID() > B.GetID() {
			return 1
		}

		return 0
	})

	return s
}
