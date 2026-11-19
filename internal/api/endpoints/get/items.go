package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"golang.org/x/sync/errgroup"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Item(r *http.Request, i handlerinput.Main[seedtypes.Item, apitypes.Item, apiresource.Named], id int32) (apitypes.Item, error) {
	item, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Item{}, err
	}
	
	response := apitypes.Item{
		ID:                 item.ID,
		Name:               item.Name,
		UntypedItem:        handlerinput.IDToTypedAPIResource(apilookups.E.AllItems, item.MasterItem.ID),
		Category:           item.Category,
		Description:        item.Description,
		SgDescription:      item.SphereGridDescription,
		Effect:             item.Effect,
		Usability:          item.Usability,
		BasePrice:          item.BasePrice,
		SellValue:          item.SellValue,
		AvailableMenus:     handlerinput.NamesToNamedAPIResources(apilookups.E.Submenus, item.AvailableMenus),
		RelatedStats:       handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, item.RelatedStats),
	}

	err = getItemRelationships(r, item, &response)
	if err != nil {
		return apitypes.Item{}, err
	}

	return response, nil
}

func getItemRelationships(r *http.Request, item seedtypes.Item, response *apitypes.Item) error {
	availabilityParams, err := apilookups.E.Items.QueryAvlRelGetParams(r, item.ID)
	if err != nil {
		return err
	}

	err = runItemRelQueries(r, item, response, availabilityParams)
	if err != nil {
		return err
	}

	if item.Category == string(database.ItemCategorySphere) {
		sphereRes := handlerinput.NameToNamedAPIResource(apilookups.E.Spheres, item.Name, nil)
		response.Sphere = &sphereRes
	}

	var itemAbilityNamePtr *string
	if len(item.BattleInteractions) > 0 {
		itemAbilityNamePtr = &item.Name
	}

	response.ItemAbility = handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.ItemAbilities, itemAbilityNamePtr, nil)

	return nil
}

func runItemRelQueries(r *http.Request, item seedtypes.Item, response *apitypes.Item, availabilityParams handlerinput.RelAvlParams) error {
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		monsters, err := apilookups.E.Monsters.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ts.VstMonster))
		if err != nil {
			return err
		}
		response.Monsters = getMonItemAmts(monsters, item.Name)
		return nil
	})

	g.Go(func() error {
		treasures, err := apilookups.E.Treasures.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ts.VstTreasure))
		if err != nil {
			return err
		}
		response.Treasures = handlerinput.ItemAmtsToChildResAmts2(apilookups.E.Treasures, treasures, item.ID)
		return nil
	})

	g.Go(func() error {
		var err error
		response.Shops, err = apilookups.E.Shops.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ts.VstShop))
		return err
	})

	g.Go(func() error {
		quests, err := apilookups.E.Quests.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ts.VstQuest))
		if err != nil {
			return err
		}
		response.Quests = handlerinput.ItemAmtsToChildResAmts(apilookups.E.Quests, quests)
		return nil
	})

	g.Go(func() error {
		blitzballPrizes, err := apilookups.E.BlitzballPrizes.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ts.VstBlitzball))
		if err != nil {
			return err
		}
		response.BlitzballPrizes = handlerinput.ItemAmtsToChildResAmts2(apilookups.E.BlitzballPrizes, blitzballPrizes, item.ID)
		return nil
	})

	g.Go(func() error {
		playerAbilities, err := apilookups.E.PlayerAbilities.DBGetRelations(ctx, item, config.Cfg.DB.GetItemPlayerAbilityIDs)
		if err != nil {
			return err
		}
		response.AeonLearnAbilities = handlerinput.ItemAmtsToChildResAmts(apilookups.E.PlayerAbilities, playerAbilities)
		return nil
	})

	g.Go(func() error {
		autoAbilities, err := apilookups.E.AutoAbilities.DBGetRelations(ctx, item, config.Cfg.DB.GetItemAutoAbilityIDs)
		if err != nil {
			return err
		}
		response.AutoAbilities = handlerinput.ItemAmtsToChildResAmts(apilookups.E.AutoAbilities, autoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		response.Mixes, err = apilookups.E.Mixes.DBGetRelations(ctx, item, config.Cfg.DB.GetItemMixIDs)
		return err
	})

	return g.Wait()
}

func getMonItemAmts(monsters []apiresource.Named, itemName string) []apitypes.MonItemAmts {
	monItemAmts := []apitypes.MonItemAmts{}

	for _, monster := range monsters {
		monItemAmt := createItemMonster(itemName, monster)
		monItemAmts = append(monItemAmts, monItemAmt)
	}

	return monItemAmts
}

func createItemMonster(itemName string, mon apiresource.Named) apitypes.MonItemAmts {
	monster, _ := lookup.GetResourceByID(mon.ID, lookup.L.MonstersID)

	return apitypes.MonItemAmts{
		Monster:             mon,
		StealCommon:         getAmountMonItem(itemName, monster.Items.StealCommon),
		StealRare:           getAmountMonItem(itemName, monster.Items.StealRare),
		DropCommon:          getAmountMonItem(itemName, monster.Items.DropCommon),
		DropRare:            getAmountMonItem(itemName, monster.Items.DropRare),
		SecondaryDropCommon: getAmountMonItem(itemName, monster.Items.SecondaryDropCommon),
		SecondaryDropRare:   getAmountMonItem(itemName, monster.Items.SecondaryDropRare),
		Bribe:               getAmountMonItem(itemName, monster.Items.Bribe),
		Other:               getAmountOther(itemName, monster.Items.OtherItems),
	}
}

func getAmountMonItem(wantedItem string, ia *seedtypes.ItemAmount) int32 {
	if ia == nil || ia.ItemName != wantedItem {
		return 0
	}

	return ia.Amount
}

func getAmountOther(wantedItem string, otherItems []seedtypes.PossibleItem) int32 {
	for _, otherItem := range otherItems {
		if otherItem.ItemAmount.ItemName == wantedItem {
			return otherItem.ItemAmount.Amount
		}
	}

	return 0
}
