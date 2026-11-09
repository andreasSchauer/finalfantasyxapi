package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getKeyItemRelationships(r *http.Request, keyItem seedtypes.KeyItem) (apitypes.KeyItem, error) {
	var rel apitypes.KeyItem
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := apilookups.E.KeyItems.QueryAvlRelGetParams(r, keyItem.ID)
	if err != nil {
		return apitypes.KeyItem{}, err
	}

	g.Go(func() error {
		var err error
		rel.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, keyItem, availabilityParams, getKeyItemSourceIDs(ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = apilookups.E.Quests.QueryAvlRelRun(ctx, keyItem, availabilityParams, getKeyItemSourceIDs(ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = apilookups.E.Areas.QueryAvlRelRun(ctx, keyItem, availabilityParams, getKeyItemAreaIDs())
		return err
	})

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = getKeyItemCelestialWeapon(ctx, keyItem)
		return err
	})

	var primer *apiresource.Named
	if keyItem.Category == string(database.KeyItemCategoryPrimer) {
		primerRes := handlerinput.NameToNamedAPIResource(apilookups.E.Primers, keyItem.Name, nil)
		primer = &primerRes
	}
	rel.Primer = primer

	err = g.Wait()
	if err != nil {
		return apitypes.KeyItem{}, err
	}

	return rel, nil
}

func getKeyItemCelestialWeapon(ctx context.Context, keyItem seedtypes.KeyItem) (*apiresource.Named, error) {
	if !(strings.HasSuffix(keyItem.Name, "crest") || strings.HasSuffix(keyItem.Name, "sigil")) {
		return nil, nil
	}

	keyItemBase, _, _ := strings.Cut(keyItem.Name, " ")

	celestialID, err := config.Cfg.DB.GetKeyItemCelestialWeapon(ctx, database.KeyItemBase(keyItemBase))
	if err != nil {
		return nil, apilookups.E.CelestialWeapons.ErrorDBOne(keyItem, err)
	}

	celestialRes := handlerinput.IDToNamedAPIResource(apilookups.E.CelestialWeapons, celestialID)
	return &celestialRes, nil
}
