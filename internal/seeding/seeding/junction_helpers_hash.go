package seeding

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func GenerateJunctionHash(j Junction, desc string) string {
	return h.CombineFields(j.ToHashFieldsJ(desc))
}