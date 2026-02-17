package seeding

import (
	"context"
	"database/sql"

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
	Junction
	Battles   int32
}

func (j AeonXStatJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		j.Battles,
	}
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
	
			asJunction.Junction, err = createJunctionSeed(qtx, aeon, baseStat, l.seedBaseStat)
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
