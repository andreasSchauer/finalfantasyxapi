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

func Loop5PlayerAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractPlayerAbilities()
	if err != nil {
		return err
	}

	params := database.CreatePlayerAbilityBulkParams{
		DataHash:            make([]string, len(abilities)),
		AbilityID:           make([]int32, len(abilities)),
		Description:         make([]sql.NullString, len(abilities)),
		Effect:              make([]string, len(abilities)),
		Category:            make([]database.PlayerAbilityCategory, len(abilities)),
		CanUseOutsideBattle: make([]bool, len(abilities)),
		MpCost:              make([]int32, len(abilities)),
		Cursor:              make([]database.NullTargetType, len(abilities)),
		TopmenuID:           make([]sql.NullInt32, len(abilities)),
		SubmenuID:           make([]sql.NullInt32, len(abilities)),
		OpenSubmenuID:       make([]sql.NullInt32, len(abilities)),
		StdGridCharID:       make([]sql.NullInt32, len(abilities)),
		ExpGridCharID:       make([]sql.NullInt32, len(abilities)),
		AeonLearnItemID:     make([]sql.NullInt32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.Description[i] = h.GetNullString(a.Description)
		params.Effect[i] = a.Effect
		params.Category[i] = database.PlayerAbilityCategory(a.Category)
		params.CanUseOutsideBattle[i] = a.CanUseOutsideBattle
		params.MpCost[i] = a.MPCost
		params.Cursor[i] = database.ToNullTargetType(a.Cursor)
		params.TopmenuID[i] = h.GetNullInt32(a.TopmenuID)
		params.SubmenuID[i] = h.GetNullInt32(a.SubmenuID)
		params.OpenSubmenuID[i] = h.GetNullInt32(a.OpenSubmenuID)
		params.StdGridCharID[i] = h.GetNullInt32(a.StandardGridCharID)
		params.ExpGridCharID[i] = h.GetNullInt32(a.ExpertGridCharID)
		params.AeonLearnItemID[i] = h.ObjPtrToNullInt32ID(a.AeonLearnItem)
	}

	dbRows, err := qtx.CreatePlayerAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create player abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		sh.RawJSON.PlayerAbilities[i].ID = row.ID
		lookup.L.PlayerAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.PlayerAbilitiesID[row.ID] = abilities[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractPlayerAbilities() ([]seedtypes.PlayerAbility, error) {
	abilities := []seedtypes.PlayerAbility{}
	var err error

	for i := range sh.RawJSON.PlayerAbilities {
		ability := &sh.RawJSON.PlayerAbilities[i]

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

		ability.StandardGridCharID, err = sh.AssignFKPtr(ability.StandardGridPos, lookup.L.Characters)
		if err != nil {
			return nil, err
		}

		ability.ExpertGridCharID, err = sh.AssignFKPtr(ability.ExpertGridPos, lookup.L.Characters)
		if err != nil {
			return nil, err
		}

		if ability.AeonLearnItem != nil {
			ability.AeonLearnItem.ID, err = sh.GetHashID(ability.AeonLearnItem)
			if err != nil {
				return nil, err
			}
		}

		abilities = append(abilities, *ability)
	}

	return sh.DedupeRows(abilities, sh.Hashes), nil
}

func CompletePlayerAbilities() error {
	for i := range sh.RawJSON.PlayerAbilities {
		ability := &sh.RawJSON.PlayerAbilities[i]

		err := completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		lookup.L.PlayerAbilities[seedtypes.Key(ability)] = *ability
		lookup.L.PlayerAbilitiesID[ability.ID] = *ability
	}

	return nil
}
