package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getCelestialWeaponRelationships(r *http.Request, cw seedtypes.CelestialWeapon) (CelestialWeapon, error) {
	crest, _ := lookup.GetResource(cw.KeyItemBase+" crest", lookup.L.KeyItems)
	sigil, _ := lookup.GetResource(cw.KeyItemBase+" sigil", lookup.L.KeyItems)
	equipment, _ := lookup.GetResource(cw, lookup.L.EquipmentNames)

	rel := CelestialWeapon{
		Equipment: nameToNamedAPIResource(E.Equipment, equipment.Name, nil),
		Crest:     nameToNamedAPIResource(E.KeyItems, crest.Name, nil),
		Sigil:     nameToNamedAPIResource(E.KeyItems, sigil.Name, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		tables, err := E.EquipmentTables.dbGetRelations(ctx, equipment, config.Cfg.DB.GetEquipmentEquipmentTableIDs)
		if err != nil {
			return err
		}
		table, _ := lookup.GetResourceByID(tables[0].ID, lookup.L.EquipmentTablesID)
		rel.AutoAbilities = namesToNamedAPIResources(E.AutoAbilities, table.RequiredAutoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.WpnTreasure, err = E.Treasures.dbGetRelation(ctx, cw, ToIntOneNull(config.Cfg.DB.GetCelestialWeaponTreasureID))
		return err
	})

	g.Go(func() error {
		crestTreasures, err := E.Treasures.dbGetRelations(ctx, crest, config.Cfg.DB.GetKeyItemTreasureIDs)
		if err != nil {
			return err
		}
		rel.CrestTreasure = crestTreasures[0]
		return nil
	})

	g.Go(func() error {
		sigilQuests, err := E.Quests.dbGetRelations(ctx, sigil, config.Cfg.DB.GetKeyItemQuestIDs)
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
