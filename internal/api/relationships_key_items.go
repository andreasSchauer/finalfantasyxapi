package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getKeyItemRelationships(cfg *Config, r *http.Request, keyItem seeding.KeyItem) (KeyItem, error) {
	var rel KeyItem
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.keyItems, keyItem.ID)
	if err != nil {
		return KeyItem{}, err
	}

	g.Go(func() error {
		var err error
		rel.Treasures, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.treasures, keyItem, availabilityParams, getKeyItemSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.quests, keyItem, availabilityParams, getKeyItemSourceIDs(cfg, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.areas, keyItem, availabilityParams, getKeyItemAreaIDs(cfg))
		return err
	})

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = getKeyItemCelestialWeapon(cfg, ctx, keyItem)
		return err
	})

	var primer *NamedAPIResource
	if keyItem.Category == string(database.KeyItemCategoryPrimer) {
		primerRes := nameToNamedAPIResource(cfg, cfg.e.primers, keyItem.Name, nil)
		primer = &primerRes
	}
	rel.Primer = primer

	err = g.Wait()
	if err != nil {
		return KeyItem{}, err
	}

	return rel, nil
}

func getKeyItemCelestialWeapon(cfg *Config, ctx context.Context, keyItem seeding.KeyItem) (*NamedAPIResource, error) {
	if !(strings.HasSuffix(keyItem.Name, "crest") || strings.HasSuffix(keyItem.Name, "sigil")) {
		return nil, nil
	}

	keyItemBase := strings.Split(keyItem.Name, " ")[0]

	celestialID, err := cfg.db.GetKeyItemCelestialWeapon(ctx, database.KeyItemBase(keyItemBase))
	if err != nil {
		return nil, newHTTPErrorDbOne(cfg.e.celestialWeapons.resTypeSingle, keyItem, err)
	}

	celestialRes := idToNamedAPIResource(cfg, cfg.e.celestialWeapons, celestialID)
	return &celestialRes, nil
}
