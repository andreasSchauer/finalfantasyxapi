package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type BattleInteraction struct {
	ID                        int32
	Target                    string            `json:"target"`
	BasedOnUserAttack         bool              `json:"based_on_user_attack"`
	Range                     *int32            `json:"range"`
	Damage                    *Damage           `json:"damage"`
	ShatterRate               int32             `json:"shatter_rate"`
	Accuracy                  Accuracy          `json:"accuracy"`
	AffectedBy                []string          `json:"affected_by"`
	HitAmount                 int32             `json:"hit_amount"`
	SpecialAction             *string           `json:"special_action"`
	InflictedDelay            *InflictedDelay   `json:"inflicted_delay"`
	InflictedStatusConditions []InflictedStatus `json:"inflicted_status_conditions"`
	RemovedStatusConditions   []string          `json:"removed_status_conditions"`
	CopiedStatusConditions    []InflictedStatus `json:"copied_status_conditions"`
	StatChanges               []StatChange      `json:"stat_changes"`
	ModifierChanges           []ModifierChange  `json:"modifier_changes"`
}

func (bi BattleInteraction) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", bi),
		bi.Target,
		bi.BasedOnUserAttack,
		h.DerefOrNil(bi.Range),
		h.ObjPtrToID(bi.Damage),
		bi.ShatterRate,
		bi.Accuracy.ID,
		bi.HitAmount,
		h.DerefOrNil(bi.SpecialAction),
		h.ObjPtrToID(bi.InflictedDelay),
	}
}

func (bi BattleInteraction) GetID() int32 {
	return bi.ID
}

func (bi BattleInteraction) Error() string {
	return fmt.Sprintf("battle interaction with target: %s, phys attack: %t, range: %v, damage id: %v, shatter rate: %v, accuracy id: %d, hit amount: %d, special action: %v", bi.Target, bi.BasedOnUserAttack, h.PtrToString(bi.Range), h.ObjPtrToID(bi.Damage), bi.ShatterRate, bi.Accuracy.ID, bi.HitAmount, h.PtrToString(bi.SpecialAction))
}

func (l *Lookup) seedBattleInteractions(qtx *database.Queries, ability Ability, battleInteractions []BattleInteraction) error {
	for _, battleInteraction := range battleInteractions {
		junction, err := createJunctionSeed(qtx, ability, battleInteraction, l.seedBattleInteraction)
		if err != nil {
			return err
		}
		battleInteraction.ID = junction.ChildID

		err = qtx.CreateAbilitiesBattleInteractionsJunction(context.Background(), database.CreateAbilitiesBattleInteractionsJunctionParams{
			DataHash:            generateDataHash(junction),
			AbilityID:           junction.ParentID,
			BattleInteractionID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(battleInteraction.Error(), err, "couldn't junction battle interaction")
		}

		l.currentBI = battleInteraction

		err = l.seedBattleInteractionRelationships(qtx, ability, battleInteraction)
		if err != nil {
			return h.NewErr(battleInteraction.Error(), err)
		}
	}

	return nil
}

func (l *Lookup) seedBattleInteraction(qtx *database.Queries, battleInteraction BattleInteraction) (BattleInteraction, error) {
	var err error

	battleInteraction.Accuracy, err = seedObjAssignID(qtx, battleInteraction.Accuracy, l.seedAccuracy)
	if err != nil {
		return BattleInteraction{}, h.NewErr(battleInteraction.Error(), err)
	}

	battleInteraction.InflictedDelay, err = seedObjPtrAssignFK(qtx, battleInteraction.InflictedDelay, l.seedInflictedDelay)
	if err != nil {
		return BattleInteraction{}, h.NewErr(battleInteraction.Error(), err)
	}

	dbBattleInteraction, err := qtx.CreateBattleInteraction(context.Background(), database.CreateBattleInteractionParams{
		DataHash:          generateDataHash(battleInteraction),
		Target:            database.TargetType(battleInteraction.Target),
		BasedOnUserAttack: battleInteraction.BasedOnUserAttack,
		Range:             h.GetNullInt32(battleInteraction.Range),
		ShatterRate:       battleInteraction.ShatterRate,
		AccuracyID:        battleInteraction.Accuracy.ID,
		InflictedDelayID:  h.ObjPtrToNullInt32ID(battleInteraction.InflictedDelay),
		HitAmount:         battleInteraction.HitAmount,
		SpecialAction:     database.ToNullSpecialActionType(battleInteraction.SpecialAction),
	})
	if err != nil {
		return BattleInteraction{}, h.NewErr(battleInteraction.Error(), err, "couldn't create battle interaction")
	}

	battleInteraction.ID = dbBattleInteraction.ID

	return battleInteraction, nil
}

func (l *Lookup) seedBattleInteractionRelationships(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	if battleInteraction.Damage != nil {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, *battleInteraction.Damage, l.seedDamage)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntDamageJunction(context.Background(), database.CreateBattleIntDamageJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			DamageID:            threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(battleInteraction.Damage.Error(), err, "couldn't junction damage")
		}
	}

	functions := []func(*database.Queries, Ability, BattleInteraction) error{
		l.seedBattleIntAffectedBy,
		l.seedBattleIntInflictedConditions,
		l.seedBattleIntRemovedConditions,
		l.seedBattleIntCopiedConditions,
		l.seedBattleIntStatChanges,
		l.seedBattleIntModifierChanges,
	}

	for _, function := range functions {
		err := function(qtx, ability, battleInteraction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntAffectedBy(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, conditionString := range battleInteraction.AffectedBy {
		threeWay, err := createThreeWayJunction(ability, battleInteraction, conditionString, l.StatusConditions)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntAffectedByJunction(context.Background(), database.CreateBattleIntAffectedByJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			StatusConditionID:   threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(conditionString, err, "couldn't junction affected by status")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntInflictedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.InflictedStatusConditions {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, condition, l.seedInflictedStatus)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntInflictedConditionsJunction(context.Background(), database.CreateBattleIntInflictedConditionsJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			InflictedStatusID:   threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(condition.Error(), err, "couldn't junction inflicted status condition")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntRemovedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, conditionString := range battleInteraction.RemovedStatusConditions {
		threeWay, err := createThreeWayJunction(ability, battleInteraction, conditionString, l.StatusConditions)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntRemovedConditionsJunction(context.Background(), database.CreateBattleIntRemovedConditionsJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			StatusConditionID:   threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(conditionString, err, "couldn't junction removed status condition")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntCopiedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.CopiedStatusConditions {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, condition, l.seedInflictedStatus)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntCopiedConditionsJunction(context.Background(), database.CreateBattleIntCopiedConditionsJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			InflictedStatusID:   threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(condition.Error(), err, "couldn't junction copied status condition")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntStatChanges(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, statChange := range battleInteraction.StatChanges {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, statChange, l.seedStatChange)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntStatChangesJunction(context.Background(), database.CreateBattleIntStatChangesJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			StatChangeID:        threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(statChange.Error(), err, "couldn't junction stat change")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntModifierChanges(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, modifierChange := range battleInteraction.ModifierChanges {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, modifierChange, l.seedModifierChange)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntModifierChangesJunction(context.Background(), database.CreateBattleIntModifierChangesJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			ModifierChangeID:    threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(modifierChange.Error(), err, "couldn't junction modifier change")
		}
	}

	return nil
}

func (l *Lookup) loop3SeedBattleInteractions(qtx *database.Queries, ctx context.Context) error {
	bis, err := l.extractBattleInteractions()
	if err != nil {
		return err
	}

	params := database.CreateBattleInteractionBulkParams{
		DataHash:       	make([]string, len(bis)),
		Target: 			make([]database.TargetType, len(bis)),
		BasedOnUserAttack: 	make([]bool, len(bis)),
		Range: 				make([]sql.NullInt32, len(bis)),
		ShatterRate: 		make([]int32, len(bis)),
		AccuracyID: 		make([]int32, len(bis)),
		InflictedDelayID: 	make([]sql.NullInt32, len(bis)),
		HitAmount: 			make([]int32, len(bis)),
		SpecialAction: 		make([]database.NullSpecialActionType, len(bis)),
	}

	for i, bi := range bis {
		params.DataHash[i] = generateDataHash(bi)
		params.Target[i] = database.TargetType(bi.Target)
		params.BasedOnUserAttack[i] = bi.BasedOnUserAttack
		params.Range[i] = h.GetNullInt32(bi.Range)
		params.ShatterRate[i] = bi.ShatterRate
		params.AccuracyID[i] = bi.Accuracy.ID
		params.InflictedDelayID[i] = h.ObjPtrToNullInt32ID(bi.InflictedDelay)
		params.HitAmount[i] = bi.HitAmount
		params.SpecialAction[i] = database.ToNullSpecialActionType(bi.SpecialAction)
	}

	dbRows, err := qtx.CreateBattleInteractionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create battle interactions: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractBattleInteractions() ([]BattleInteraction, error) {
	bis := []BattleInteraction{}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		bisNew, err := l.prepareBattleInteractions(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		bisNew, err := l.prepareBattleInteractions(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	return dedupeRows(bis, l.Hashes), nil
}


func (l *Lookup) prepareBattleInteractions(bis []BattleInteraction) ([]BattleInteraction, error) {
	bisNew := []BattleInteraction{}
	var err error

	for j := range bis {
		bi := &bis[j]

		bi.Accuracy.ID, err = l.getHashID(bi.Accuracy)
		if err != nil {
			return nil, err
		}

		if bi.InflictedDelay != nil {
			bi.InflictedDelay.ID, err = l.getHashID(bi.InflictedDelay)
		}

		bisNew = append(bisNew, *bi)
	}

	return bisNew, nil
}


func (l *Lookup) loop2SeedDamages(qtx *database.Queries, ctx context.Context) error {
	damages, err := l.extractDamages()
	if err != nil {
		return err
	}

	params := database.CreateDamageBulkParams{
		DataHash:        make([]string, len(damages)),
		Critical:        make([]database.NullCriticalType, len(damages)),
		CriticalPlusVal: make([]sql.NullInt32, len(damages)),
		IsPiercing:      make([]bool, len(damages)),
		BreakDmgLimit:   make([]database.NullBreakDmgLmtType, len(damages)),
		ElementID:       make([]sql.NullInt32, len(damages)),
	}

	for i, d := range damages {
		params.DataHash[i] = generateDataHash(d)
		params.Critical[i] = database.ToNullCriticalType(d.Critical)
		params.CriticalPlusVal[i] = h.GetNullInt32(d.CriticalPlusVal)
		params.IsPiercing[i] = d.IsPiercing
		params.BreakDmgLimit[i] = database.ToNullBreakDmgLmtType(d.BreakDmgLimit)
		params.ElementID[i] = h.GetNullInt32(d.ElementID)
	}

	dbRows, err := qtx.CreateDamageBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create damages: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractDamages() ([]Damage, error) {
	damages := []Damage{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		newDamages, err := l.prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newDamages, err := l.prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newDamages, err := l.prepareDamages(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newDamages, err := l.prepareDamages(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		newDamages, err := l.prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newDamages, err := l.prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return dedupeRows(damages, l.Hashes), nil
}

func (l *Lookup) prepareDamages(battleInteractions []BattleInteraction) ([]Damage, error) {
	damages := []Damage{}
	var err error

	for i := range battleInteractions {
		bi := &battleInteractions[i]

		if bi.Damage != nil {
			bi.Damage.ElementID, err = assignFKPtr(bi.Damage.Element, l.Elements)
			if err != nil {
				return nil, err
			}

			damages = append(damages, *bi.Damage)
		}
	}

	return damages, nil
}


func (l *Lookup) loop1SeedAccuracies(qtx *database.Queries, ctx context.Context) error {
	accuracies := l.extractAccuracies()

	params := database.CreateAbilityAccuracyBulkParams{
		DataHash:    make([]string, len(accuracies)),
		AccSource:   make([]database.AccSourceType, len(accuracies)),
		HitChance:   make([]sql.NullInt32, len(accuracies)),
		AccModifier: make([]sql.NullFloat64, len(accuracies)),
	}

	for i, a := range accuracies {
		params.DataHash[i] = generateDataHash(a)
		params.AccSource[i] = database.AccSourceType(a.AccSource)
		params.HitChance[i] = h.GetNullInt32(a.HitChance)
		params.AccModifier[i] = h.GetNullFloat64(a.AccModifier)
	}

	dbRows, err := qtx.CreateAbilityAccuracyBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability accuracies: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAccuracies() []Accuracy {
	accuracies := []Accuracy{}

	for _, ability := range l.json.enemyAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.items {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.overdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.playerAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.triggerCommands {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.unspecifiedAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, aeon := range l.json.aeons {
		if aeon.PhysAtkAccuracy != nil {
			accuracies = append(accuracies, *aeon.PhysAtkAccuracy)
		}
	}

	return dedupeRows(accuracies, l.Hashes)
}


func (l *Lookup) loop1SeedInflictedDelays(qtx *database.Queries, ctx context.Context) error {
	delays := l.extractInflictedDelays()

	params := database.CreateInflictedDelayBulkParams{
		DataHash:       make([]string, len(delays)),
		Condition:      make([]sql.NullString, len(delays)),
		CtbAttackType:  make([]database.CtbAttackType, len(delays)),
		DelayType:      make([]database.DelayType, len(delays)),
		DamageConstant: make([]int32, len(delays)),
	}

	for i, d := range delays {
		params.DataHash[i] = generateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.CtbAttackType[i] = database.CtbAttackType(d.CTBAttackType)
		params.DelayType[i] = database.DelayType(d.DelayType)
		params.DamageConstant[i] = d.DamageConstant
	}

	dbRows, err := qtx.CreateInflictedDelayBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create inflicted delays: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractInflictedDelays() []InflictedDelay {
	delays := []InflictedDelay{}

	for _, ability := range l.json.enemyAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.items {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.overdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.playerAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.triggerCommands {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.unspecifiedAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, status := range l.json.statusConditions {
		if status.CtbOnInfliction != nil {
			delays = append(delays, *status.CtbOnInfliction)
		}
	}

	return dedupeRows(delays, l.Hashes)
}

func (l *Lookup) loop2SeedModifierChanges(qtx *database.Queries, ctx context.Context) error {
	changes, err := l.extractModifierChanges()
	if err != nil {
		return err
	}

	params := database.CreateModifierChangeBulkParams{
		DataHash:        make([]string, len(changes)),
		ModifierID:      make([]int32, len(changes)),
		CalculationType: make([]database.CalculationType, len(changes)),
		Value:           make([]float32, len(changes)),
	}

	for i, c := range changes {
		params.DataHash[i] = generateDataHash(c)
		params.ModifierID[i] = c.ModifierID
		params.CalculationType[i] = database.CalculationType(c.CalculationType)
		params.Value[i] = c.Value
	}

	dbRows, err := qtx.CreateModifierChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create modifier changes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractModifierChanges() ([]ModifierChange, error) {
	changes := []ModifierChange{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		newChanges, err := l.prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newChanges, err := l.prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newChanges, err := l.prepareAbilityModifierChanges(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newChanges, err := l.prepareAbilityModifierChanges(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		newChanges, err := l.prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newChanges, err := l.prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		newChanges, err := l.prepareModifierChanges(autoAbility.ModifierChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.properties {
		property := &l.json.properties[i]

		newChanges, err := l.prepareModifierChanges(property.ModifierChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.statusConditions {
		status := &l.json.statusConditions[i]

		newChanges, err := l.prepareModifierChanges(status.ModifierChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	return dedupeRows(changes, l.Hashes), nil
}

func (l *Lookup) prepareAbilityModifierChanges(battleInteractions []BattleInteraction) ([]ModifierChange, error) {
	changes := []ModifierChange{}

	for i := range battleInteractions {
		bi := &battleInteractions[i]

		changesNew, err := l.prepareModifierChanges(bi.ModifierChanges)
		if err != nil {
			return nil, err
		}
		changes = append(changes, changesNew...)
	}

	return changes, nil
}

func (l *Lookup) prepareModifierChanges(changes []ModifierChange) ([]ModifierChange, error) {
	changesNew := []ModifierChange{}
	var err error

	for i := range changes {
		change := &changes[i]

		change.ModifierID, err = assignFK(change.ModifierName, l.Modifiers)
		if err != nil {
			return nil, err
		}

		changesNew = append(changesNew, *change)
	}

	return changesNew, nil
}
