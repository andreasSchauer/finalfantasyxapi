package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetCelestialWeaponRelationships(r *http.Request, cw seedtypes.CelestialWeapon) (apitypes.CelestialWeapon, error) {
	crest, _ := lookup.GetResource(cw.KeyItemBase+" crest", lookup.L.KeyItems)
	sigil, _ := lookup.GetResource(cw.KeyItemBase+" sigil", lookup.L.KeyItems)
	equipment, _ := lookup.GetResource(cw, lookup.L.EquipmentNames)

	rel := apitypes.CelestialWeapon{
		Equipment: handlerinput.NameToNamedAPIResource(apilookups.E.Equipment, equipment.Name, nil),
		Crest:     handlerinput.NameToNamedAPIResource(apilookups.E.KeyItems, crest.Name, nil),
		Sigil:     handlerinput.NameToNamedAPIResource(apilookups.E.KeyItems, sigil.Name, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		tables, err := apilookups.E.EquipmentTables.DBGetRelations(ctx, equipment, config.Cfg.DB.GetEquipmentEquipmentTableIDs)
		if err != nil {
			return err
		}
		table, _ := lookup.GetResourceByID(tables[0].ID, lookup.L.EquipmentTablesID)
		rel.AutoAbilities = handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, table.RequiredAutoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.WpnTreasure, err = apilookups.E.Treasures.DBGetRelation(ctx, cw, handlerinput.ToIntOneNull(config.Cfg.DB.GetCelestialWeaponTreasureID))
		return err
	})

	g.Go(func() error {
		crestTreasures, err := apilookups.E.Treasures.DBGetRelations(ctx, crest, config.Cfg.DB.GetKeyItemTreasureIDs)
		if err != nil {
			return err
		}
		rel.CrestTreasure = crestTreasures[0]
		return nil
	})

	g.Go(func() error {
		sigilQuests, err := apilookups.E.Quests.DBGetRelations(ctx, sigil, config.Cfg.DB.GetKeyItemQuestIDs)
		if err != nil {
			return err
		}
		rel.SigilQuest = sigilQuests[0]
		return nil
	})

	err := g.Wait()
	if err != nil {
		return apitypes.CelestialWeapon{}, err
	}

	return rel, nil
}
