package seeding

import (
	"context"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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

func (l *Lookup) getAeonBaseStatsA(a Aeon) ([]BaseStat, error) {
	return a.BaseStats.AVals, nil
}

func (l *Lookup) seedJuncAeonBaseStatsA(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats a"
	jParams, err := processJunctions(l, desc, l.json.aeons, l.getAeonBaseStatsA)
	if err != nil {
		return err
	}

	return qtx.CreateAeonsBaseStatAJunctionBulk(ctx, database.CreateAeonsBaseStatAJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		AeonID:     jParams.ParentIDs,
		BaseStatID: jParams.ChildIDs,
	})
}

func (l *Lookup) getAeonBaseStatsB(a Aeon) ([]BaseStat, error) {
	return a.BaseStats.BVals, nil
}

func (l *Lookup) seedJuncAeonBaseStatsB(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats b"
	jParams, err := processJunctions(l, desc, l.json.aeons, l.getAeonBaseStatsB)
	if err != nil {
		return err
	}

	return qtx.CreateAeonsBaseStatBJunctionBulk(ctx, database.CreateAeonsBaseStatBJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		AeonID:     jParams.ParentIDs,
		BaseStatID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAeonBaseStatsX(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats x"
	params := database.CreateAeonsBaseStatXJunctionBulkParams{
		DataHash:   make([]string, 0),
		AeonID:     make([]int32, 0),
		BaseStatID: make([]int32, 0),
		Battles:    make([]int32, 0),
	}

	for _, aeon := range l.json.aeons {
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
