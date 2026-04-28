package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ItemAmount struct {
	ID           int32	`json:"-"`
	MasterItem
	ItemName     string `json:"name"`
	Amount       int32  `json:"amount"`
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

	masterItem, err := GetResource(itemAmount.ItemName, l.MasterItems)
	if err != nil {
		return ItemAmount{}, h.NewErr(itemAmount.Error(), err)
	}

	itemAmount.MasterItem = masterItem

	dbItemAmount, err := qtx.CreateItemAmount(context.Background(), database.CreateItemAmountParams{
		DataHash:     generateDataHash(itemAmount),
		MasterItemID: itemAmount.MasterItem.ID,
		Amount:       itemAmount.Amount,
	})
	if err != nil {
		return ItemAmount{}, h.NewErr(itemAmount.Error(), err, "couldn't create item amount")
	}

	itemAmount.ID = dbItemAmount.ID

	return itemAmount, nil
}


func (l *Lookup) loop2SeedItemAmounts(qtx *database.Queries, ctx context.Context) error {
	itemAmts, err := l.extractItemAmounts()
	if err != nil {
		return err
	}

	params := database.CreateItemAmountBulkParams{
		DataHash:     make([]string, len(itemAmts)),
		MasterItemID: make([]int32, len(itemAmts)),
		Amount:       make([]int32, len(itemAmts)),
	}

	for i, ia := range itemAmts {
		params.DataHash[i] = generateDataHash(ia)
		params.MasterItemID[i] = ia.MasterItem.ID
		params.Amount[i] = ia.Amount
	}

	dbRows, err := qtx.CreateItemAmountBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create item amounts: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractItemAmounts() ([]ItemAmount, error) {
	itemAmounts := []ItemAmount{}
	var err error

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		if autoAbility.RequiredItem != nil {
			itemAmt, err := l.prepareItemAmt(autoAbility.RequiredItem)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	for i := range l.json.blitzballPositions {
		position := &l.json.blitzballPositions[i]

		for j := range position.Items {
			itemAmt, err := l.prepareItemAmt(&position.Items[j].ItemAmount)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	for i := range l.json.monsters {
		monster := l.json.monsters[i]

		if monster.Items != nil {
			items := monster.Items
			sc := items.StealCommon
			sr := items.StealRare
			dc := items.DropCommon
			dr := items.DropRare
			sdc := items.SecondaryDropCommon
			sdr := items.SecondaryDropRare
			br := items.Bribe

			if sc != nil {
				sc, err = l.prepareItemAmt(sc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sc)
			}

			if sr != nil {
				sr, err = l.prepareItemAmt(sr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sr)
			}

			if dc != nil {
				dc, err = l.prepareItemAmt(dc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *dc)
			}

			if dr != nil {
				dr, err = l.prepareItemAmt(dr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *dr)
			}

			if sdc != nil {
				sdc, err = l.prepareItemAmt(sdc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sdc)
			}

			if sdr != nil {
				sdr, err = l.prepareItemAmt(sdr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sdr)
			}

			if br != nil {
				br, err = l.prepareItemAmt(br)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *br)
			}

			for j := range items.OtherItems {
				itemAmt, err := l.prepareItemAmt(&items.OtherItems[j].ItemAmount)
				if err != nil {
					return nil, err
				}

				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	for i := range l.json.playerAbilities {
		playerAbility := &l.json.playerAbilities[i]

		if playerAbility.AeonLearnItem != nil {
			itemAmt, err := l.prepareItemAmt(playerAbility.AeonLearnItem)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		if sidequest.Completion != nil {
			itemAmt, err := l.prepareItemAmt(&sidequest.Completion.Reward)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}

		for j := range sidequest.Subquests {
			subquest := sidequest.Subquests[j]

			if subquest.Completion != nil {
				itemAmt, err := l.prepareItemAmt(&subquest.Completion.Reward)
				if err != nil {
					return nil, err
				}

				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	for i := range l.json.treasureLists {
		treasureList := &l.json.treasureLists[i]
		for j := range treasureList.Treasures {
			treasure := &treasureList.Treasures[j]

			for j := range treasure.Items {
				itemAmt, err := l.prepareItemAmt(&treasure.Items[j])
				if err != nil {
					return nil, err
				}

				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	return dedupeRows(itemAmounts, l.Hashes), nil
}

func (l *Lookup) prepareItemAmt(ia *ItemAmount) (*ItemAmount, error) {
	var err error
	ia.MasterItem.Name = ia.ItemName

	ia.MasterItem.ID, err = assignFK(ia.MasterItem.Name, l.MasterItems)
	if err != nil {
		return nil, err
	}

	mi, err := GetResourceByID(ia.MasterItem.ID, l.MasterItemsID)
	if err != nil {
		return nil, err
	}
	ia.MasterItem.Type = mi.Type

	return ia, nil
}