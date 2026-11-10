package seeding

import (
	"context"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

type AeonXStatJunction struct {
	StdJunction
	Battles int32
}

func (j AeonXStatJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		j.Battles,
	}
}

func (j AeonXStatJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}

func getAeonBaseStatsA(a seedtypes.Aeon) ([]seedtypes.BaseStat, error) {
	return a.BaseStats.AVals, nil
}

func seedJuncAeonBaseStatsA(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats a"
	jParams, err := processJunctions(desc, rawJSON.aeons, getAeonBaseStatsA)
	if err != nil {
		return err
	}

	return qtx.CreateAeonsBaseStatAJunctionBulk(ctx, database.CreateAeonsBaseStatAJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		AeonID:     jParams.ParentIDs,
		BaseStatID: jParams.ChildIDs,
	})
}

func getAeonBaseStatsB(a seedtypes.Aeon) ([]seedtypes.BaseStat, error) {
	return a.BaseStats.BVals, nil
}

func seedJuncAeonBaseStatsB(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats b"
	jParams, err := processJunctions(desc, rawJSON.aeons, getAeonBaseStatsB)
	if err != nil {
		return err
	}

	return qtx.CreateAeonsBaseStatBJunctionBulk(ctx, database.CreateAeonsBaseStatBJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		AeonID:     jParams.ParentIDs,
		BaseStatID: jParams.ChildIDs,
	})
}

func seedJuncAeonBaseStatsX(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats x"
	params := database.CreateAeonsBaseStatXJunctionBulkParams{
		DataHash:   make([]string, 0),
		AeonID:     make([]int32, 0),
		BaseStatID: make([]int32, 0),
		Battles:    make([]int32, 0),
	}

	for _, aeon := range rawJSON.aeons {
		for _, xVal := range aeon.BaseStats.XVals {
			for _, baseStat := range xVal.BaseStats {
				j := AeonXStatJunction{}
				j.ParentID = aeon.ID
				j.ChildID = baseStat.ID
				j.Battles = xVal.Battles
				dataHash := generateJunctionHash(j, desc)

				params.DataHash = append(params.DataHash, dataHash)
				params.AeonID = append(params.AeonID, aeon.ID)
				params.BaseStatID = append(params.BaseStatID, baseStat.ID)
				params.Battles = append(params.Battles, xVal.Battles)
			}
		}
	}

	return qtx.CreateAeonsBaseStatXJunctionBulk(ctx, params)
}
