package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type ItemAmount struct {
	ID				*int32
	MasterItemID	int32
	ItemName		string		`json:"name"`
	Amount			int32		`json:"amount"`
}

func (ia ItemAmount) ToHashFields() []any {
	return []any{
		ia.MasterItemID,
		ia.Amount,
	}
}


func (ia ItemAmount) GetID() *int32 {	
	return ia.ID
}



func (l *lookup) seedItemAmount(qtx *database.Queries, itemAmount ItemAmount) (ItemAmount, error) {
	item, err := l.getItem(itemAmount.ItemName)
	if err != nil {
		return ItemAmount{}, err
	}
	itemAmount.MasterItemID = item.MasterItem.ID

	dbItemAmount, err := qtx.CreateItemAmount(context.Background(), database.CreateItemAmountParams{
		DataHash: generateDataHash(itemAmount),
		MasterItemID: itemAmount.MasterItemID,
		Amount: itemAmount.Amount,
	})
	if err != nil {
		return ItemAmount{}, fmt.Errorf("couldn't create Item Amount: %v", err)
	}

	itemAmount.ID = &dbItemAmount.ID

	return itemAmount, nil
}