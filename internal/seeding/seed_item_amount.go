package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ItemAmount struct {
	ID           int32	`json:"-"`
	MasterItemID int32	`json:"-"`
	ItemName     string `json:"name"`
	Amount       int32  `json:"amount"`
}

func (ia ItemAmount) ToHashFields() []any {
	return []any{
		ia.MasterItemID,
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

func (l *Lookup) seedItemAmount(qtx *database.Queries, itemAmount ItemAmount) (ItemAmount, error) {
	var err error

	itemAmount.MasterItemID, err = assignFK(itemAmount.ItemName, l.MasterItems)
	if err != nil {
		return ItemAmount{}, h.NewErr(itemAmount.Error(), err)
	}

	dbItemAmount, err := qtx.CreateItemAmount(context.Background(), database.CreateItemAmountParams{
		DataHash:     generateDataHash(itemAmount),
		MasterItemID: itemAmount.MasterItemID,
		Amount:       itemAmount.Amount,
	})
	if err != nil {
		return ItemAmount{}, h.NewErr(itemAmount.Error(), err, "couldn't create item amount")
	}

	itemAmount.ID = dbItemAmount.ID

	return itemAmount, nil
}
