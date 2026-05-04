package seeding

import "fmt"

type ItemAmount struct {
	ID int32 `json:"-"`
	MasterItem
	ItemName string `json:"name"`
	Amount   int32  `json:"amount"`
}

func (ia ItemAmount) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", ia),
		ia.MasterItem.ID,
		ia.Amount,
	}
}

func (ia ItemAmount) ToKeyFields() []any {
	return []any{
		ia.ItemName,
		ia.Amount,
	}
}

func (ia ItemAmount) GetID() int32 {
	return ia.ID
}

func (ia *ItemAmount) SetID(id int32) {
	ia.ID = id
}

func (ia ItemAmount) GetName() string {
	return ia.ItemName
}

func (ia ItemAmount) GetVersion() *int32 {
	return nil
}

func (ia ItemAmount) GetVal() int32 {
	return ia.Amount
}

func (ia ItemAmount) Error() string {
	return fmt.Sprintf("item amount with item: %s, amount: %d", ia.ItemName, ia.Amount)
}
