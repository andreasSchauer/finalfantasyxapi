package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedMiscAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractMiscAbilities()
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
		l.json.miscAbilities[i].ID = row.ID
		l.MiscAbilities[Key(abilities[i])] = abilities[i]
		l.MiscAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMiscAbilities() ([]MiscAbility, error) {
	abilities := []MiscAbility{}
	var err error

	for i := range l.json.miscAbilities {
		ability := &l.json.miscAbilities[i]

		ability.Ability.ID, err = l.getHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		ability.TopmenuID, err = assignFKPtr(ability.Topmenu, l.Topmenus)
		if err != nil {
			return nil, err
		}

		ability.SubmenuID, err = assignFKPtr(ability.Submenu, l.Submenus)
		if err != nil {
			return nil, err
		}

		ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, l.Submenus)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completeMiscAbilities() error {
	for i := range l.json.miscAbilities {
		ability := &l.json.miscAbilities[i]

		err := l.completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		l.MiscAbilities[Key(ability)] = *ability
		l.MiscAbilitiesID[ability.ID] = *ability
	}

	return nil
}

func (l *Lookup) getMiscAbilityLearnedBy(ua MiscAbility) ([]CharacterClass, error) {
	return getResources(ua.LearnedBy, l.CharClasses)
}

func (l *Lookup) seedJuncMiscAbilitiesLearnedBy(qtx *database.Queries, ctx context.Context) error {
	const desc string = "misc abilities + learned by"
	jParams, err := processJunctions(l, desc, l.json.miscAbilities, l.getMiscAbilityLearnedBy)
	if err != nil {
		return err
	}

	return qtx.CreateMiscAbilitiesLearnedByJunctionBulk(ctx, database.CreateMiscAbilitiesLearnedByJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		MiscAbilityID:    jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
