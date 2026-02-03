package main

import (
	"fmt"
	"slices"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SubRef struct {
	ID            int32   `json:"id,omitempty"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
}

func createSubReference(id int32, name string) SubRef {
	return SubRef{
		ID:   id,
		Name: name,
	}
}

func nameVersionToString(name string, version *int32, spec *string) string {
	var verStr string
	var specStr string

	if version != nil {
		intVer := int(*version)
		verStr = fmt.Sprintf(" %s", strconv.Itoa(intVer))
	}

	if spec != nil {
		specStr = fmt.Sprintf(" (%s)", *spec)
	}

	return name + verStr + specStr
}

func nameAmountString(name string, amount int32) string {
	return fmt.Sprintf("%s x%d", name, amount)
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