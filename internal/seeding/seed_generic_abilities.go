package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type GenericAbility struct {
	ID int32
	Ability
	SubmenuID           *int32
	OpenSubmenuID       *int32
	Description         *string             `json:"description"`
	Effect              string              `json:"effect"`
	RelatedStats        []string            `json:"related_stats"`
	Topmenu             *string             `json:"topmenu"`
	Submenu             *string             `json:"submenu"`
	OpenSubmenu         *string             `json:"open_submenu"`
	LearnedBy           []string            `json:"learned_by"`
	Cursor              *string             `json:"cursor"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

func (g GenericAbility) ToHashFields() []any {
	return []any{
		g.Ability.ID,
		h.DerefOrNil(g.Description),
		g.Effect,
		h.DerefOrNil(g.Topmenu),
		h.DerefOrNil(g.Cursor),
		h.DerefOrNil(g.SubmenuID),
		h.DerefOrNil(g.OpenSubmenuID),
	}
}

func (g GenericAbility) ToKeyFields() []any {
	return []any{
		g.Ability.Name,
		h.DerefOrNil(g.Ability.Version),
	}
}

func (g GenericAbility) GetID() int32 {
	return g.ID
}

func (g GenericAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        g.Name,
		Version:     g.Version,
		AbilityType: string(database.AbilityTypeGenericAbility),
	}
}

func (g GenericAbility) Error() string {
	return fmt.Sprintf("generic ability %s, version %v", g.Name, h.DerefOrNil(g.Version))
}

func (g GenericAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            g.ID,
		Name:          g.Name,
		Version:       g.Version,
		Specification: g.Specification,
	}
}

func (l *Lookup) seedGenericAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/generic_abilities.json"

	var genericAbilities []GenericAbility

	err := loadJSONFile(string(srcPath), &genericAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, genericAbility := range genericAbilities {
			var err error
			genericAbility.Type = database.AbilityTypeGenericAbility

			genericAbility.Ability, err = seedObjAssignID(qtx, genericAbility.Ability, l.seedAbility)
			if err != nil {
				return h.NewErr(genericAbility.Error(), err)
			}

			dbGenericAbility, err := qtx.CreateGenericAbility(context.Background(), database.CreateGenericAbilityParams{
				DataHash:            generateDataHash(genericAbility),
				AbilityID:           genericAbility.Ability.ID,
				Description:         h.GetNullString(genericAbility.Description),
				Effect:              genericAbility.Effect,
				Topmenu:             h.NullTopmenuType(genericAbility.Topmenu),
				Cursor:              h.NullTargetType(genericAbility.Cursor),
			})
			if err != nil {
				return h.NewErr(genericAbility.Error(), err, "couldn't create generic ability")
			}

			genericAbility.ID = dbGenericAbility.ID
			key := CreateLookupKey(genericAbility)
			l.GenericAbilities[key] = genericAbility
			l.GenericAbilitiesID[genericAbility.ID] = genericAbility
		}
		return nil
	})
}

func (l *Lookup) seedGenericAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/generic_abilities.json"

	var genericAbilities []GenericAbility

	err := loadJSONFile(string(srcPath), &genericAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range genericAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			genericAbility, err := GetResource(abilityRef.Untyped(), l.GenericAbilities)
			if err != nil {
				return err
			}

			err = l.seedGenericAbilityFKs(qtx, genericAbility)
			if err != nil {
				return h.NewErr(genericAbility.Error(), err)
			}

			err = l.seedGenericAbilityRelatedStats(qtx, genericAbility)
			if err != nil {
				return h.NewErr(genericAbility.Error(), err)
			}

			err = l.seedGenericAbilityLearnedBy(qtx, genericAbility)
			if err != nil {
				return h.NewErr(genericAbility.Error(), err)
			}

			l.currentAbility = genericAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, genericAbility.BattleInteractions)
			if err != nil {
				return h.NewErr(genericAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedGenericAbilityFKs(qtx *database.Queries, ability GenericAbility) error {
	var err error

	ability.SubmenuID, err = assignFKPtr(ability.Submenu, l.Submenus)
	if err != nil {
		return err
	}

	ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, l.Submenus)
	if err != nil {
		return err
	}

	err = qtx.UpdateGenericAbility(context.Background(), database.UpdateGenericAbilityParams{
		DataHash:           generateDataHash(ability),
		SubmenuID:          h.GetNullInt32(ability.SubmenuID),
		OpenSubmenuID:      h.GetNullInt32(ability.OpenSubmenuID),
		ID:                 ability.ID,
	})
	if err != nil {
		return h.NewErr("", err, "couldn't update generic ability")
	}

	return nil
}

func (l *Lookup) seedGenericAbilityRelatedStats(qtx *database.Queries, ability GenericAbility) error {
	for _, jsonStat := range ability.RelatedStats {
		junction, err := createJunction(ability, jsonStat, l.Stats)
		if err != nil {
			return err
		}

		err = qtx.CreateGenericAbilitiesRelatedStatsJunction(context.Background(), database.CreateGenericAbilitiesRelatedStatsJunctionParams{
			DataHash:        	generateDataHash(junction),
			GenericAbilityID: 	junction.ParentID,
			StatID:          	junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonStat, err, "couldn't junction related stat")
		}
	}

	return nil
}

func (l *Lookup) seedGenericAbilityLearnedBy(qtx *database.Queries, ability GenericAbility) error {
	for _, charClass := range ability.LearnedBy {
		junction, err := createJunction(ability, charClass, l.CharClasses)
		if err != nil {
			return err
		}

		err = qtx.CreateGenericAbilitiesLearnedByJunction(context.Background(), database.CreateGenericAbilitiesLearnedByJunctionParams{
			DataHash:         	generateDataHash(junction),
			GenericAbilityID:  	junction.ParentID,
			CharacterClassID: 	junction.ChildID,
		})
		if err != nil {
			return h.NewErr(charClass, err, "couldn't junction 'learned by' class")
		}
	}

	return nil
}
