package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Shop struct {
	ID           int32
	Version      *int32       `json:"version"`
	LocationArea LocationArea `json:"location_area"`
	AreaID       int32
	Notes        *string  `json:"notes"`
	Category     string   `json:"category"`
	Availability string   `json:"availability"`
	PreAirship   *SubShop `json:"pre_airship"`
	PostAirship  *SubShop `json:"post_airship"`
}

func (s Shop) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		h.DerefOrNil(s.Version),
		s.AreaID,
		h.DerefOrNil(s.Notes),
		s.Category,
		s.Availability,
	}
}

func (s Shop) ToKeyFields() []any {
	return []any{
		Key(s.LocationArea),
		h.DerefOrNil(s.Version),
	}
}

func (s Shop) GetID() int32 {
	return s.ID
}

func (s Shop) Error() string {
	return fmt.Sprintf("shop %s, %v", s.LocationArea, h.PtrToString(s.Version))
}

func (s Shop) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: s.ID,
	}
}

type SubShop struct {
	Items     []ShopItem      `json:"items"`
	Equipment []ShopEquipment `json:"equipment"`
	Type      database.ShopType
}

func (s SubShop) Error() string {
	return fmt.Sprintf("subshop type: %s", s.Type)
}
