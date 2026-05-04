package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Damage struct {
	ID              int32
	DamageCalc      []AbilityDamage `json:"damage_calc"`
	Critical        *string         `json:"critical"`
	CriticalPlusVal *int32          `json:"critical_plus_val"`
	IsPiercing      bool            `json:"is_piercing"`
	BreakDmgLimit   *string         `json:"break_dmg_lmt"`
	ElementID       *int32
	Element         *string `json:"element"`
}

func (d Damage) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", d),
		h.DerefOrNil(d.Critical),
		h.DerefOrNil(d.CriticalPlusVal),
		d.IsPiercing,
		h.DerefOrNil(d.BreakDmgLimit),
		h.DerefOrNil(d.ElementID),
	}
}

func (d Damage) GetID() int32 {
	return d.ID
}

func (d *Damage) SetID(id int32) {
	d.ID = id
}

func (d Damage) Error() string {
	return fmt.Sprintf("damage with critical: %v, crit plus: %v, piercing: %t, bdl: %v, element: %v", h.PtrToString(d.Critical), h.PtrToString(d.CriticalPlusVal), d.IsPiercing, h.PtrToString(d.BreakDmgLimit), h.PtrToString(d.Element))
}

type AbilityDamage struct {
	ID             int32
	Condition      *string `json:"condition"`
	AttackType     string  `json:"attack_type"`
	StatID         int32
	TargetStat     string `json:"target_stat"`
	DamageType     string `json:"damage_type"`
	DamageFormula  string `json:"damage_formula"`
	DamageConstant int32  `json:"damage_constant"`
}

func (ad AbilityDamage) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", ad),
		h.DerefOrNil(ad.Condition),
		ad.AttackType,
		ad.StatID,
		ad.DamageType,
		ad.DamageFormula,
		ad.DamageConstant,
	}
}

func (ad AbilityDamage) GetID() int32 {
	return ad.ID
}

func (ad *AbilityDamage) SetID(id int32) {
	ad.ID = id
}

func (ad AbilityDamage) Error() string {
	return fmt.Sprintf("ability damage with attack type: %s, target stat: %s, damage type: %s, formula: %s, damage constant %d, condition: %v", ad.AttackType, ad.TargetStat, ad.DamageType, ad.DamageFormula, ad.DamageConstant, h.PtrToString(ad.Condition))
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

func (l *Lookup) completeDamage(damage *Damage) error {
	if damage == nil {
		return nil
	}

	err := l.assignID(damage)
	if err != nil {
		return err
	}

	err = assignIDs(l, damage.DamageCalc)
	if err != nil {
		return err
	}

	return nil
}


func (l *Lookup) loop5SeedAbilityDamages(qtx *database.Queries, ctx context.Context) error {
	damages, err := l.extractAbilityDamages()
	if err != nil {
		return err
	}

	params := database.CreateAbilityDamageBulkParams{
		DataHash:       make([]string, len(damages)),
		Condition: 		make([]sql.NullString, len(damages)),
		AttackType: 	make([]database.AttackType, len(damages)),
		StatID: 		make([]int32, len(damages)),
		DamageType: 	make([]database.DamageType, len(damages)),
		DamageFormula: 	make([]database.DamageFormula, len(damages)),
		DamageConstant: make([]int32, len(damages)),
	}

	for i, d := range damages {
		params.DataHash[i] = generateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.AttackType[i] = database.AttackType(d.AttackType)
		params.StatID[i] = d.StatID
		params.DamageType[i] = database.DamageType(d.DamageType)
		params.DamageFormula[i] = database.DamageFormula(d.DamageFormula)
		params.DamageConstant[i] = d.DamageConstant
	}

	dbRows, err := qtx.CreateAbilityDamageBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability damages: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAbilityDamages() ([]AbilityDamage, error) {
	damages := []AbilityDamage{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		newDamages, err := l.prepareAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}
		damages = append(damages, newDamages...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newDamages, err := l.prepareAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newDamages, err := l.prepareAbilityDamages(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newDamages, err := l.prepareAbilityDamages(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		newDamages, err := l.prepareAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newDamages, err := l.prepareAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return dedupeRows(damages, l.Hashes), nil
}


func (l *Lookup) prepareAbilityDamages(battleInteractions []BattleInteraction) ([]AbilityDamage, error) {
	damages := []AbilityDamage{}

	for j := range battleInteractions {
		bi := &battleInteractions[j]

		if bi.Damage == nil {
			continue
		}

		newDamages, err := l.prepareAbilityDamage(bi.Damage.DamageCalc)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return damages, nil
}

func (l *Lookup) prepareAbilityDamage(abilityDamages []AbilityDamage) ([]AbilityDamage, error) {
	damages := []AbilityDamage{}
	var err error

	for i := range abilityDamages {
		ad := &abilityDamages[i]

		ad.StatID, err = assignFK(ad.TargetStat, l.Stats)
		if err != nil {
			return nil, err
		}

		damages = append(damages, *ad)
	}

	return damages, nil
}

func (l *Lookup) seedJuncDamagesDamageCalc(qtx *database.Queries, ctx context.Context) error {
	const desc string = "damages + damage calc"
	params := database.CreateDamagesDamageCalcJunctionBulkParams{
		DataHash: 				make([]string, 0),
		AbilityID: 				make([]int32, 0),
		BattleInteractionID: 	make([]int32, 0),
		DamageID: 				make([]int32, 0),
		AbilityDamageID: 		make([]int32, 0),
	}

	for _, ability := range l.getAbilities() {
		bis, err := l.getAbilityBattleInteractions(ability)
		if err != nil {
			return err
		}
		
		for _, bi := range bis {
			if bi.Damage == nil {
				continue
			}

			for _, ad := range bi.Damage.DamageCalc {
				j := FourWayJunction{}
				j.GreatGrandparentID = ability.ID
				j.GrandparentID = bi.ID
				j.ParentID = bi.Damage.ID
				j.ChildID = ad.ID
				dataHash := generateJunctionHash(j, desc)

				params.DataHash = append(params.DataHash, dataHash)
				params.AbilityID = append(params.AbilityID, ability.ID)
				params.BattleInteractionID = append(params.BattleInteractionID, bi.ID)
				params.DamageID = append(params.DamageID, bi.Damage.ID)
				params.AbilityDamageID = append(params.AbilityDamageID, ad.ID)
			}
		}
	}

	return qtx.CreateDamagesDamageCalcJunctionBulk(ctx, params)
}