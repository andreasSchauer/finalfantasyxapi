package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)



func seedLocationBased(db *database.Queries, dbConn *sql.DB) error {
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		locationAreas, err := qtx.GetLocationAreas(context.Background())
		if err != nil {
			return err
		}

		locationAreaToID := make(map[string]int32)
		for _, locationArea := range locationAreas {
			keyStruct := LocationArea{
				Location: locationArea.LocationName.String,
				SubLocation: locationArea.SubLocationName.String,
				SVersion: convertNullInt32(locationArea.SVersion),
				Area: locationArea.AreaName,
				AVersion: convertNullInt32(locationArea.AVersion),
			}
			key := generateDataHash(keyStruct)
			locationAreaToID[key] = locationArea.AreaID
		}

		err = seedTreasures(qtx, locationAreaToID)
		if err != nil {
			return err
		}

		err = seedShops(qtx, locationAreaToID)
		if err != nil {
			return err
		}

		err = seedMonsterFormations(qtx, locationAreaToID)
		if err != nil {
			return err
		}

		return nil
	})
}




func getAreaID(locationArea LocationArea, lookup map[string]int32) (int32, error) {
	locationKey := generateDataHash(locationArea)
	locationAreaID, found := lookup[locationKey]
	if !found {
		return 0, fmt.Errorf("couldn't find location area: %s - %s - %d - %s - %d", locationArea.Location, locationArea.SubLocation, derefOrNil(locationArea.SVersion), locationArea.Area, derefOrNil(locationArea.AVersion))
	}

	return locationAreaID, nil
}