package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Item struct {
	ID int32
	MasterItem
	ItemAbility
	Description           string   `json:"description"`
	Effect                string   `json:"effect"`
	RelatedStats          []string `json:"related_stats"`
	SphereGridDescription *string  `json:"sphere_grid_description"`
	Category              string   `json:"category"`
	Usability             string   `json:"usability"`
	AvailableMenus        []string `json:"available_menus"`
	BasePrice             *int32   `json:"base_price"`
	SellValue             int32    `json:"sell_value"`
}

func (i Item) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", i),
		i.MasterItem.ID,
		i.Description,
		i.Effect,
		h.DerefOrNil(i.SphereGridDescription),
		i.Category,
		i.Usability,
		h.DerefOrNil(i.BasePrice),
		i.SellValue,
	}
}

func (i Item) GetID() int32 {
	return i.ID
}

func (i Item) Error() string {
	return fmt.Sprintf("item %s", i.Name)
}

func (i Item) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   i.ID,
		Name: i.Name,
	}
}
