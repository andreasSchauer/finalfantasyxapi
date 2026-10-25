package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getCelestialWeaponRelationships(cfg *Config, r *http.Request, cw seeding.CelestialWeapon) (CelestialWeapon, error) {
	crest, _ := seeding.GetResource(cw.KeyItemBase+" crest", cfg.L.KeyItems)
	sigil, _ := seeding.GetResource(cw.KeyItemBase+" sigil", cfg.L.KeyItems)
	equipment, _ := seeding.GetResource(cw, cfg.L.EquipmentNames)

	rel := CelestialWeapon{
		Equipment: nameToNamedAPIResource(cfg, cfg.E.Equipment, equipment.Name, nil),
		Crest:     nameToNamedAPIResource(cfg, cfg.E.KeyItems, crest.Name, nil),
		Sigil:     nameToNamedAPIResource(cfg, cfg.E.KeyItems, sigil.Name, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		tables, err := getResourcesDbItem(cfg, ctx, cfg.E.EquipmentTables, equipment, cfg.db.GetEquipmentEquipmentTableIDs)
		if err != nil {
			return err
		}
		table, _ := seeding.GetResourceByID(tables[0].ID, cfg.L.EquipmentTablesID)
		rel.AutoAbilities = namesToNamedAPIResources(cfg, cfg.E.AutoAbilities, table.RequiredAutoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.WpnTreasure, err = getResDbItemOne(cfg, ctx, cfg.E.Treasures, cw, ToIntOneNull(cfg.db.GetCelestialWeaponTreasureID))
		return err
	})

	g.Go(func() error {
		crestTreasures, err := getResourcesDbItem(cfg, ctx, cfg.E.Treasures, crest, cfg.db.GetKeyItemTreasureIDs)
		if err != nil {
			return err
		}
		rel.CrestTreasure = crestTreasures[0]
		return nil
	})

	g.Go(func() error {
		sigilQuests, err := getResourcesDbItem(cfg, ctx, cfg.E.Quests, sigil, cfg.db.GetKeyItemQuestIDs)
		if err != nil {
			return err
		}
		rel.SigilQuest = sigilQuests[0]
		return nil
	})

	err := g.Wait()
	if err != nil {
		return CelestialWeapon{}, err
	}

	return rel, nil
}
