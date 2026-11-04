package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getCelestialWeaponRelationships(cfg *Config, r *http.Request, cw seedtypes.CelestialWeapon) (CelestialWeapon, error) {
	crest, _ := lookup.GetResource(cw.KeyItemBase+" crest", lookup.L.KeyItems)
	sigil, _ := lookup.GetResource(cw.KeyItemBase+" sigil", lookup.L.KeyItems)
	equipment, _ := lookup.GetResource(cw, lookup.L.EquipmentNames)

	rel := CelestialWeapon{
		Equipment: nameToNamedAPIResource(cfg, cfg.E.Equipment, equipment.Name, nil),
		Crest:     nameToNamedAPIResource(cfg, cfg.E.KeyItems, crest.Name, nil),
		Sigil:     nameToNamedAPIResource(cfg, cfg.E.KeyItems, sigil.Name, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		tables, err := cfg.E.EquipmentTables.dbGetRelations(cfg, ctx, equipment, cfg.db.GetEquipmentEquipmentTableIDs)
		if err != nil {
			return err
		}
		table, _ := lookup.GetResourceByID(tables[0].ID, lookup.L.EquipmentTablesID)
		rel.AutoAbilities = namesToNamedAPIResources(cfg, cfg.E.AutoAbilities, table.RequiredAutoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.WpnTreasure, err = cfg.E.Treasures.dbGetRelation(cfg, ctx, cw, ToIntOneNull(cfg.db.GetCelestialWeaponTreasureID))
		return err
	})

	g.Go(func() error {
		crestTreasures, err := cfg.E.Treasures.dbGetRelations(cfg, ctx, crest, cfg.db.GetKeyItemTreasureIDs)
		if err != nil {
			return err
		}
		rel.CrestTreasure = crestTreasures[0]
		return nil
	})

	g.Go(func() error {
		sigilQuests, err := cfg.E.Quests.dbGetRelations(cfg, ctx, sigil, cfg.db.GetKeyItemQuestIDs)
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
