package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type AgilityTier struct {
	//id 				int32
	//dataHash			string
	MinAgility			int32				`json:"min_agility"`
	MaxAgility			int32				`json:"max_agility"`
	TickSpeed			int32				`json:"tick_speed"`
	MonsterMinICV		*int32				`json:"monster_min_icv"`
	MonsterMaxICV		*int32				`json:"monster_max_icv"`
	CharacterMaxICV		*int32				`json:"character_max_icv"`
	CharacterMinICVs	[]AgilitySubtier	`json:"character_min_icvs"`
}

func(a AgilityTier) ToHashFields() []any {
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
	AgilityTierID			int32
	SubtierMinAgility		int32 	`json:"subtier_min_agility"`
	SubtierMaxAgility		int32	`json:"subtier_max_agility"`
	CharacterMinICV			*int32	`json:"character_min_icv"`
}

func(a AgilitySubtier) ToHashFields() []any {
	return []any{
		a.AgilityTierID,
		a.SubtierMinAgility,
		a.SubtierMaxAgility,
		derefOrNil(a.CharacterMinICV),
	}
}


func seedAgilityTiers(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/agility_tiers.json"

	var agilityTiers []AgilityTier
	err := loadJSONFile(string(srcPath), &agilityTiers)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for i, agilityTier := range agilityTiers {
			tier, err := qtx.CreateAgilityTier(context.Background(), database.CreateAgilityTierParams{
				DataHash: 			generateDataHash(agilityTier),
				MinAgility: 		agilityTier.MinAgility,
				MaxAgility: 		agilityTier.MaxAgility,
				TickSpeed: 			agilityTier.TickSpeed,
				MonsterMinIcv: 		getNullInt32(agilityTier.MonsterMinICV),
				MonsterMaxIcv: 		getNullInt32(agilityTier.MonsterMaxICV),
				CharacterMaxIcv: 	getNullInt32(agilityTier.CharacterMaxICV),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Agility Tier: %d: %v", i, err)
			}

			for j, subtier := range agilityTier.CharacterMinICVs {
				subtier.AgilityTierID = tier.ID

				err = qtx.CreateAgilitySubtier(context.Background(), database.CreateAgilitySubtierParams{
					DataHash: 			generateDataHash(subtier),
					AgilityTierID: 		subtier.AgilityTierID,
					SubtierMinAgility: 	subtier.SubtierMinAgility,
					SubtierMaxAgility: 	subtier.SubtierMaxAgility,
					CharacterMinIcv: 	getNullInt32(subtier.CharacterMinICV),
				})
				if err != nil {
					return fmt.Errorf("couldn't create Agility Subtier: %d - %d: %v", i, j, err)
				}
			}
		}
		return nil
	})
}