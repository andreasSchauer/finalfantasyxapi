package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop3SeedMiscAbilities(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = generateDataHash(a)
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
		rawJSON.miscAbilities[i].ID = row.ID
		lookup.L.MiscAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.MiscAbilitiesID[row.ID] = abilities[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMiscAbilities() ([]seedtypes.MiscAbility, error) {
	abilities := []seedtypes.MiscAbility{}
	var err error

	for i := range rawJSON.miscAbilities {
		ability := &rawJSON.miscAbilities[i]

		ability.Ability.ID, err = GetHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		ability.TopmenuID, err = assignFKPtr(ability.Topmenu, lookup.L.Topmenus)
		if err != nil {
			return nil, err
		}

		ability.SubmenuID, err = assignFKPtr(ability.Submenu, lookup.L.Submenus)
		if err != nil {
			return nil, err
		}

		ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, lookup.L.Submenus)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, hashes), nil
}

func completeMiscAbilities() error {
	for i := range rawJSON.miscAbilities {
		ability := &rawJSON.miscAbilities[i]

		err := completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		lookup.L.MiscAbilities[seedtypes.Key(ability)] = *ability
		lookup.L.MiscAbilitiesID[ability.ID] = *ability
	}

	return nil
}

func getMiscAbilityLearnedBy(ua seedtypes.MiscAbility) ([]seedtypes.CharacterClass, error) {
	return lookup.GetResources(ua.LearnedBy, lookup.L.CharClasses)
}

func seedJuncMiscAbilitiesLearnedBy(qtx *database.Queries, ctx context.Context) error {
	const desc string = "misc abilities + learned by"
	jParams, err := processJunctions(desc, rawJSON.miscAbilities, getMiscAbilityLearnedBy)
	if err != nil {
		return err
	}

	return qtx.CreateMiscAbilitiesLearnedByJunctionBulk(ctx, database.CreateMiscAbilitiesLearnedByJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		MiscAbilityID:    jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
