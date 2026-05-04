package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop5SeedPlayerAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractPlayerAbilities()
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
		params.DataHash[i] = generateDataHash(a)
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
		l.json.playerAbilities[i].ID = row.ID
		key := Key(abilities[i])
		l.PlayerAbilities[key] = abilities[i]
		l.PlayerAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractPlayerAbilities() ([]PlayerAbility, error) {
	abilities := []PlayerAbility{}
	var err error

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

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

		ability.StandardGridCharID, err = assignFKPtr(ability.StandardGridPos, l.Characters)
		if err != nil {
			return nil, err
		}

		ability.ExpertGridCharID, err = assignFKPtr(ability.ExpertGridPos, l.Characters)
		if err != nil {
			return nil, err
		}

		if ability.AeonLearnItem != nil {
			ability.AeonLearnItem.ID, err = l.getHashID(ability.AeonLearnItem)
			if err != nil {
				return nil, err
			}
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completePlayerAbilities() error {
	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		err := l.completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		l.PlayerAbilities[Key(ability)] = *ability
		l.PlayerAbilitiesID[ability.ID] = *ability
	}

	return nil
}

func (l *Lookup) getPlayerAbilityLearnedBy(pa PlayerAbility) ([]CharacterClass, error) {
	return getResources(pa.LearnedBy, l.CharClasses)
}

func (l *Lookup) seedJuncPlayerAbilitiesLearnedBy(qtx *database.Queries, ctx context.Context) error {
	const desc string = "player abilities + learned by"
	jParams, err := processJunctions(l, desc, l.json.playerAbilities, l.getPlayerAbilityLearnedBy)
	if err != nil {
		return err
	}

	return qtx.CreatePlayerAbilitiesLearnedByJunctionBulk(ctx, database.CreatePlayerAbilitiesLearnedByJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		PlayerAbilityID:  jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}

func (l *Lookup) getPlayerAbilityRelatedStats(pa PlayerAbility) ([]Stat, error) {
	return getResources(pa.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncPlayerAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "player abilites + related stats"
	jParams, err := processJunctions(l, desc, l.json.playerAbilities, l.getPlayerAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreatePlayerAbilitiesRelatedStatsJunctionBulk(ctx, database.CreatePlayerAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		PlayerAbilityID: jParams.ParentIDs,
		StatID:          jParams.ChildIDs,
	})
}
