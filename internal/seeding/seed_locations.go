package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Location struct {
	Name         string        `json:"location"`
	SubLocations []SubLocation `json:"sub_locations"`
}

func (l Location) ToHashFields() []any {
	return []any{
		l.Name,
	}
}

type SubLocation struct {
	locationID    int32
	Name          string  `json:"sub_location"`
	Specification *string `json:"specification"`
	Areas         []Area  `json:"areas"`
}

func (s SubLocation) ToHashFields() []any {
	return []any{
		s.locationID,
		s.Name,
		derefOrNil(s.Specification),
	}
}

type Area struct {
	//id 					int32
	//dataHash				string
	SubLocationID        int32
	Name                 string  `json:"area"`
	Version              *int32  `json:"version"`
	Specification        *string `json:"specification"`
	StoryOnly            bool    `json:"story_only"`
	HasSaveSphere        bool    `json:"has_save_sphere"`
	AirshipDropOff       bool    `json:"airship_drop_off"`
	HasCompilationSphere bool    `json:"has_compilation_sphere"`
	CanRideChocobo       bool    `json:"can_ride_chocobo"`
}

func (a Area) ToHashFields() []any {
	return []any{
		a.SubLocationID,
		a.Name,
		derefOrNil(a.Version),
		derefOrNil(a.Specification),
		a.StoryOnly,
		a.HasSaveSphere,
		a.AirshipDropOff,
		a.HasCompilationSphere,
		a.CanRideChocobo,
	}
}

type LocationArea struct {
	Location    string `json:"location"`
	SubLocation string `json:"sub_location"`
	Area        string `json:"area"`
	Version     *int32 `json:"version"`
}


func (la LocationArea) ToKeyFields() []any {
	return []any{
		la.Location,
		la.SubLocation,
		la.Area,
		derefOrNil(la.Version),
	}
}


type LocationAreaLookup struct {
	LocationArea
	ID 				int32
}


func (l *lookup) seedLocations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/locations.json"

	var locations []Location
	err := loadJSONFile(string(srcPath), &locations)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, location := range locations {
			dbLocation, err := qtx.CreateLocation(context.Background(), database.CreateLocationParams{
				DataHash: generateDataHash(location),
				Name:     location.Name,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Location: %s: %v", location.Name, err)
			}

			err = l.seedSubLocations(qtx, location, dbLocation.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (l *lookup) seedSubLocations(qtx *database.Queries, location Location, locationID int32) error {
	for _, subLocation := range location.SubLocations {
		subLocation.locationID = locationID

		dbSubLocation, err := qtx.CreateSubLocation(context.Background(), database.CreateSubLocationParams{
			DataHash:      generateDataHash(subLocation),
			LocationID:    subLocation.locationID,
			Name:          subLocation.Name,
			Specification: getNullString(subLocation.Specification),
		})
		if err != nil {
			return fmt.Errorf("couldn't create Sub Location: %s - %s: %v", location.Name, subLocation.Name, err)
		}

		err = l.seedAreas(qtx, location, subLocation, dbSubLocation.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *lookup) seedAreas(qtx *database.Queries, location Location, subLocation SubLocation, subLocationID int32) error {
	for _, area := range subLocation.Areas {
		area.SubLocationID = subLocationID

		dbArea, err := qtx.CreateArea(context.Background(), database.CreateAreaParams{
			DataHash:             generateDataHash(area),
			SubLocationID:        area.SubLocationID,
			Name:                 area.Name,
			Version:              getNullInt32(area.Version),
			Specification:        getNullString(area.Specification),
			StoryOnly:            area.StoryOnly,
			HasSaveSphere:        area.HasSaveSphere,
			AirshipDropOff:       area.AirshipDropOff,
			HasCompilationSphere: area.HasCompilationSphere,
			CanRideChocobo:       area.CanRideChocobo,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Area: %s - %s - %s: %v", location.Name, subLocation.Name, area.Name, err)
		}

		locationArea := LocationArea{
			Location: 		location.Name,
			SubLocation: 	subLocation.Name,
			Area: 			area.Name,
			Version: 		area.Version,
		}
		key := createLookupKey(locationArea)
		l.locationAreas[key] = LocationAreaLookup{
			LocationArea: 	locationArea,
			ID: 			dbArea.ID,
		}
	}

	return nil
}
