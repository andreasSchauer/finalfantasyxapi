package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AgilityTier struct {
	ID               int32
	MinAgility       int32            `json:"min_agility"`
	MaxAgility       int32            `json:"max_agility"`
	TickSpeed        int32            `json:"tick_speed"`
	MonsterMinICV    *int32           `json:"monster_min_icv"`
	MonsterMaxICV    *int32           `json:"monster_max_icv"`
	CharacterMaxICV  *int32           `json:"character_max_icv"`
	CharacterMinICVs []AgilitySubtier `json:"character_min_icvs"`
}

func (a AgilityTier) ToHashFields() []any {
	return []any{
		a.MinAgility,
		a.MaxAgility,
		a.TickSpeed,
		h.DerefOrNil(a.MonsterMinICV),
		h.DerefOrNil(a.MonsterMaxICV),
		h.DerefOrNil(a.CharacterMaxICV),
	}
}

func (a AgilityTier) Error() string {
	return fmt.Sprintf("agility tier with min agility: %d, max agility: %d", a.MinAgility, a.MaxAgility)
}

type AgilitySubtier struct {
	//id 			int32
	//dataHash		string
	AgilityTierID   int32
	MinAgility      int32  `json:"subtier_min_agility"`
	MaxAgility      int32  `json:"subtier_max_agility"`
	CharacterMinICV *int32 `json:"character_min_icv"`
}

func (a AgilitySubtier) ToHashFields() []any {
	return []any{
		a.AgilityTierID,
		a.MinAgility,
		a.MaxAgility,
		h.DerefOrNil(a.CharacterMinICV),
	}
}

func (a AgilitySubtier) Error() string {
	return fmt.Sprintf("agility subtier with min agility: %d, max agility: %d", a.MinAgility, a.MaxAgility)
}

func (l *Lookup) seedAgilityTiers(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/agility_tiers.json"

	var agilityTiers []AgilityTier
	err := loadJSONFile(string(srcPath), &agilityTiers)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, agilityTier := range agilityTiers {
			dbAgilityTier, err := qtx.CreateAgilityTier(context.Background(), database.CreateAgilityTierParams{
				DataHash:        generateDataHash(agilityTier),
				MinAgility:      agilityTier.MinAgility,
				MaxAgility:      agilityTier.MaxAgility,
				TickSpeed:       agilityTier.TickSpeed,
				MonsterMinIcv:   h.GetNullInt32(agilityTier.MonsterMinICV),
				MonsterMaxIcv:   h.GetNullInt32(agilityTier.MonsterMaxICV),
				CharacterMaxIcv: h.GetNullInt32(agilityTier.CharacterMaxICV),
			})
			if err != nil {
				return h.GetErr(agilityTier.Error(), err, "couldn't create agility tier")
			}

			agilityTier.ID = dbAgilityTier.ID

			err = l.seedAgilitySubtiers(qtx, agilityTier)
			if err != nil {
				return h.GetErr(agilityTier.Error(), err)
			}
		}
		return nil
	})
}

func (l *Lookup) seedAgilitySubtiers(qtx *database.Queries, agilityTier AgilityTier) error {
	for _, subtier := range agilityTier.CharacterMinICVs {
		subtier.AgilityTierID = agilityTier.ID

		err := qtx.CreateAgilitySubtier(context.Background(), database.CreateAgilitySubtierParams{
			DataHash:          generateDataHash(subtier),
			AgilityTierID:     subtier.AgilityTierID,
			SubtierMinAgility: subtier.MinAgility,
			SubtierMaxAgility: subtier.MaxAgility,
			CharacterMinIcv:   h.GetNullInt32(subtier.CharacterMinICV),
		})
		if err != nil {
			return h.GetErr(subtier.Error(), err, "couldn't create agility subtier")
		}
	}

	return nil
}
