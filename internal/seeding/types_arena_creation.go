package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ArenaCreation struct {
	ID                        int32
	SubquestID                int32
	MonsterID                 *int32
	Name                      string  `json:"name"`
	Category                  string  `json:"category"`
	RequiredArea              *string `json:"required_area"`
	RequiredSpecies           *string `json:"required_species"`
	UnderwaterOnly            bool    `json:"underwater_only"`
	CreationsUnlockedCategory *string `json:"creations_unlocked_category"`
	Amount                    int32   `json:"amount"`
}

func (a ArenaCreation) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.SubquestID,
		h.DerefOrNil(a.MonsterID),
		a.Category,
		h.DerefOrNil(a.RequiredArea),
		h.DerefOrNil(a.RequiredSpecies),
		a.UnderwaterOnly,
		h.DerefOrNil(a.CreationsUnlockedCategory),
		a.Amount,
	}
}

func (a ArenaCreation) ToKeyFields() []any {
	return []any{
		a.Name,
	}
}

func (a ArenaCreation) GetID() int32 {
	return a.ID
}

func (a ArenaCreation) Error() string {
	return fmt.Sprintf("monster arena creation %s", a.Name)
}

func (a ArenaCreation) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   a.ID,
		Name: a.Name,
	}
}
