package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type MonsterFormationList struct {
	//id 		int32
	//dataHash	string
	Version      *int32       `json:"version"`
	LocationArea LocationArea `json:"location_area"`
	AreaID       int32
	Notes        *string `json:"notes"`
}

func (ml MonsterFormationList) ToHashFields() []any {
	return []any{
		derefOrNil(ml.Version),
		ml.AreaID,
		derefOrNil(ml.Notes),
	}
}


func (l *lookup) seedMonsterFormations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var monsterFormationLists []MonsterFormationList
	err := loadJSONFile(string(srcPath), &monsterFormationLists)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, list := range monsterFormationLists {
			var err error

			locationArea := list.LocationArea
			list.AreaID, err = assignFK(locationArea, l.getArea)
			if err != nil {
				return fmt.Errorf("monster formations: %v", err)
			}

			err = qtx.CreateMonsterFormationList(context.Background(), database.CreateMonsterFormationListParams{
				DataHash: generateDataHash(list),
				Version:  getNullInt32(list.Version),
				AreaID:   list.AreaID,
				Notes:    getNullString(list.Notes),
			})
			if err != nil {
				return fmt.Errorf("couldn't create monster formation list: %s - %s - %s - %d - %s: %v", locationArea.Location, locationArea.SubLocation, locationArea.Area, derefOrNil(locationArea.Version), derefOrNil(list.Notes), err)
			}
		}
		return nil
	})
}
