package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AutoAbility struct {
	ID                   int32
	GradRecoveryStatID   *int32
	OnHitElementID       *int32
	AddedPropertyID      *int32
	CnvrsnFromModID      *int32
	CnvrsnToModID        *int32
	RequiredItemAmountID *int32
	AddedElemAffinityID  *int32
	OnHitStatusID        *int32
	Name                 string             `json:"name"`
	Description          *string            `json:"description"`
	Effect               string             `json:"effect"`
	Type                 *string            `json:"type"`
	Category             string             `json:"category"`
	RelatedStats         []string           `json:"related_stats"`
	AbilityValue         *int32             `json:"ability_value"`
	RequiredItem         *ItemAmount        `json:"required_item"`
	LockedOutAbilities   []string           `json:"locked_out_abilities"`
	ActivationCondition  *string            `json:"activation_condition"`
	Counter              *string            `json:"counter"`
	GradualRecovery      *string            `json:"gradual_recovery"`
	AutoItemUse          []string           `json:"auto_item_use"`
	OnHitElement         *string            `json:"on_hit_element"`
	AddedElemAffinity    *ElementalAffinity `json:"added_elem_affinity"`
	OnHitStatus          *InflictedStatus   `json:"on_hit_status"`
	AddedStatusResists   []StatusResist     `json:"added_status_resists"`
	AddedStatusses       []string           `json:"added_statusses"`
	AddedProperty        *string            `json:"added_property"`
	ConversionFrom       *string            `json:"conversion_from"`
	ConversionTo         *string            `json:"conversion_to"`
	StatChanges          []StatChange       `json:"stat_changes"`
	ModifierChanges      []ModifierChange   `json:"modifier_changes"`
}

func (a AutoAbility) ToHashFields() []any {
	return []any{
		a.Name,
		derefOrNil(a.Description),
		a.Effect,
		derefOrNil(a.Type),
		a.Category,
		derefOrNil(a.AbilityValue),
		derefOrNil(a.RequiredItemAmountID),
		derefOrNil(a.ActivationCondition),
		derefOrNil(a.Counter),
		derefOrNil(a.GradRecoveryStatID),
		derefOrNil(a.OnHitElementID),
		derefOrNil(a.AddedElemAffinityID),
		derefOrNil(a.OnHitStatusID),
		derefOrNil(a.AddedPropertyID),
		derefOrNil(a.CnvrsnFromModID),
		derefOrNil(a.CnvrsnToModID),
	}
}

func (a AutoAbility) GetID() *int32 {
	return &a.ID
}


func (l *lookup) seedAutoAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/auto_abilities.json"

	var autoAbilities []AutoAbility
	err := loadJSONFile(string(srcPath), &autoAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, autoAbility := range autoAbilities {
			dbAutoAbility, err := qtx.CreateAutoAbility(context.Background(), database.CreateAutoAbilityParams{
				DataHash:            generateDataHash(autoAbility),
				Name:                autoAbility.Name,
				Description:         getNullString(autoAbility.Description),
				Effect:              autoAbility.Effect,
				Type:                nullEquipType(autoAbility.Type),
				Category:            database.AutoAbilityCategory(autoAbility.Category),
				AbilityValue:        getNullInt32(autoAbility.AbilityValue),
				ActivationCondition: nullAaActivationCondition(autoAbility.ActivationCondition),
				Counter:             nullCounterType(autoAbility.Counter),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Auto-Ability: %s: %v", autoAbility.Name, err)
			}

			autoAbility.ID = dbAutoAbility.ID
			l.autoAbilities[autoAbility.Name] = autoAbility
		}
		return nil
	})
}

func (l *lookup) createAutoAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/auto_abilities.json"

	var autoAbilities []AutoAbility
	err := loadJSONFile(string(srcPath), &autoAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAutoAbility := range autoAbilities {
			autoAbility, err := l.getAutoAbility(jsonAutoAbility.Name)
			if err != nil {
				return err
			}

			autoAbility, err = l.assignAutoAbilityFKs(qtx, autoAbility)
			if err != nil {
				return err
			}

			err = qtx.UpdateAutoAbility(context.Background(), database.UpdateAutoAbilityParams{
				DataHash:             generateDataHash(autoAbility),
				Name:                 autoAbility.Name,
				Description:          getNullString(autoAbility.Description),
				Effect:               autoAbility.Effect,
				Type:                 nullEquipType(autoAbility.Type),
				Category:             database.AutoAbilityCategory(autoAbility.Category),
				AbilityValue:         getNullInt32(autoAbility.AbilityValue),
				RequiredItemAmountID: ptrObjIDToNullInt32(autoAbility.RequiredItem),
				ActivationCondition:  nullAaActivationCondition(autoAbility.ActivationCondition),
				Counter:              nullCounterType(autoAbility.Counter),
				GradRcvryStatID:      getNullInt32(autoAbility.GradRecoveryStatID),
				OnHitElementID:       getNullInt32(autoAbility.OnHitElementID),
				AddedElemAffinityID:  ptrObjIDToNullInt32(autoAbility.AddedElemAffinity),
				OnHitStatusID:        ptrObjIDToNullInt32(autoAbility.OnHitStatus),
				AddedPropertyID:      getNullInt32(autoAbility.AddedPropertyID),
				CnvrsnFromModID:      getNullInt32(autoAbility.CnvrsnFromModID),
				CnvrsnToModID:        getNullInt32(autoAbility.CnvrsnToModID),
				ID:                   autoAbility.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't update auto ability: %s: %v", autoAbility.Name, err)
			}

			err = l.createAutoAbilityJunctions(qtx, autoAbility)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (l *lookup) assignAutoAbilityFKs(qtx *database.Queries, autoAbility AutoAbility) (AutoAbility, error) {
	var err error

	autoAbility.GradRecoveryStatID, err = assignFK(autoAbility.GradualRecovery, l.getStat)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.OnHitElementID, err = assignFK(autoAbility.OnHitElement, l.getElement)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.AddedPropertyID, err = assignFK(autoAbility.AddedProperty, l.getProperty)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.CnvrsnFromModID, err = assignFK(autoAbility.ConversionFrom, l.getModifier)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.CnvrsnToModID, err = assignFK(autoAbility.ConversionTo, l.getModifier)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.RequiredItem, err = assignFKSeed(qtx, autoAbility.RequiredItem, l.seedItemAmount)
	if err != nil {
		return AutoAbility{}, fmt.Errorf("auto ability: %s: %v", autoAbility.Name, err)
	}

	autoAbility.AddedElemAffinity, err = assignFKSeed(qtx, autoAbility.AddedElemAffinity, l.seedElementalAffinity)
	if err != nil {
		return AutoAbility{}, fmt.Errorf("auto ability: %s: %v", autoAbility.Name, err)
	}
	
	autoAbility.OnHitStatus, err = assignFKSeed(qtx, autoAbility.OnHitStatus, l.seedInflictedStatus)
	if err != nil {
		return AutoAbility{}, fmt.Errorf("auto ability: %s: %v", autoAbility.Name, err)
	}

	return autoAbility, nil
}


func (l *lookup) createAutoAbilityJunctions(qtx *database.Queries, autoAbility AutoAbility) error {
	relationShipFunctions := []func(*database.Queries, AutoAbility) error{
		l.createAutoAbilityRelatedStats,
		l.createAutoAbilityLockedOutAbilities,
		l.createAutoAbilityAutoItemUse,
		l.createAutoAbilityAddedStatusResists,
		l.createAutoAbilityAddedStatusses,
		l.createAutoAbilityStatChanges,
		l.createAutoAbilityModifierChanges,
	}

	for _, function := range relationShipFunctions {
		err := function(qtx, autoAbility)
		if err != nil {
			return fmt.Errorf("auto ability: %s: %v", autoAbility.Name, err)
		}
	}

	return nil
}


func (l *lookup) createAutoAbilityRelatedStats(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, jsonStat := range autoAbility.RelatedStats {
		junction, err := createJunction(autoAbility, jsonStat, l.getStat)
		if err != nil {
			return err
		}

		err = qtx.CreateAutoAbilityStatJunction(context.Background(), database.CreateAutoAbilityStatJunctionParams{
			DataHash:      generateDataHash(junction),
			AutoAbilityID: junction.ParentID,
			StatID:        junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create related stats: %v", err)
		}
	}

	return nil
}


func (l *lookup) createAutoAbilityLockedOutAbilities(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, jsonAbility := range autoAbility.LockedOutAbilities {
		lockedAbility, err := l.getAutoAbility(jsonAbility)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: autoAbility.ID,
			ChildID:  lockedAbility.ID,
		}

		err = qtx.CreateAutoAbilitySelfJunction(context.Background(), database.CreateAutoAbilitySelfJunctionParams{
			DataHash:        generateDataHash(junction),
			ParentAbilityID: junction.ParentID,
			ChildAbilityID:  junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create locked out abilities: %v", err)
		}
	}

	return nil
}

func (l *lookup) createAutoAbilityAutoItemUse(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, jsonItem := range autoAbility.AutoItemUse {
		item, err := l.getItem(jsonItem)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: autoAbility.ID,
			ChildID:  item.ID,
		}

		err = qtx.CreateAutoAbilityItemJunction(context.Background(), database.CreateAutoAbilityItemJunctionParams{
			DataHash:      generateDataHash(junction),
			AutoAbilityID: junction.ParentID,
			ItemID:        junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create auto item use: %v", err)
		}
	}

	return nil
}

func (l *lookup) createAutoAbilityAddedStatusses(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, jsonStatus := range autoAbility.AddedStatusses {
		condition, err := l.getStatusCondition(jsonStatus)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: autoAbility.ID,
			ChildID:  condition.ID,
		}

		err = qtx.CreateAutoAbilityStatusConditionJunction(context.Background(), database.CreateAutoAbilityStatusConditionJunctionParams{
			DataHash:          generateDataHash(junction),
			AutoAbilityID:     junction.ParentID,
			StatusConditionID: junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create added statusses: %v", err)
		}
	}

	return nil
}

func (l *lookup) createAutoAbilityAddedStatusResists(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, statusResist := range autoAbility.AddedStatusResists {
		dbStatusResist, err := l.seedStatusResist(qtx, statusResist)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: autoAbility.ID,
			ChildID:  dbStatusResist.ID,
		}

		err = qtx.CreateAutoAbilityStatusResistJunction(context.Background(), database.CreateAutoAbilityStatusResistJunctionParams{
			DataHash:       generateDataHash(junction),
			AutoAbilityID:  junction.ParentID,
			StatusResistID: junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create status resist junction for %s: %v", statusResist.StatusCondition, err)
		}
	}

	return nil
}

func (l *lookup) createAutoAbilityStatChanges(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, statChange := range autoAbility.StatChanges {
		dbStatChange, err := l.seedStatChange(qtx, statChange)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: autoAbility.ID,
			ChildID:  dbStatChange.ID,
		}

		err = qtx.CreateAutoAbilityStatChangeJunction(context.Background(), database.CreateAutoAbilityStatChangeJunctionParams{
			DataHash:      generateDataHash(junction),
			AutoAbilityID: junction.ParentID,
			StatChangeID:  junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *lookup) createAutoAbilityModifierChanges(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, modifierChange := range autoAbility.ModifierChanges {
		dbModifierChange, err := l.seedModifierChange(qtx, modifierChange)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: autoAbility.ID,
			ChildID:  dbModifierChange.ID,
		}

		err = qtx.CreateAutoAbilityModifierChangeJunction(context.Background(), database.CreateAutoAbilityModifierChangeJunctionParams{
			DataHash:         generateDataHash(junction),
			AutoAbilityID:    junction.ParentID,
			ModifierChangeID: junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
