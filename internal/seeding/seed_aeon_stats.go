package seeding

import (
	"context"
	"database/sql"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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

func (l *Lookup) seedAeonStats(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/aeon_stats.json"

	var aeonStats []AeonStat
	err := loadJSONFile(string(srcPath), &aeonStats)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, aeonStat := range aeonStats {
			aeon, err := GetResource(aeonStat.Name, l.Aeons)
			if err != nil {
				return h.NewErr(aeonStat.Name, err)
			}

			err = l.seedAeonBaseStatsA(qtx, aeon, aeonStat.AVals)
			if err != nil {
				return h.NewErr(aeonStat.Name, err)
			}

			err = l.seedAeonBaseStatsB(qtx, aeon, aeonStat.BVals)
			if err != nil {
				return h.NewErr(aeonStat.Name, err)
			}

			err = l.seedAeonBaseStatsX(qtx, aeon, aeonStat.XVals)
			if err != nil {
				return h.NewErr(aeonStat.Name, err)
			}

			aeon.BaseStats = aeonStat
			l.Aeons[aeon.Name] = aeon
			l.AeonsID[aeon.ID] = aeon
		}

		return nil
	})
}

func (l *Lookup) seedAeonBaseStatsA(qtx *database.Queries, aeon Aeon, baseStats []BaseStat) error {
	for _, baseStat := range baseStats {
		junction, err := createJunctionSeed(qtx, aeon, baseStat, l.seedBaseStat)
		if err != nil {
			return h.NewErr(baseStat.Error(), err)
		}

		err = qtx.CreateAeonsBaseStatAJunction(context.Background(), database.CreateAeonsBaseStatAJunctionParams{
			DataHash:   generateDataHash(junction),
			AeonID:     junction.ParentID,
			BaseStatID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(baseStat.Error(), err, "couldn't junction base stat a")
		}
	}
	return nil
}

func (l *Lookup) seedAeonBaseStatsB(qtx *database.Queries, aeon Aeon, baseStats []BaseStat) error {
	for _, baseStat := range baseStats {
		junction, err := createJunctionSeed(qtx, aeon, baseStat, l.seedBaseStat)
		if err != nil {
			return h.NewErr(baseStat.Error(), err)
		}

		err = qtx.CreateAeonsBaseStatBJunction(context.Background(), database.CreateAeonsBaseStatBJunctionParams{
			DataHash:   generateDataHash(junction),
			AeonID:     junction.ParentID,
			BaseStatID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(baseStat.Error(), err, "couldn't junction base stat b")
		}
	}
	return nil
}

func (l *Lookup) seedAeonBaseStatsX(qtx *database.Queries, aeon Aeon, xVals []XVal) error {
	for _, xVal := range xVals {
		for _, baseStat := range xVal.BaseStats {
			var err error
			asJunction := AeonXStatJunction{}

			asJunction.StdJunction, err = createJunctionSeed(qtx, aeon, baseStat, l.seedBaseStat)
			if err != nil {
				return h.NewErr(baseStat.Error(), err)
			}

			asJunction.Battles = xVal.Battles

			err = qtx.CreateAeonsBaseStatXJunction(context.Background(), database.CreateAeonsBaseStatXJunctionParams{
				DataHash:   generateDataHash(asJunction),
				AeonID:     asJunction.ParentID,
				BaseStatID: asJunction.ChildID,
				Battles:    asJunction.Battles,
			})
			if err != nil {
				return h.NewErr(baseStat.Error(), err, "couldn't junction base stat x")
			}
		}
	}

	return nil
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