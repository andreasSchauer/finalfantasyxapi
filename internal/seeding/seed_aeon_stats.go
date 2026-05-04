package seeding

import (
	"context"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AeonStat struct {
	AeonID int32
	Name   string     `json:"name"`
	AVals  []BaseStat `json:"a_vals"`
	BVals  []BaseStat `json:"b_vals"`
	XVals  []XVal     `json:"x_vals"`
}

type XVal struct {
	Battles   int32      `json:"battles"`
	BaseStats []BaseStat `json:"base_stats"`
}

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


func (l *Lookup) completeAeonStats() error {
	for i := range l.json.aeonStats {
		as := &l.json.aeonStats[i]
		err := assignIDs(l, as.AVals)
		if err != nil {
			return err
		}

		err = assignIDs(l, as.BVals)
		if err != nil {
			return err
		}

		err = l.completeAeonXVals(as.XVals)
		if err != nil {
			return err
		}

		aeon := l.Aeons[as.Name]
		as.AeonID = aeon.ID
		aeon.BaseStats = *as
		l.Aeons[as.Name] = aeon
		l.AeonsID[as.AeonID] = aeon
		l.json.aeons[i].BaseStats = *as
	}

	return nil
}

func (l *Lookup) completeAeonXVals(xVals []XVal) error {
	for i := range xVals {
		xVal := &xVals[i]
		err := assignIDs(l, xVal.BaseStats)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) getAeonBaseStatsA(a Aeon) ([]BaseStat, error) {
	return a.BaseStats.AVals, nil
}

func (l *Lookup) getAeonBaseStatsB(a Aeon) ([]BaseStat, error) {
	return a.BaseStats.BVals, nil
}

func (l *Lookup) seedJuncAeonBaseStatsA(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats a"
	jParams, err := processJunctions(l, desc, l.json.aeons, l.getAeonBaseStatsA)
	if err != nil {
		return err
	}

	return qtx.CreateAeonsBaseStatAJunctionBulk(ctx, database.CreateAeonsBaseStatAJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		AeonID: 		jParams.ParentIDs,
		BaseStatID:  	jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAeonBaseStatsB(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats b"
	jParams, err := processJunctions(l, desc, l.json.aeons, l.getAeonBaseStatsB)
	if err != nil {
		return err
	}

	return qtx.CreateAeonsBaseStatBJunctionBulk(ctx, database.CreateAeonsBaseStatBJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		AeonID: 		jParams.ParentIDs,
		BaseStatID:  	jParams.ChildIDs,
	})
}


func (l *Lookup) seedJuncAeonBaseStatsX(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + base stats x"
	params := database.CreateAeonsBaseStatXJunctionBulkParams{
		DataHash: 	make([]string, 0),
		AeonID:  	make([]int32, 0),
		BaseStatID: make([]int32, 0),
		Battles: 	make([]int32, 0),
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