package seedhandlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3MiscAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractMiscAbilities()
	if err != nil {
		return err
	}

	params := database.CreateMiscAbilityBulkParams{
		DataHash:      make([]string, len(abilities)),
		AbilityID:     make([]int32, len(abilities)),
		Description:   make([]string, len(abilities)),
		Effect:        make([]string, len(abilities)),
		Cursor:        make([]database.NullTargetType, len(abilities)),
		TopmenuID:     make([]sql.NullInt32, len(abilities)),
		SubmenuID:     make([]sql.NullInt32, len(abilities)),
		OpenSubmenuID: make([]sql.NullInt32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.Description[i] = a.Description
		params.Effect[i] = a.Effect
		params.Cursor[i] = database.ToNullTargetType(a.Cursor)
		params.TopmenuID[i] = h.GetNullInt32(a.TopmenuID)
		params.SubmenuID[i] = h.GetNullInt32(a.SubmenuID)
		params.OpenSubmenuID[i] = h.GetNullInt32(a.OpenSubmenuID)
	}

	dbRows, err := qtx.CreateMiscAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create misc abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		sh.RawJSON.MiscAbilities[i].ID = row.ID
		lookup.L.MiscAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.MiscAbilitiesID[row.ID] = abilities[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMiscAbilities() ([]seedtypes.MiscAbility, error) {
	abilities := []seedtypes.MiscAbility{}
	var err error

	for i := range sh.RawJSON.MiscAbilities {
		ability := &sh.RawJSON.MiscAbilities[i]

		ability.Ability.ID, err = sh.GetHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		ability.TopmenuID, err = sh.AssignFKPtr(ability.Topmenu, lookup.L.Topmenus)
		if err != nil {
			return nil, err
		}

		ability.SubmenuID, err = sh.AssignFKPtr(ability.Submenu, lookup.L.Submenus)
		if err != nil {
			return nil, err
		}

		ability.OpenSubmenuID, err = sh.AssignFKPtr(ability.OpenSubmenu, lookup.L.Submenus)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return sh.DedupeRows(abilities, sh.Hashes), nil
}

func CompleteMiscAbilities() error {
	for i := range sh.RawJSON.MiscAbilities {
		ability := &sh.RawJSON.MiscAbilities[i]

		err := completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		lookup.L.MiscAbilities[seedtypes.Key(ability)] = *ability
		lookup.L.MiscAbilitiesID[ability.ID] = *ability
	}

	return nil
}
