package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getMonsterMonsterAbilities(m seedtypes.Monster) ([]seedtypes.MonsterAbility, error) {
	return m.Abilities, nil
}

func MonstersMonsterAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + monster abilities"
	jParams, err := processJunctions(desc, sh.RawJSON.Monsters, getMonsterMonsterAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersAbilitiesJunctionBulk(ctx, database.CreateMonstersAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		MonsterID:        jParams.ParentIDs,
		MonsterAbilityID: jParams.ChildIDs,
	})
}

func getMonsterAutoAbilities(m seedtypes.Monster) ([]seedtypes.AutoAbility, error) {
	return lookup.GetResources(m.AutoAbilities, lookup.L.AutoAbilities)
}

func MonstersAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + auto-abilities"
	jParams, err := processJunctions(desc, sh.RawJSON.Monsters, getMonsterAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersAutoAbilitiesJunctionBulk(ctx, database.CreateMonstersAutoAbilitiesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		MonsterID:     jParams.ParentIDs,
		AutoAbilityID: jParams.ChildIDs,
	})
}

func getMonsterBaseStats(m seedtypes.Monster) ([]seedtypes.BaseStat, error) {
	return m.BaseStats, nil
}

func MonstersBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + base stats"
	jParams, err := processJunctions(desc, sh.RawJSON.Monsters, getMonsterBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersBaseStatsJunctionBulk(ctx, database.CreateMonstersBaseStatsJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		MonsterID:  jParams.ParentIDs,
		BaseStatID: jParams.ChildIDs,
	})
}

func getMonsterElementalResists(m seedtypes.Monster) ([]seedtypes.ElementalResist, error) {
	return m.ElemResists, nil
}

func MonstersElementalResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + elemental resists"
	jParams, err := processJunctions(desc, sh.RawJSON.Monsters, getMonsterElementalResists)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersElemResistsJunctionBulk(ctx, database.CreateMonstersElemResistsJunctionBulkParams{
		DataHash:     jParams.DataHashes,
		MonsterID:    jParams.ParentIDs,
		ElemResistID: jParams.ChildIDs,
	})
}

func getMonsterStatusImmunities(m seedtypes.Monster) ([]seedtypes.StatusCondition, error) {
	return lookup.GetResources(m.StatusImmunities, lookup.L.StatusConditions)
}

func MonstersStatusImmunities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + status immunities"
	jParams, err := processJunctions(desc, sh.RawJSON.Monsters, getMonsterStatusImmunities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersImmunitiesJunctionBulk(ctx, database.CreateMonstersImmunitiesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		MonsterID:         jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}

func getMonsterProperties(m seedtypes.Monster) ([]seedtypes.Property, error) {
	return lookup.GetResources(m.Properties, lookup.L.Properties)
}

func MonstersProperties(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + properties"
	jParams, err := processJunctions(desc, sh.RawJSON.Monsters, getMonsterProperties)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersPropertiesJunctionBulk(ctx, database.CreateMonstersPropertiesJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		MonsterID:  jParams.ParentIDs,
		PropertyID: jParams.ChildIDs,
	})
}

func getMonsterRonsoRages(m seedtypes.Monster) ([]seedtypes.RonsoRage, error) {
	return lookup.GetResources(m.RonsoRages, lookup.L.RonsoRages)
}

func MonstersRonsoRages(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + ronso rages"
	jParams, err := processJunctions(desc, sh.RawJSON.Monsters, getMonsterRonsoRages)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersRonsoRagesJunctionBulk(ctx, database.CreateMonstersRonsoRagesJunctionBulkParams{
		DataHash:    jParams.DataHashes,
		MonsterID:   jParams.ParentIDs,
		RonsoRageID: jParams.ChildIDs,
	})
}

func getMonsterStatusResists(m seedtypes.Monster) ([]seedtypes.StatusResist, error) {
	return m.StatusResists, nil
}

func MonstersStatusResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + status resists"
	jParams, err := processJunctions(desc, sh.RawJSON.Monsters, getMonsterStatusResists)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersStatusResistsJunctionBulk(ctx, database.CreateMonstersStatusResistsJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		MonsterID:      jParams.ParentIDs,
		StatusResistID: jParams.ChildIDs,
	})
}
