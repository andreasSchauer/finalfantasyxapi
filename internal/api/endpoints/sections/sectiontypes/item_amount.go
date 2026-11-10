package sectiontypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

type ItemAmountSimple struct {
	ia            seedtypes.ItemAmount `json:"-"`
	ItemAmountStr string               `json:"item"`
	ItemType      database.ItemType    `json:"item_type"`
}
