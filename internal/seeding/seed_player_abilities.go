package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerAbility struct {
	ID int32
	Ability
	SubmenuID			*int32
	OpenSubmenuID		*int32
	StandardGridCharID	*int32
	ExpertGridCharID	*int32
	Description         *string             `json:"description"`
	Effect              string              `json:"effect"`
	RelatedStats		[]string			`json:"related_stats"`
	Topmenu             *string             `json:"topmenu"`
	Submenu				*string				`json:"submenu"`
	OpenSubmenu			*string				`json:"open_submenu"`
	LearnedBy			[]string			`json:"learned_by"`
	StandardGridPos		*string				`json:"standard_grid_pos"`
	ExpertGridPos		*string				`json:"expert_grid_pos"`
	CanUseOutsideBattle bool                `json:"can_use_outside_battle"`
	AeonLearnItem		*ItemAmount			`json:"aeon_learn_item"`
	MPCost              *int32              `json:"mp_cost"`
	Cursor              *string             `json:"cursor"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

func (p PlayerAbility) ToHashFields() []any {
	return []any{
		p.Ability.ID,
		derefOrNil(p.Description),
		p.Effect,
		derefOrNil(p.Topmenu),
		p.CanUseOutsideBattle,
		derefOrNil(p.MPCost),
		derefOrNil(p.Cursor),
		derefOrNil(p.SubmenuID),
		derefOrNil(p.OpenSubmenuID),
		derefOrNil(p.StandardGridCharID),
		derefOrNil(p.ExpertGridCharID),
		ObjPtrToHashID(p.AeonLearnItem),
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

func (l *lookup) seedPlayerAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/player_abilities.json"

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
				return err
			}

			dbPlayerAbility, err := qtx.CreatePlayerAbility(context.Background(), database.CreatePlayerAbilityParams{
				DataHash:            generateDataHash(playerAbility),
				AbilityID:           playerAbility.Ability.ID,
				Description:         getNullString(playerAbility.Description),
				Effect:              playerAbility.Effect,
				Topmenu:             nullTopmenuType(playerAbility.Topmenu),
				CanUseOutsideBattle: playerAbility.CanUseOutsideBattle,
				MpCost:              getNullInt32(playerAbility.MPCost),
				Cursor:              nullTargetType(playerAbility.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Player Ability: %s: %v", playerAbility.Name, err)
			}

			playerAbility.ID = dbPlayerAbility.ID
			key := createLookupKey(playerAbility.Ability)
			l.playerAbilities[key] = playerAbility
		}
		return nil
	})
}

func (l *lookup) seedPlayerAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/player_abilities.json"

	var playerAbilities []PlayerAbility

	err := loadJSONFile(string(srcPath), &playerAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range playerAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			ability, err := l.getPlayerAbility(abilityRef)
			if err != nil {
				return fmt.Errorf("ability %s: %v", createLookupKey(ability.Ability), err)
			}

			err = l.seedPlayerAbilityFKs(qtx, ability)
			if err != nil {
				return fmt.Errorf("ability %s: %v", createLookupKey(ability.Ability), err)
			}

			err = l.seedPlayerAbilityRelatedStats(qtx, ability)
			if err != nil {
				return fmt.Errorf("ability %s: %v", createLookupKey(ability.Ability), err)
			}

			err = l.seedPlayerAbilityLearnedBy(qtx, ability)
			if err != nil {
				return fmt.Errorf("ability %s: %v", createLookupKey(ability.Ability), err)
			}

			l.currentAbility = ability.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, ability.BattleInteractions)
			if err != nil {
				return fmt.Errorf("ability %s: %v", createLookupKey(ability.Ability), err)
			}
		}

		return nil
	})
}


func (l *lookup) seedPlayerAbilityFKs(qtx *database.Queries, ability PlayerAbility) error {
	var err error

	ability.SubmenuID, err = assignFKPtr(ability.Submenu, l.getSubmenu)
	if err != nil {
		return err
	}

	ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, l.getSubmenu)
	if err != nil {
		return err
	}

	ability.StandardGridCharID, err = assignFKPtr(ability.StandardGridPos, l.getCharacter)
	if err != nil {
		return err
	}

	ability.ExpertGridCharID, err = assignFKPtr(ability.ExpertGridPos, l.getCharacter)
	if err != nil {
		return err
	}

	ability.AeonLearnItem, err = seedObjPtrAssignFK(qtx, ability.AeonLearnItem, l.seedItemAmount)
	if err != nil {
		return err
	}

	err = qtx.UpdatePlayerAbility(context.Background(), database.UpdatePlayerAbilityParams{
		DataHash: 			generateDataHash(ability),
		SubmenuID: 			getNullInt32(ability.SubmenuID),
		OpenSubmenuID: 		getNullInt32(ability.OpenSubmenuID),
		StandardGridCharID: getNullInt32(ability.StandardGridCharID),
		ExpertGridCharID: 	getNullInt32(ability.ExpertGridCharID),
		AeonLearnItemID: 	ObjPtrToNullInt32ID(ability.AeonLearnItem),
		ID: 				ability.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't update player ability %s: %v", createLookupKey(ability.Ability), err)
	}

	return nil
}



func (l *lookup) seedPlayerAbilityRelatedStats(qtx *database.Queries, ability PlayerAbility) error {
	for _, jsonStat := range ability.RelatedStats {
		junction, err := createJunction(ability, jsonStat, l.getStat)
		if err != nil {
			return err
		}

		err = qtx.CreatePlayerAbilitiesRelatedStatsJunction(context.Background(), database.CreatePlayerAbilitiesRelatedStatsJunctionParams{
			DataHash: 			generateDataHash(junction),
			PlayerAbilityID: 	junction.ParentID,
			StatID: 			junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("ability %s: %v", createLookupKey(ability.Ability), err)
		}
	}

	return nil
}


func (l *lookup) seedPlayerAbilityLearnedBy(qtx *database.Queries, ability PlayerAbility) error {
	for _, charClass := range ability.LearnedBy {
		junction, err := createJunction(ability, charClass, l.getCharacterClass)
		if err != nil {
			return err
		}

		err = qtx.CreatePlayerAbilitiesLearnedByJunction(context.Background(), database.CreatePlayerAbilitiesLearnedByJunctionParams{
			DataHash: 			generateDataHash(junction),
			PlayerAbilityID: 	junction.ParentID,
			CharacterClassID: 	junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("ability %s: %v", createLookupKey(ability.Ability), err)
		}
	}

	return nil
}