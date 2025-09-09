package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Location struct {
	Name 			string 			`json:"location"`
	SubLocations 	[]SubLocation 	`json:"sub_locations"`
}


func(l Location) ToHashFields() []any {
	return []any{
		l.Name,
	}
}


type SubLocation struct {	
	locationID		int32
	Name			string		`json:"sub_location"`
	Specification	*string		`json:"specification"`
	Areas			[]Area		`json:"areas"`
}


func(s SubLocation) ToHashFields() []any {
	return []any{
		s.locationID,
		s.Name,
		derefOrNil(s.Specification),
	}
}


type Area struct {
	//id 					int32
	//dataHash				string
	SubLocationID			int32
	Name					string		`json:"area"`
	Section					*string		`json:"section"`
	CanRevisit				bool		`json:"can_revisit"`
	HasSaveSphere			bool		`json:"has_save_sphere"`
	AirshipDropOff			bool		`json:"airship_drop_off"`
	HasCompilationSphere	bool		`json:"has_compilation_sphere"`
}


func(a Area) ToHashFields() []any {
	return []any{
		a.SubLocationID,
		a.Name,
		derefOrNil(a.Section),
		a.CanRevisit,
		a.HasSaveSphere,
		a.AirshipDropOff,
		a.HasCompilationSphere,
	}
}




func seedLocations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/locations.json"

	var locations []Location
	err := loadJSONFile(string(srcPath), &locations)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, location := range locations {
			dbLocation, err := qtx.CreateLocation(context.Background(), database.CreateLocationParams{
				DataHash: 	generateDataHash(location),
				Name: 		location.Name,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Location: %s: %v", location.Name, err)
			}
				
			err = seedSubLocations(qtx, location, dbLocation.ID)
			if err != nil {
				return err
			}
		}
		
		return nil
	})
}


func seedSubLocations(qtx *database.Queries, location Location, locationID int32) error {
	for _, subLocation := range location.SubLocations {
		subLocation.locationID = locationID

		dbSubLocation, err := qtx.CreateSubLocation(context.Background(), database.CreateSubLocationParams{
			DataHash: 		generateDataHash(subLocation),
			LocationID: 	subLocation.locationID,
			Name: 			subLocation.Name,
			Specification: 	getNullString(subLocation.Specification),
		})
		if err != nil {
			return fmt.Errorf("couldn't create Sub Location: %s - %s: %v", location.Name, subLocation.Name, err)
		}
		
		err = seedAreas(qtx, location, subLocation, dbSubLocation.ID)
		if err != nil {
			return err
		}
	}

	return nil
}



func seedAreas(qtx *database.Queries, location Location, subLocation SubLocation, subLocationID int32) error {
	for _, area := range subLocation.Areas {
		area.SubLocationID = subLocationID
		
		err := qtx.CreateArea(context.Background(), database.CreateAreaParams{
			DataHash: 				generateDataHash(area),
			SubLocationID: 			area.SubLocationID,
			Name: 					area.Name,
			Section: 				getNullString(area.Section),
			CanRevisit: 			area.CanRevisit,
			HasSaveSphere: 			area.HasSaveSphere,
			AirshipDropOff: 		area.AirshipDropOff,
			HasCompilationSphere: 	area.HasCompilationSphere,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Area: %s - %s - %s: %v", location.Name, subLocation.Name, area.Name, err)
		}
	}

	return nil
}