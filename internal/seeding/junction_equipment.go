package seeding

import (
	"context"
	"database/sql"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EquipmentTableNameJunction struct {
	StdJunction
	CelestialWeaponID *int32
}

func (j EquipmentTableNameJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		h.DerefOrNil(j.CelestialWeaponID),
	}
}

func (j EquipmentTableNameJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}

func (l *Lookup) getEquipmentReqAutoAbilities(et EquipmentTable) ([]AutoAbility, error) {
	return getResources(et.RequiredAutoAbilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncEquipmentReqAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment table + required auto-abilities"
	jParams, err := processJunctions(l, desc, l.json.equipment, l.getEquipmentReqAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateEquipmentTablesRequiredAutoAbilitiesJunctionBulk(ctx, database.CreateEquipmentTablesRequiredAutoAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		EquipmentTableID: jParams.ParentIDs,
		AutoAbilityID:    jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncEquipmentTablesNames(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment tables + equipment names"
	params := database.CreateEquipmentTablesNamesJunctionBulkParams{
		DataHash:          make([]string, 0),
		EquipmentTableID:  make([]int32, 0),
		EquipmentNameID:   make([]int32, 0),
		CelestialWeaponID: make([]sql.NullInt32, 0),
	}

	for _, table := range l.json.equipment {
		for _, name := range table.EquipmentNames {
			j := EquipmentTableNameJunction{}
			j.ParentID = table.ID
			j.ChildID = name.ID

			if table.Classification == string(database.EquipClassCelestialWeapon) {
				var err error
				j.CelestialWeaponID, err = assignFKPtr(&name.Name, l.CelestialWeapons)
				if err != nil {
					return err
				}
			}

			dataHash := generateJunctionHash(j, desc)

			params.DataHash = append(params.DataHash, dataHash)
			params.EquipmentTableID = append(params.EquipmentTableID, table.ID)
			params.EquipmentNameID = append(params.EquipmentNameID, name.ID)
			params.CelestialWeaponID = append(params.CelestialWeaponID, h.GetNullInt32(j.CelestialWeaponID))
		}
	}

	return qtx.CreateEquipmentTablesNamesJunctionBulk(ctx, params)
}

func (l *Lookup) getAbilityPools() []AbilityPool {
	pools := []AbilityPool{}

	for _, table := range l.json.equipment {
		pools = append(pools, table.SelectableAutoAbilities...)
	}

	return pools
}

func (l *Lookup) getAbilityPoolAutoAbilities(ap AbilityPool) ([]AutoAbility, error) {
	return getResources(ap.AutoAbilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncAbilityPoolsAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "ability pools + auto-abilities"
	jParams, err := processJunctions(l, desc, l.getAbilityPools(), l.getAbilityPoolAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAbilityPoolsAutoAbilitiesJunctionBulk(ctx, database.CreateAbilityPoolsAutoAbilitiesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AbilityPoolID: jParams.ParentIDs,
		AutoAbilityID: jParams.ChildIDs,
	})
}
