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

type AeonBaseStatJunction struct {
	Junction
	ValueType database.AeonStatValue
	Battles   *int32
}

func (j AeonBaseStatJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		j.ValueType,
		h.DerefOrNil(j.Battles),
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

			err = l.seedAeonBaseStats(qtx, aeon, aeonStat.AVals, database.AeonStatValueA, nil)
			if err != nil {
				return h.NewErr(aeonStat.Name, err)
			}

			err = l.seedAeonBaseStats(qtx, aeon, aeonStat.BVals, database.AeonStatValueB, nil)
			if err != nil {
				return h.NewErr(aeonStat.Name, err)
			}

			for _, xVal := range aeonStat.XVals {
				err := l.seedAeonBaseStats(qtx, aeon, xVal.BaseStats, database.AeonStatValueX, &xVal.Battles)
				if err != nil {
					subjects := h.JoinErrSubjects(aeonStat.Name, string(xVal.Battles))
					return h.NewErr(subjects, err)
				}
			}
		}

		return nil
	})
}

func (l *Lookup) seedAeonBaseStats(qtx *database.Queries, aeon Aeon, baseStats []BaseStat, valType database.AeonStatValue, battles *int32) error {
	for _, baseStat := range baseStats {
		var err error
		asJunction := AeonBaseStatJunction{}

		asJunction.Junction, err = createJunctionSeed(qtx, aeon, baseStat, l.seedBaseStat)
		if err != nil {
			return h.NewErr(baseStat.Error(), err)
		}

		asJunction.ValueType = valType
		asJunction.Battles = battles

		err = qtx.CreateAeonsBaseStatJunction(context.Background(), database.CreateAeonsBaseStatJunctionParams{
			DataHash:   generateDataHash(asJunction),
			AeonID:     asJunction.ParentID,
			BaseStatID: asJunction.ChildID,
			ValueType:  asJunction.ValueType,
			Battles:    h.GetNullInt32(asJunction.Battles),
		})
		if err != nil {
			return h.NewErr(baseStat.Error(), err, "couldn't junction base stat")
		}
	}

	return nil
}
