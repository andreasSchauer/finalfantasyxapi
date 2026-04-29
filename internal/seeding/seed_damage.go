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

func (ad AbilityDamage) Error() string {
	return fmt.Sprintf("ability damage with attack type: %s, target stat: %s, damage type: %s, formula: %s, damage constant %d, condition: %v", ad.AttackType, ad.TargetStat, ad.DamageType, ad.DamageFormula, ad.DamageConstant, h.PtrToString(ad.Condition))
}

func (l *Lookup) seedDamage(qtx *database.Queries, damage Damage) (Damage, error) {
	var err error

	damage.ElementID, err = assignFKPtr(damage.Element, l.Elements)
	if err != nil {
		return Damage{}, h.NewErr(damage.Error(), err)
	}

	dbDamage, err := qtx.CreateDamage(context.Background(), database.CreateDamageParams{
		DataHash:        generateDataHash(damage),
		Critical:        database.ToNullCriticalType(damage.Critical),
		CriticalPlusVal: h.GetNullInt32(damage.CriticalPlusVal),
		IsPiercing:      damage.IsPiercing,
		BreakDmgLimit:   database.ToNullBreakDmgLmtType(damage.BreakDmgLimit),
		ElementID:       h.GetNullInt32(damage.ElementID),
	})
	if err != nil {
		return Damage{}, h.NewErr(damage.Error(), err, "couldn't create damage")
	}

	damage.ID = dbDamage.ID

	err = l.seedAbilityDamages(qtx, damage)
	if err != nil {
		return Damage{}, h.NewErr(damage.Error(), err)
	}

	return damage, nil
}

func (l *Lookup) seedAbilityDamages(qtx *database.Queries, damage Damage) error {
	ability := l.currentAbility
	battleInteraction := l.currentBI

	for _, abilityDamage := range damage.DamageCalc {
		fourWay, err := createFourWayJunctionSeed(qtx, ability, battleInteraction, damage, abilityDamage, l.seedAbilityDamage)
		if err != nil {
			return err
		}

		err = qtx.CreateDamagesDamageCalcJunction(context.Background(), database.CreateDamagesDamageCalcJunctionParams{
			DataHash:            generateDataHash(fourWay),
			AbilityID:           fourWay.GreatGrandparentID,
			BattleInteractionID: fourWay.GrandparentID,
			DamageID:            fourWay.ParentID,
			AbilityDamageID:     fourWay.ChildID,
		})
		if err != nil {
			return h.NewErr(abilityDamage.Error(), err, "couldn't junction ability damage")
		}
	}

	return nil
}

func (l *Lookup) seedAbilityDamage(qtx *database.Queries, abilityDamage AbilityDamage) (AbilityDamage, error) {
	var err error

	abilityDamage.StatID, err = assignFK(abilityDamage.TargetStat, l.Stats)
	if err != nil {
		return AbilityDamage{}, h.NewErr(abilityDamage.Error(), err)
	}

	dbAbilityDamage, err := qtx.CreateAbilityDamage(context.Background(), database.CreateAbilityDamageParams{
		DataHash:       generateDataHash(abilityDamage),
		Condition:      h.GetNullString(abilityDamage.Condition),
		AttackType:     database.AttackType(abilityDamage.AttackType),
		StatID:         abilityDamage.StatID,
		DamageType:     database.DamageType(abilityDamage.DamageType),
		DamageFormula:  database.DamageFormula(abilityDamage.DamageFormula),
		DamageConstant: abilityDamage.DamageConstant,
	})
	if err != nil {
		return AbilityDamage{}, h.NewErr(abilityDamage.Error(), err, "couldn't create ability damage")
	}

	abilityDamage.ID = dbAbilityDamage.ID

	return abilityDamage, nil
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
