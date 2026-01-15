package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AutoAbility struct {
	ID                  int32
	GradRecoveryStatID  *int32
	OnHitElementID      *int32
	AddedPropertyID     *int32
	CnvrsnFromModID     *int32
	CnvrsnToModID       *int32
	Name                string           `json:"name"`
	Description         *string          `json:"description"`
	Effect              string           `json:"effect"`
	Type                *string          `json:"type"`
	Category            string           `json:"category"`
	RelatedStats        []string         `json:"related_stats"`
	AbilityValue        *int32           `json:"ability_value"`
	RequiredItem        *ItemAmount      `json:"required_item"`
	LockedOutAbilities  []string         `json:"locked_out_abilities"`
	ActivationCondition *string          `json:"activation_condition"`
	Counter             *string          `json:"counter"`
	GradualRecovery     *string          `json:"gradual_recovery"`
	AutoItemUse         []string         `json:"auto_item_use"`
	OnHitElement        *string          `json:"on_hit_element"`
	AddedElemResist     *ElementalResist `json:"added_elem_resist"`
	OnHitStatus         *InflictedStatus `json:"on_hit_status"`
	AddedStatusResists  []StatusResist   `json:"added_status_resists"`
	AddedStatusses      []string         `json:"added_statusses"`
	AddedProperty       *string          `json:"added_property"`
	ConversionFrom      *string          `json:"conversion_from"`
	ConversionTo        *string          `json:"conversion_to"`
	StatChanges         []StatChange     `json:"stat_changes"`
	ModifierChanges     []ModifierChange `json:"modifier_changes"`
}

func (a AutoAbility) ToHashFields() []any {
	return []any{
		a.Name,
		h.DerefOrNil(a.Description),
		a.Effect,
		h.DerefOrNil(a.Type),
		a.Category,
		h.DerefOrNil(a.AbilityValue),
		h.ObjPtrToID(a.RequiredItem),
		h.DerefOrNil(a.ActivationCondition),
		h.DerefOrNil(a.Counter),
		h.DerefOrNil(a.GradRecoveryStatID),
		h.DerefOrNil(a.OnHitElementID),
		h.ObjPtrToID(a.AddedElemResist),
		h.ObjPtrToID(a.OnHitStatus),
		h.DerefOrNil(a.AddedPropertyID),
		h.DerefOrNil(a.CnvrsnFromModID),
		h.DerefOrNil(a.CnvrsnToModID),
	}
}

func (a AutoAbility) GetID() int32 {
	return a.ID
}

func (a AutoAbility) Error() string {
	return fmt.Sprintf("auto ability %s", a.Name)
}

func (l *Lookup) seedAutoAbilities(db *database.Queries, dbConn *sql.DB) error {
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
				Description:         h.GetNullString(autoAbility.Description),
				Effect:              autoAbility.Effect,
				Type:                h.NullEquipType(autoAbility.Type),
				Category:            database.AutoAbilityCategory(autoAbility.Category),
				AbilityValue:        h.GetNullInt32(autoAbility.AbilityValue),
				ActivationCondition: h.NullAaActivationCondition(autoAbility.ActivationCondition),
				Counter:             h.NullCounterType(autoAbility.Counter),
			})
			if err != nil {
				return h.NewErr(autoAbility.Error(), err, "couldn't create auto-ability")
			}

			autoAbility.ID = dbAutoAbility.ID
			l.AutoAbilities[autoAbility.Name] = autoAbility
			l.AutoAbilitiesID[autoAbility.ID] = autoAbility
		}
		return nil
	})
}

func (l *Lookup) seedAutoAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/auto_abilities.json"

	var autoAbilities []AutoAbility
	err := loadJSONFile(string(srcPath), &autoAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAutoAbility := range autoAbilities {
			autoAbility, err := GetResource(jsonAutoAbility.Name, l.AutoAbilities)
			if err != nil {
				return err
			}

			autoAbility, err = l.assignAutoAbilityFKs(qtx, autoAbility)
			if err != nil {
				return h.NewErr(autoAbility.Error(), err)
			}

			err = qtx.UpdateAutoAbility(context.Background(), database.UpdateAutoAbilityParams{
				DataHash:          generateDataHash(autoAbility),
				GradRcvryStatID:   h.GetNullInt32(autoAbility.GradRecoveryStatID),
				OnHitElementID:    h.GetNullInt32(autoAbility.OnHitElementID),
				AddedElemResistID: h.ObjPtrToNullInt32ID(autoAbility.AddedElemResist),
				OnHitStatusID:     h.ObjPtrToNullInt32ID(autoAbility.OnHitStatus),
				AddedPropertyID:   h.GetNullInt32(autoAbility.AddedPropertyID),
				CnvrsnFromModID:   h.GetNullInt32(autoAbility.CnvrsnFromModID),
				CnvrsnToModID:     h.GetNullInt32(autoAbility.CnvrsnToModID),
				ID:                autoAbility.ID,
			})
			if err != nil {
				return h.NewErr(autoAbility.Error(), err, "couldn't update auto-ability")
			}

			err = l.seedAutoAbilityJunctions(qtx, autoAbility)
			if err != nil {
				return h.NewErr(autoAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) assignAutoAbilityFKs(qtx *database.Queries, autoAbility AutoAbility) (AutoAbility, error) {
	var err error

	autoAbility.GradRecoveryStatID, err = assignFKPtr(autoAbility.GradualRecovery, l.Stats)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.OnHitElementID, err = assignFKPtr(autoAbility.OnHitElement, l.Elements)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.AddedPropertyID, err = assignFKPtr(autoAbility.AddedProperty, l.Properties)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.CnvrsnFromModID, err = assignFKPtr(autoAbility.ConversionFrom, l.Modifiers)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.CnvrsnToModID, err = assignFKPtr(autoAbility.ConversionTo, l.Modifiers)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.RequiredItem, err = seedObjPtrAssignFK(qtx, autoAbility.RequiredItem, l.seedItemAmount)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.AddedElemResist, err = seedObjPtrAssignFK(qtx, autoAbility.AddedElemResist, l.seedElementalResist)
	if err != nil {
		return AutoAbility{}, err
	}

	autoAbility.OnHitStatus, err = seedObjPtrAssignFK(qtx, autoAbility.OnHitStatus, l.seedInflictedStatus)
	if err != nil {
		return AutoAbility{}, err
	}

	return autoAbility, nil
}

func (l *Lookup) seedAutoAbilityJunctions(qtx *database.Queries, autoAbility AutoAbility) error {
	functions := []func(*database.Queries, AutoAbility) error{
		l.seedAutoAbilityRelatedStats,
		l.seedAutoAbilityLockedOutAbilities,
		l.seedAutoAbilityAutoItemUse,
		l.seedAutoAbilityAddedStatusResists,
		l.seedAutoAbilityAddedStatusses,
		l.seedAutoAbilityStatChanges,
		l.seedAutoAbilityModifierChanges,
	}

	for _, function := range functions {
		err := function(qtx, autoAbility)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedAutoAbilityRelatedStats(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, jsonStat := range autoAbility.RelatedStats {
		junction, err := createJunction(autoAbility, jsonStat, l.Stats)
		if err != nil {
			return err
		}

		err = qtx.CreateAutoAbilitiesRelatedStatsJunction(context.Background(), database.CreateAutoAbilitiesRelatedStatsJunctionParams{
			DataHash:      generateDataHash(junction),
			AutoAbilityID: junction.ParentID,
			StatID:        junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonStat, err, "couldn't junction related stat")
		}
	}

	return nil
}

func (l *Lookup) seedAutoAbilityLockedOutAbilities(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, jsonAbility := range autoAbility.LockedOutAbilities {
		junction, err := createJunction(autoAbility, jsonAbility, l.AutoAbilities)
		if err != nil {
			return err
		}

		err = qtx.CreateAutoAbilitiesLockedOutJunction(context.Background(), database.CreateAutoAbilitiesLockedOutJunctionParams{
			DataHash:        generateDataHash(junction),
			ParentAbilityID: junction.ParentID,
			ChildAbilityID:  junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonAbility, err, "couldn't junction locked out ability")
		}
	}

	return nil
}

func (l *Lookup) seedAutoAbilityAutoItemUse(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, jsonItem := range autoAbility.AutoItemUse {
		junction, err := createJunction(autoAbility, jsonItem, l.Items)
		if err != nil {
			return err
		}

		err = qtx.CreateAutoAbilitiesRequiredItemJunction(context.Background(), database.CreateAutoAbilitiesRequiredItemJunctionParams{
			DataHash:      generateDataHash(junction),
			AutoAbilityID: junction.ParentID,
			ItemID:        junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonItem, err, "couldn't junction auto item use item")
		}
	}

	return nil
}

func (l *Lookup) seedAutoAbilityAddedStatusses(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, jsonStatus := range autoAbility.AddedStatusses {
		junction, err := createJunction(autoAbility, jsonStatus, l.StatusConditions)
		if err != nil {
			return err
		}

		err = qtx.CreateAutoAbilitiesAddedStatussesJunction(context.Background(), database.CreateAutoAbilitiesAddedStatussesJunctionParams{
			DataHash:          generateDataHash(junction),
			AutoAbilityID:     junction.ParentID,
			StatusConditionID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonStatus, err, "couldn't junction added status")
		}
	}

	return nil
}

func (l *Lookup) seedAutoAbilityAddedStatusResists(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, statusResist := range autoAbility.AddedStatusResists {
		junction, err := createJunctionSeed(qtx, autoAbility, statusResist, l.seedStatusResist)
		if err != nil {
			return err
		}

		err = qtx.CreateAutoAbilitiesAddedStatusResistsJunction(context.Background(), database.CreateAutoAbilitiesAddedStatusResistsJunctionParams{
			DataHash:       generateDataHash(junction),
			AutoAbilityID:  junction.ParentID,
			StatusResistID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(statusResist.Error(), err, "couldn't junction status resist")
		}
	}

	return nil
}

func (l *Lookup) seedAutoAbilityStatChanges(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, statChange := range autoAbility.StatChanges {
		junction, err := createJunctionSeed(qtx, autoAbility, statChange, l.seedStatChange)
		if err != nil {
			return err
		}

		err = qtx.CreateAutoAbilitiesStatChangesJunction(context.Background(), database.CreateAutoAbilitiesStatChangesJunctionParams{
			DataHash:      generateDataHash(junction),
			AutoAbilityID: junction.ParentID,
			StatChangeID:  junction.ChildID,
		})
		if err != nil {
			return h.NewErr(statChange.Error(), err, "couldn't junction stat change")
		}
	}

	return nil
}

func (l *Lookup) seedAutoAbilityModifierChanges(qtx *database.Queries, autoAbility AutoAbility) error {
	for _, modifierChange := range autoAbility.ModifierChanges {
		junction, err := createJunctionSeed(qtx, autoAbility, modifierChange, l.seedModifierChange)
		if err != nil {
			return err
		}

		err = qtx.CreateAutoAbilitiesModifierChangesJunction(context.Background(), database.CreateAutoAbilitiesModifierChangesJunctionParams{
			DataHash:         generateDataHash(junction),
			AutoAbilityID:    junction.ParentID,
			ModifierChangeID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(modifierChange.Error(), err, "couldn't junction modifier change")
		}
	}

	return nil
}
