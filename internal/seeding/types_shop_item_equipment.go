package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type ShopItem struct {
	ID     int32
	ItemID int32
	Name   string `json:"name"`
	Price  int32  `json:"price"`
}

func (s ShopItem) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.ItemID,
		s.Price,
	}
}

func (s ShopItem) ToKeyFields() []any {
	return []any{
		s.Name,
		s.Price,
	}
}

func (s ShopItem) GetID() int32 {
	return s.ID
}

func (s *ShopItem) SetID(id int32) {
	s.ID = id
}

func (s ShopItem) Error() string {
	return fmt.Sprintf("shop item %s, price %d", s.Name, s.Price)
}

type ShopEquipment struct {
	ID       int32
	ShopID   int32
	ShopType database.ShopType
	TreasureEquipment
	Price int32 `json:"price"`
}

func (s ShopEquipment) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.ShopID,
		s.EquipmentNameID,
		s.ShopType,
		s.EmptySlotsAmount,
		s.Price,
	}
}

func (s ShopEquipment) ToKeyFields() []any {
	return []any{
		s.Name,
		s.EmptySlotsAmount,
		s.Price,
	}
}

func (s ShopEquipment) GetID() int32 {
	return s.ID
}

func (s *ShopEquipment) SetID(id int32) {
	s.ID = id
}

func (s ShopEquipment) Error() string {
	return fmt.Sprintf("shop equipment %s, empty slots %d, price %d", s.Name, s.EmptySlotsAmount, s.Price)
}
