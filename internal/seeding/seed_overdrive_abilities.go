package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OverdriveAbility struct {
	ID int32
	Ability
	Overdrive          LookupObject        `json:"overdrive"` // not meant for seeding
	RelatedStats       []string            `json:"related_stats"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (o OverdriveAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", o),
		o.Ability.ID,
	}
}

func (o OverdriveAbility) ToKeyFields() []any {
	return []any{
		o.Ability.Name,
		h.DerefOrNil(o.Ability.Version),
	}
}

func (o OverdriveAbility) GetID() int32 {
	return o.ID
}

func (o OverdriveAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        o.Name,
		Version:     o.Version,
		AbilityType: string(database.AbilityTypeOverdriveAbility),
	}
}

func (o OverdriveAbility) Error() string {
	return fmt.Sprintf("overdrive ability '%s'", h.NameToString(o.Name, o.Version, o.Specification))
}

func (o OverdriveAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            o.ID,
		Name:          o.Name,
		Version:       o.Version,
		Specification: o.Specification,
	}
}

func (l *Lookup) seedOverdriveAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrive_abilities.json"

	var overdriveAbilities []OverdriveAbility

	err := loadJSONFile(string(srcPath), &overdriveAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, overdriveAbility := range overdriveAbilities {
			var err error
			overdriveAbility.Type = database.AbilityTypeOverdriveAbility

			overdriveAbility.Ability, err = seedObjAssignID(qtx, overdriveAbility.Ability, l.seedAbility)
			if err != nil {
				return h.NewErr(overdriveAbility.Error(), err)
			}

			dbOverdriveAbility, err := qtx.CreateOverdriveAbility(context.Background(), database.CreateOverdriveAbilityParams{
				DataHash:  generateDataHash(overdriveAbility),
				AbilityID: overdriveAbility.Ability.ID,
			})
			if err != nil {
				return h.NewErr(overdriveAbility.Error(), err, "couldn't create overdrive ability")
			}

			overdriveAbility.ID = dbOverdriveAbility.ID
			key := Key(overdriveAbility)
			l.OverdriveAbilities[key] = overdriveAbility
			l.OverdriveAbilitiesID[overdriveAbility.ID] = overdriveAbility
		}
		return nil
	})
}

func (l *Lookup) seedOverdriveAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrive_abilities.json"

	var overdriveAbilities []OverdriveAbility

	err := loadJSONFile(string(srcPath), &overdriveAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range overdriveAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			overdriveAbility, err := GetResource(abilityRef.Untyped(), l.OverdriveAbilities)
			if err != nil {
				return h.NewErr(abilityRef.Error(), err)
			}

			err = l.seedOverdriveAbilityRelatedStats(qtx, overdriveAbility)
			if err != nil {
				return h.NewErr(overdriveAbility.Error(), err)
			}

			l.currentAbility = overdriveAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, overdriveAbility.BattleInteractions)
			if err != nil {
				return h.NewErr(overdriveAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedOverdriveAbilityRelatedStats(qtx *database.Queries, ability OverdriveAbility) error {
	for _, jsonStat := range ability.RelatedStats {
		junction, err := createJunction(ability, jsonStat, l.Stats)
		if err != nil {
			return err
		}

		err = qtx.CreateOverdriveAbilitiesRelatedStatsJunction(context.Background(), database.CreateOverdriveAbilitiesRelatedStatsJunctionParams{
			DataHash:           generateDataHash(junction),
			OverdriveAbilityID: junction.ParentID,
			StatID:             junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonStat, err, "couldn't junction related stat")
		}
	}

	return nil
}

func (l *Lookup) loop3SeedOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractOverdriveAbilities()
	if err != nil {
		return err
	}

	params := database.CreateOverdriveAbilityBulkParams{
		DataHash:  make([]string, len(abilities)),
		AbilityID: make([]int32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
	}

	dbRows, err := qtx.CreateOverdriveAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrive abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		l.json.overdriveAbilities[i].ID = row.ID
		key := Key(abilities[i])
		l.OverdriveAbilities[key] = abilities[i]
		l.OverdriveAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractOverdriveAbilities() ([]OverdriveAbility, error) {
	abilities := []OverdriveAbility{}
	var err error

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		ability.Ability.ID, err = l.getHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completeOverdriveAbilities() error {
	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		err := l.completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		l.OverdriveAbilities[Key(ability)] = *ability
		l.OverdriveAbilitiesID[ability.ID] = *ability
	}

	return nil
}