package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getKeyItemRelationships(server *Server, r *http.Request, keyItem seedtypes.KeyItem) (KeyItem, error) {
	var rel KeyItem
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.KeyItems.queryAvlRelGetParams(server, r, keyItem.ID)
	if err != nil {
		return KeyItem{}, err
	}

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.queryAvlRelRun(server, ctx, keyItem, availabilityParams, getKeyItemSourceIDs(server, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = E.Quests.queryAvlRelRun(server, ctx, keyItem, availabilityParams, getKeyItemSourceIDs(server, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = E.Areas.queryAvlRelRun(server, ctx, keyItem, availabilityParams, getKeyItemAreaIDs(server))
		return err
	})

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = getKeyItemCelestialWeapon(server, ctx, keyItem)
		return err
	})

	var primer *NamedAPIResource
	if keyItem.Category == string(database.KeyItemCategoryPrimer) {
		primerRes := nameToNamedAPIResource(server, E.Primers, keyItem.Name, nil)
		primer = &primerRes
	}
	rel.Primer = primer

	err = g.Wait()
	if err != nil {
		return KeyItem{}, err
	}

	return rel, nil
}

func getKeyItemCelestialWeapon(server *Server, ctx context.Context, keyItem seedtypes.KeyItem) (*NamedAPIResource, error) {
	if !(strings.HasSuffix(keyItem.Name, "crest") || strings.HasSuffix(keyItem.Name, "sigil")) {
		return nil, nil
	}

	keyItemBase := strings.Split(keyItem.Name, " ")[0]

	celestialID, err := server.Cfg.DB.GetKeyItemCelestialWeapon(ctx, database.KeyItemBase(keyItemBase))
	if err != nil {
		return nil, newHTTPErrorDBOne(E.CelestialWeapons.resTypeSingle, keyItem, err)
	}

	celestialRes := idToNamedAPIResource(server, E.CelestialWeapons, celestialID)
	return &celestialRes, nil
}
