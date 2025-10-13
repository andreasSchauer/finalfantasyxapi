package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AgilityTier struct {
	ID					int32
	MinAgility       	int32            `json:"min_agility"`
	MaxAgility       	int32            `json:"max_agility"`
	TickSpeed        	int32            `json:"tick_speed"`
	MonsterMinICV    	*int32           `json:"monster_min_icv"`
	MonsterMaxICV    	*int32           `json:"monster_max_icv"`
	CharacterMaxICV  	*int32           `json:"character_max_icv"`
	CharacterMinICVs 	[]AgilitySubtier `json:"character_min_icvs"`
}

func (a AgilityTier) ToHashFields() []any {
	return []any{
		a.MinAgility,
		a.MaxAgility,
		a.TickSpeed,
		derefOrNil(a.MonsterMinICV),
		derefOrNil(a.MonsterMaxICV),
		derefOrNil(a.CharacterMaxICV),
	}
}

type AgilitySubtier struct {
	//id 			int32
	//dataHash		string
	AgilityTierID     int32
	SubtierMinAgility int32  `json:"subtier_min_agility"`
	SubtierMaxAgility int32  `json:"subtier_max_agility"`
	CharacterMinICV   *int32 `json:"character_min_icv"`
}

func (a AgilitySubtier) ToHashFields() []any {
	return []any{
		a.AgilityTierID,
		a.SubtierMinAgility,
		a.SubtierMaxAgility,
		derefOrNil(a.CharacterMinICV),
	}
}

func (l *lookup) seedAgilityTiers(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/agility_tiers.json"

	var agilityTiers []AgilityTier
	err := loadJSONFile(string(srcPath), &agilityTiers)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for i, agilityTier := range agilityTiers {
			dbAgilityTier, err := qtx.CreateAgilityTier(context.Background(), database.CreateAgilityTierParams{
				DataHash:        generateDataHash(agilityTier),
				MinAgility:      agilityTier.MinAgility,
				MaxAgility:      agilityTier.MaxAgility,
				TickSpeed:       agilityTier.TickSpeed,
				MonsterMinIcv:   getNullInt32(agilityTier.MonsterMinICV),
				MonsterMaxIcv:   getNullInt32(agilityTier.MonsterMaxICV),
				CharacterMaxIcv: getNullInt32(agilityTier.CharacterMaxICV),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Agility Tier: %d: %v", i, err)
			}

			agilityTier.ID = dbAgilityTier.ID

			err = l.seedAgilitySubtiers(qtx, agilityTier)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (l *lookup) seedAgilitySubtiers(qtx *database.Queries, agilityTier AgilityTier) error {
	for i, subtier := range agilityTier.CharacterMinICVs {
		subtier.AgilityTierID = agilityTier.ID

		err := qtx.CreateAgilitySubtier(context.Background(), database.CreateAgilitySubtierParams{
			DataHash:          generateDataHash(subtier),
			AgilityTierID:     subtier.AgilityTierID,
			SubtierMinAgility: subtier.SubtierMinAgility,
			SubtierMaxAgility: subtier.SubtierMaxAgility,
			CharacterMinIcv:   getNullInt32(subtier.CharacterMinICV),
		})
		if err != nil {
			agilityTierIndex := agilityTier.ID - 1
			return fmt.Errorf("couldn't create Agility Subtier: %d - %d: %v", agilityTierIndex, i, err)
		}
	}

	return nil
}
