package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type CelestialWeapon struct {
	ID          int32
	Name        string `json:"name"`
	Character   string `json:"character"`
	CharacterID *int32
	KeyItemBase string  `json:"key_item_base"`
	Aeon        *string `json:"aeon"`
	AeonID      *int32
	Formula     string `json:"formula"`
}

func (cw CelestialWeapon) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", cw),
		cw.Name,
		cw.KeyItemBase,
		cw.Formula,
	}
}

func (cw CelestialWeapon) ToKeyFields() []any {
	return []any{
		cw.Name,
	}
}

func (cw CelestialWeapon) GetID() int32 {
	return cw.ID
}

func (cw CelestialWeapon) Error() string {
	return fmt.Sprintf("celestial weapon %s", cw.Name)
}

func (cw CelestialWeapon) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   cw.ID,
		Name: cw.Name,
	}
}
