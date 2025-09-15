package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type MonsterFormationList struct {
	//id 		int32
	//dataHash	string
	Version			*int32			`json:"version"`
	LocationArea	LocationArea 	`json:"location_area"`
	AreaID			int32
	Notes			*string 		`json:"notes"`
}


func(ml MonsterFormationList) ToHashFields() []any {
	return []any{
		derefOrNil(ml.Version),
		ml.AreaID,
		derefOrNil(ml.Notes),
	}
}


func seedMonsterFormations(qtx *database.Queries, lookup map[string]int32) error {
	const srcPath = "./data/monster_formations.json"

	var monsterFormationLists []MonsterFormationList
	err := loadJSONFile(string(srcPath), &monsterFormationLists)
	if err != nil {
		return err
	}


	for _, list := range monsterFormationLists {
		locationArea := list.LocationArea
		locationAreaID, err := getAreaID(locationArea, lookup)
		if err != nil {
			return fmt.Errorf("monster formations: %v", err)
		}

		list.AreaID = locationAreaID

		err = qtx.CreateMonsterFormationList(context.Background(), database.CreateMonsterFormationListParams{
			DataHash: 		generateDataHash(list),
			Version: 		getNullInt32(list.Version),
			AreaID: 		list.AreaID,	
			Notes: 			getNullString(list.Notes),
		})
		if err != nil {
			return fmt.Errorf("couldn't create monster formation list: %s - %s - %d - %s - %d - %s: %v", locationArea.Location, locationArea.SubLocation, derefOrNil(locationArea.SVersion), locationArea.Area, derefOrNil(locationArea.AVersion), derefOrNil(list.Notes), err)
		}
	}
	return nil

}

