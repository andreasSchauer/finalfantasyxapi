package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getMonsterMonsterAbilities(m Monster) ([]MonsterAbility, error) {
	return m.Abilities, nil
}

func (l *Lookup) seedJuncMonstersMonsterAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + monster abilities"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterMonsterAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersAbilitiesJunctionBulk(ctx, database.CreateMonstersAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		MonsterID:        jParams.ParentIDs,
		MonsterAbilityID: jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterAutoAbilities(m Monster) ([]AutoAbility, error) {
	return getResources(m.AutoAbilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncMonstersAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + auto-abilities"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersAutoAbilitiesJunctionBulk(ctx, database.CreateMonstersAutoAbilitiesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		MonsterID:     jParams.ParentIDs,
		AutoAbilityID: jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterBaseStats(m Monster) ([]BaseStat, error) {
	return m.BaseStats, nil
}

func (l *Lookup) seedJuncMonstersBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + base stats"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersBaseStatsJunctionBulk(ctx, database.CreateMonstersBaseStatsJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		MonsterID:  jParams.ParentIDs,
		BaseStatID: jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterElementalResists(m Monster) ([]ElementalResist, error) {
	return m.ElemResists, nil
}

func (l *Lookup) seedJuncMonstersElementalResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + elemental resists"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterElementalResists)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersElemResistsJunctionBulk(ctx, database.CreateMonstersElemResistsJunctionBulkParams{
		DataHash:     jParams.DataHashes,
		MonsterID:    jParams.ParentIDs,
		ElemResistID: jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterStatusImmunities(m Monster) ([]StatusCondition, error) {
	return getResources(m.StatusImmunities, l.StatusConditions)
}

func (l *Lookup) seedJuncMonstersStatusImmunities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + status immunities"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterStatusImmunities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersImmunitiesJunctionBulk(ctx, database.CreateMonstersImmunitiesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		MonsterID:         jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterProperties(m Monster) ([]Property, error) {
	return getResources(m.Properties, l.Properties)
}

func (l *Lookup) seedJuncMonstersProperties(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + properties"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterProperties)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersPropertiesJunctionBulk(ctx, database.CreateMonstersPropertiesJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		MonsterID:  jParams.ParentIDs,
		PropertyID: jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterRonsoRages(m Monster) ([]RonsoRage, error) {
	return getResources(m.RonsoRages, l.RonsoRages)
}

func (l *Lookup) seedJuncMonstersRonsoRages(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + ronso rages"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterRonsoRages)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersRonsoRagesJunctionBulk(ctx, database.CreateMonstersRonsoRagesJunctionBulkParams{
		DataHash:    jParams.DataHashes,
		MonsterID:   jParams.ParentIDs,
		RonsoRageID: jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterStatusResists(m Monster) ([]StatusResist, error) {
	return m.StatusResists, nil
}

func (l *Lookup) seedJuncMonstersStatusResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + status resists"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterStatusResists)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersStatusResistsJunctionBulk(ctx, database.CreateMonstersStatusResistsJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		MonsterID:      jParams.ParentIDs,
		StatusResistID: jParams.ChildIDs,
	})
}
