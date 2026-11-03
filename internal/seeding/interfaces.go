package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type needsID interface {
	SetID(int32)
	seedtypes.Hashable
	fmt.Stringer
}