package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type PlayerAbility struct {
	ID int32
	Ability
	SubmenuID           *int32
	OpenSubmenuID       *int32
	StandardGridCharID  *int32
	ExpertGridCharID    *int32
	Description         *string             `json:"description"`
	Effect              string              `json:"effect"`
	RelatedStats        []string            `json:"related_stats"`
	Category			string				`json:"category"`
	Topmenu             *string             `json:"topmenu"`
	Submenu             *string             `json:"submenu"`
	OpenSubmenu         *string             `json:"open_submenu"`
	LearnedBy           []string            `json:"learned_by"`
	StandardGridPos     *string             `json:"standard_grid_pos"`
	ExpertGridPos       *string             `json:"expert_grid_pos"`
	CanUseOutsideBattle bool                `json:"can_use_outside_battle"`
	AeonLearnItem       *ItemAmount         `json:"aeon_learn_item"`
	MPCost              *int32              `json:"mp_cost"`
	Cursor              *string             `json:"cursor"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

func (p PlayerAbility) ToHashFields() []any {
	return []any{
		p.Ability.ID,
		h.DerefOrNil(p.Description),
		p.Effect,
		p.Category,
		h.DerefOrNil(p.Topmenu),
		p.CanUseOutsideBattle,
		h.DerefOrNil(p.MPCost),
		h.DerefOrNil(p.Cursor),
		h.DerefOrNil(p.SubmenuID),
		h.DerefOrNil(p.OpenSubmenuID),
		h.DerefOrNil(p.StandardGridCharID),
		h.DerefOrNil(p.ExpertGridCharID),
		h.ObjPtrToID(p.AeonLearnItem),
	}
}

func (p PlayerAbility) ToKeyFields() []any {
	return []any{
		p.Ability.Name,
		h.DerefOrNil(p.Ability.Version),
	}
}

func (p PlayerAbility) GetID() int32 {
	return p.ID
}

func (p PlayerAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        p.Name,
		Version:     p.Version,
		AbilityType: string(database.AbilityTypePlayerAbility),
	}
}

func (p PlayerAbility) Error() string {
	return fmt.Sprintf("player ability %s, version %v", p.Name, h.DerefOrNil(p.Version))
}

func (p PlayerAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            p.ID,
		Name:          p.Name,
		Version:       p.Version,
		Specification: p.Specification,
	}
}

func (l *Lookup) seedPlayerAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/player_abilities.json"

	var playerAbilities []PlayerAbility

	err := loadJSONFile(string(srcPath), &playerAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, playerAbility := range playerAbilities {
			var err error
			playerAbility.Type = database.AbilityTypePlayerAbility

			playerAbility.Ability, err = seedObjAssignID(qtx, playerAbility.Ability, l.seedAbility)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}

			dbPlayerAbility, err := qtx.CreatePlayerAbility(context.Background(), database.CreatePlayerAbilityParams{
				DataHash:            generateDataHash(playerAbility),
				AbilityID:           playerAbility.Ability.ID,
				Description:         h.GetNullString(playerAbility.Description),
				Effect:              playerAbility.Effect,
				Category: 			 database.PlayerAbilityCategory(playerAbility.Category),
				Topmenu:             h.NullTopmenuType(playerAbility.Topmenu),
				CanUseOutsideBattle: playerAbility.CanUseOutsideBattle,
				MpCost:              h.GetNullInt32(playerAbility.MPCost),
				Cursor:              h.NullTargetType(playerAbility.Cursor),
			})
			if err != nil {
				return h.NewErr(playerAbility.Error(), err, "couldn't create player ability")
			}

			playerAbility.ID = dbPlayerAbility.ID
			key := CreateLookupKey(playerAbility)
			l.PlayerAbilities[key] = playerAbility
			l.PlayerAbilitiesID[playerAbility.ID] = playerAbility
		}
		return nil
	})
}

func (l *Lookup) seedPlayerAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/player_abilities.json"

	var playerAbilities []PlayerAbility

	err := loadJSONFile(string(srcPath), &playerAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range playerAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			playerAbility, err := GetResource(abilityRef.Untyped(), l.PlayerAbilities)
			if err != nil {
				return err
			}

			err = l.seedPlayerAbilityFKs(qtx, playerAbility)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}

			err = l.seedPlayerAbilityRelatedStats(qtx, playerAbility)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}

			err = l.seedPlayerAbilityLearnedBy(qtx, playerAbility)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}

			l.currentAbility = playerAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, playerAbility.BattleInteractions)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedPlayerAbilityFKs(qtx *database.Queries, ability PlayerAbility) error {
	var err error

	ability.SubmenuID, err = assignFKPtr(ability.Submenu, l.Submenus)
	if err != nil {
		return err
	}

	ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, l.Submenus)
	if err != nil {
		return err
	}

	ability.StandardGridCharID, err = assignFKPtr(ability.StandardGridPos, l.Characters)
	if err != nil {
		return err
	}

	ability.ExpertGridCharID, err = assignFKPtr(ability.ExpertGridPos, l.Characters)
	if err != nil {
		return err
	}

	ability.AeonLearnItem, err = seedObjPtrAssignFK(qtx, ability.AeonLearnItem, l.seedItemAmount)
	if err != nil {
		return err
	}

	err = qtx.UpdatePlayerAbility(context.Background(), database.UpdatePlayerAbilityParams{
		DataHash:           generateDataHash(ability),
		SubmenuID:          h.GetNullInt32(ability.SubmenuID),
		OpenSubmenuID:      h.GetNullInt32(ability.OpenSubmenuID),
		StandardGridCharID: h.GetNullInt32(ability.StandardGridCharID),
		ExpertGridCharID:   h.GetNullInt32(ability.ExpertGridCharID),
		AeonLearnItemID:    h.ObjPtrToNullInt32ID(ability.AeonLearnItem),
		ID:                 ability.ID,
	})
	if err != nil {
		return h.NewErr("", err, "couldn't update player ability")
	}

	return nil
}

func (l *Lookup) seedPlayerAbilityRelatedStats(qtx *database.Queries, ability PlayerAbility) error {
	for _, jsonStat := range ability.RelatedStats {
		junction, err := createJunction(ability, jsonStat, l.Stats)
		if err != nil {
			return err
		}

		err = qtx.CreatePlayerAbilitiesRelatedStatsJunction(context.Background(), database.CreatePlayerAbilitiesRelatedStatsJunctionParams{
			DataHash:        generateDataHash(junction),
			PlayerAbilityID: junction.ParentID,
			StatID:          junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonStat, err, "couldn't junction related stat")
		}
	}

	return nil
}

func (l *Lookup) seedPlayerAbilityLearnedBy(qtx *database.Queries, ability PlayerAbility) error {
	for _, charClass := range ability.LearnedBy {
		junction, err := createJunction(ability, charClass, l.CharClasses)
		if err != nil {
			return err
		}

		err = qtx.CreatePlayerAbilitiesLearnedByJunction(context.Background(), database.CreatePlayerAbilitiesLearnedByJunctionParams{
			DataHash:         generateDataHash(junction),
			PlayerAbilityID:  junction.ParentID,
			CharacterClassID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(charClass, err, "couldn't junction 'learned by' class")
		}
	}

	return nil
}
