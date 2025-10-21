package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type ItemAmount struct {
	ID           int32
	MasterItemID int32
	ItemName     string `json:"name"`
	Amount       int32  `json:"amount"`
}

func (ia ItemAmount) ToHashFields() []any {
	return []any{
		ia.MasterItemID,
		ia.Amount,
	}
}

func (ia ItemAmount) GetID() int32 {
	return ia.ID
}

func (l *lookup) seedItemAmount(qtx *database.Queries, itemAmount ItemAmount) (ItemAmount, error) {
	var err error

	itemAmount.MasterItemID, err = assignFK(itemAmount.ItemName, l.getMasterItem)
	if err != nil {
		return ItemAmount{}, err
	}

	dbItemAmount, err := qtx.CreateItemAmount(context.Background(), database.CreateItemAmountParams{
		DataHash:     generateDataHash(itemAmount),
		MasterItemID: itemAmount.MasterItemID,
		Amount:       itemAmount.Amount,
	})
	if err != nil {
		return ItemAmount{}, fmt.Errorf("couldn't create Item Amount: %s - %d: %v", itemAmount.ItemName, itemAmount.Amount, err)
	}

	itemAmount.ID = dbItemAmount.ID

	return itemAmount, nil
}
