package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop2SeedItemAmounts(qtx *database.Queries, ctx context.Context) error {
	itemAmts, err := extractItemAmounts()
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
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractItemAmounts() ([]seedtypes.ItemAmount, error) {
	itemAmounts := []seedtypes.ItemAmount{}
	var err error

	for i := range rawJSON.autoAbilities {
		autoAbility := &rawJSON.autoAbilities[i]

		if autoAbility.RequiredItem != nil {
			itemAmt, err := prepareItemAmt(autoAbility.RequiredItem)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	for i := range rawJSON.blitzballPositions {
		position := &rawJSON.blitzballPositions[i]

		for j := range position.Items {
			itemAmt, err := prepareItemAmt(&position.Items[j].ItemAmount)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	monsterItems, err := extractItemAmtsMonsters()
	if err != nil {
		return nil, err
	}
	itemAmounts = append(itemAmounts, monsterItems...)

	for i := range rawJSON.playerAbilities {
		playerAbility := &rawJSON.playerAbilities[i]

		if playerAbility.AeonLearnItem != nil {
			itemAmt, err := prepareItemAmt(playerAbility.AeonLearnItem)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	for i := range rawJSON.sidequests {
		sidequest := &rawJSON.sidequests[i]

		if sidequest.Completion != nil {
			itemAmt, err := prepareItemAmt(&sidequest.Completion.Reward)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}

		for j := range sidequest.Subquests {
			subquest := &sidequest.Subquests[j]

			if subquest.Completion != nil {
				itemAmt, err := prepareItemAmt(&subquest.Completion.Reward)
				if err != nil {
					return nil, err
				}

				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	for i := range rawJSON.treasureLists {
		treasureList := &rawJSON.treasureLists[i]
		for j := range treasureList.Treasures {
			treasure := &treasureList.Treasures[j]

			for j := range treasure.Items {
				itemAmt, err := prepareItemAmt(&treasure.Items[j])
				if err != nil {
					return nil, err
				}

				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	return dedupeRows(itemAmounts, hashes), nil
}

func extractItemAmtsMonsters() ([]seedtypes.ItemAmount, error) {
	itemAmounts := []seedtypes.ItemAmount{}
	var err error

	for i := range rawJSON.monsters {
		monster := rawJSON.monsters[i]

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
				sc, err = prepareItemAmt(sc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sc)
			}

			if sr != nil {
				sr, err = prepareItemAmt(sr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sr)
			}

			if dc != nil {
				dc, err = prepareItemAmt(dc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *dc)
			}

			if dr != nil {
				dr, err = prepareItemAmt(dr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *dr)
			}

			if sdc != nil {
				sdc, err = prepareItemAmt(sdc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sdc)
			}

			if sdr != nil {
				sdr, err = prepareItemAmt(sdr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sdr)
			}

			if br != nil {
				br, err = prepareItemAmt(br)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *br)
			}

			for j := range items.OtherItems {
				itemAmt, err := prepareItemAmt(&items.OtherItems[j].ItemAmount)
				if err != nil {
					return nil, err
				}

				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	return itemAmounts, nil
}

func prepareItemAmt(ia *seedtypes.ItemAmount) (*seedtypes.ItemAmount, error) {
	var err error
	ia.MasterItem.Name = ia.ItemName

	ia.MasterItem.ID, err = assignFK(ia.MasterItem.Name, lookup.L.MasterItems)
	if err != nil {
		return nil, err
	}

	mi, err := lookup.GetResourceByID(ia.MasterItem.ID, lookup.L.MasterItemsID)
	if err != nil {
		return nil, err
	}
	ia.MasterItem.Type = mi.Type

	return ia, nil
}
