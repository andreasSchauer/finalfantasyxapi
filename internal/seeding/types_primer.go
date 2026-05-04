package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Primer struct {
	ID            int32
	Name          string `json:"name"`
	AlBhedLetter  string `json:"al_bhed_letter"`
	EnglishLetter string `json:"english_letter"`
	KeyItemID     int32
}

func (p Primer) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", p),
		p.AlBhedLetter,
		p.EnglishLetter,
		p.KeyItemID,
	}
}

func (p Primer) GetID() int32 {
	return p.ID
}

func (p Primer) Error() string {
	return fmt.Sprintf("primer %s", p.Name)
}

func (p Primer) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   p.ID,
		Name: p.Name,
	}
}
