package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type TestElemResist struct {
	element  int32
	affinity database.ElementalAffinity
}
