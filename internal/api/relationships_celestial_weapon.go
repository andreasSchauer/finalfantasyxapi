package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getCelestialWeaponRelationships(server *Server, r *http.Request, cw seedtypes.CelestialWeapon) (CelestialWeapon, error) {
	crest, _ := lookup.GetResource(cw.KeyItemBase+" crest", lookup.L.KeyItems)
	sigil, _ := lookup.GetResource(cw.KeyItemBase+" sigil", lookup.L.KeyItems)
	equipment, _ := lookup.GetResource(cw, lookup.L.EquipmentNames)

	rel := CelestialWeapon{
		Equipment: nameToNamedAPIResource(server, E.Equipment, equipment.Name, nil),
		Crest:     nameToNamedAPIResource(server, E.KeyItems, crest.Name, nil),
		Sigil:     nameToNamedAPIResource(server, E.KeyItems, sigil.Name, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		tables, err := E.EquipmentTables.dbGetRelations(server, ctx, equipment, server.Cfg.DB.GetEquipmentEquipmentTableIDs)
		if err != nil {
			return err
		}
		table, _ := lookup.GetResourceByID(tables[0].ID, lookup.L.EquipmentTablesID)
		rel.AutoAbilities = namesToNamedAPIResources(server, E.AutoAbilities, table.RequiredAutoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.WpnTreasure, err = E.Treasures.dbGetRelation(server, ctx, cw, ToIntOneNull(server.Cfg.DB.GetCelestialWeaponTreasureID))
		return err
	})

	g.Go(func() error {
		crestTreasures, err := E.Treasures.dbGetRelations(server, ctx, crest, server.Cfg.DB.GetKeyItemTreasureIDs)
		if err != nil {
			return err
		}
		rel.CrestTreasure = crestTreasures[0]
		return nil
	})

	g.Go(func() error {
		sigilQuests, err := E.Quests.dbGetRelations(server, ctx, sigil, server.Cfg.DB.GetKeyItemQuestIDs)
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
