package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getAbilities() []Ability {
	abilities := []Ability{}

	for _, ability := range l.json.playerAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, ability := range l.json.overdriveAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, item := range l.json.items {
		if len(item.BattleInteractions) == 0 {
			continue
		}
		abilities = append(abilities, item.Ability)
	}

	for _, command := range l.json.triggerCommands {
		abilities = append(abilities, command.Ability)
	}

	for _, ability := range l.json.miscAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, ability := range l.json.enemyAbilities {
		abilities = append(abilities, ability.Ability)
	}

	return abilities
}

func (l *Lookup) getAbilityBattleInteractions(a Ability) ([]BattleInteraction, error) {
	switch a.Type {
	case database.AbilityTypePlayerAbility:
		pa, err := GetResource(a.GetAbilityRef().Untyped(), l.PlayerAbilities)
		if err != nil {
			return nil, err
		}
		return pa.BattleInteractions, nil

	case database.AbilityTypeOverdriveAbility:
		oa, err := GetResource(a.GetAbilityRef().Untyped(), l.OverdriveAbilities)
		if err != nil {
			return nil, err
		}
		return oa.BattleInteractions, nil

	case database.AbilityTypeItemAbility:
		ia, err := GetResource(a.Name, l.ItemAbilities)
		if err != nil {
			return nil, err
		}
		return ia.BattleInteractions, nil

	case database.AbilityTypeTriggerCommand:
		tc, err := GetResource(a.GetAbilityRef().Untyped(), l.TriggerCommands)
		if err != nil {
			return nil, err
		}
		return tc.BattleInteractions, nil

	case database.AbilityTypeMiscAbility:
		ua, err := GetResource(a.GetAbilityRef().Untyped(), l.MiscAbilities)
		if err != nil {
			return nil, err
		}
		return ua.BattleInteractions, nil

	case database.AbilityTypeEnemyAbility:
		ea, err := GetResource(a.GetAbilityRef().Untyped(), l.EnemyAbilities)
		if err != nil {
			return nil, err
		}
		return ea.BattleInteractions, nil

	default:
		return nil, fmt.Errorf("%s has no type", a)
	}
}

func (l *Lookup) seedJuncAbilitiesBattleInteractions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "abilities + battle interactions"
	jParams, err := processJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions)
	if err != nil {
		return err
	}

	return qtx.CreateAbilitiesBattleInteractionsJunctionBulk(ctx, database.CreateAbilitiesBattleInteractionsJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.ParentIDs,
		BattleInteractionID: jParams.ChildIDs,
	})
}
