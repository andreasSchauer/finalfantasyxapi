package seeding

import (
	"context"
	"database/sql"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
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

func getEquipmentReqAutoAbilities(et seedtypes.EquipmentTable) ([]seedtypes.AutoAbility, error) {
	return lookup.GetResources(et.RequiredAutoAbilities, lookup.L.AutoAbilities)
}

func SeedJuncEquipmentReqAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment table + required auto-abilities"
	jParams, err := processJunctions(desc, sh.RawJSON.Equipment, getEquipmentReqAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateEquipmentTablesRequiredAutoAbilitiesJunctionBulk(ctx, database.CreateEquipmentTablesRequiredAutoAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		EquipmentTableID: jParams.ParentIDs,
		AutoAbilityID:    jParams.ChildIDs,
	})
}

func SeedJuncEquipmentTablesNames(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment tables + equipment names"
	params := database.CreateEquipmentTablesNamesJunctionBulkParams{
		DataHash:          make([]string, 0),
		EquipmentTableID:  make([]int32, 0),
		EquipmentNameID:   make([]int32, 0),
		CelestialWeaponID: make([]sql.NullInt32, 0),
	}

	for _, table := range sh.RawJSON.Equipment {
		for _, name := range table.EquipmentNames {
			j := EquipmentTableNameJunction{
				ParentID: table.ID,
				ChildID:  name.ID,
			}

			if table.Classification == string(database.EquipClassCelestialWeapon) {
				var err error
				j.CelestialWeaponID, err = sh.AssignFKPtr(&name.Name, lookup.L.CelestialWeapons)
				if err != nil {
					return err
				}
			}

			dataHash := GenerateJunctionHash(j, desc)

			params.DataHash = append(params.DataHash, dataHash)
			params.EquipmentTableID = append(params.EquipmentTableID, table.ID)
			params.EquipmentNameID = append(params.EquipmentNameID, name.ID)
			params.CelestialWeaponID = append(params.CelestialWeaponID, h.GetNullInt32(j.CelestialWeaponID))
		}
	}

	return qtx.CreateEquipmentTablesNamesJunctionBulk(ctx, params)
}

func getAbilityPools() []seedtypes.AbilityPool {
	pools := []seedtypes.AbilityPool{}

	for _, table := range sh.RawJSON.Equipment {
		pools = append(pools, table.SelectableAutoAbilities...)
	}

	return pools
}

func getAbilityPoolAutoAbilities(ap seedtypes.AbilityPool) ([]seedtypes.AutoAbility, error) {
	return lookup.GetResources(ap.AutoAbilities, lookup.L.AutoAbilities)
}

func SeedJuncAbilityPoolsAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "ability pools + auto-abilities"
	jParams, err := processJunctions(desc, getAbilityPools(), getAbilityPoolAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAbilityPoolsAutoAbilitiesJunctionBulk(ctx, database.CreateAbilityPoolsAutoAbilitiesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AbilityPoolID: jParams.ParentIDs,
		AutoAbilityID: jParams.ChildIDs,
	})
}
