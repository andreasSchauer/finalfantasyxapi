package seeding

import (
	"context"
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

func (l *Lookup) getOverdriveAbilityRelatedStats(oa OverdriveAbility) ([]Stat, error) {
	return getResources(oa.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncOverdriveAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrive abilities + related stats"
	jParams, err := processJunctions(l, desc, l.json.overdriveAbilities, l.getOverdriveAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateOverdriveAbilitiesRelatedStatsJunctionBulk(ctx, database.CreateOverdriveAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		OverdriveAbilityID: jParams.ParentIDs,
		StatID:             jParams.ChildIDs,
	})
}
