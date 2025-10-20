package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Location struct {
	ID				int32
	Name         	string        `json:"location"`
	SubLocations 	[]SubLocation `json:"sub_locations"`
}

func (l Location) ToHashFields() []any {
	return []any{
		l.Name,
	}
}

type SubLocation struct {
	ID				int32
	locationID    	int32
	Name          	string  `json:"sub_location"`
	Specification 	*string `json:"specification"`
	Areas         	[]Area  `json:"areas"`
}

func (s SubLocation) ToHashFields() []any {
	return []any{
		s.locationID,
		s.Name,
		derefOrNil(s.Specification),
	}
}

type Area struct {
	ID                   int32
	SubLocationID        int32
	Name                 string  			`json:"area"`
	Version              *int32  			`json:"version"`
	Specification        *string 			`json:"specification"`
	StoryOnly            bool    			`json:"story_only"`
	HasSaveSphere        bool    			`json:"has_save_sphere"`
	AirshipDropOff       bool    			`json:"airship_drop_off"`
	HasCompilationSphere bool    			`json:"has_compilation_sphere"`
	CanRideChocobo       bool    			`json:"can_ride_chocobo"`
	ConnectedAreas		[]AreaConnection 	`json:"connected_areas"`
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

func (a Area) GetID() int32 {
	return a.ID
}


type LocationArea struct {
	ID			int32
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


type AreaConnection struct {
	ID				int32
	AreaID			int32
	LocationArea 	LocationArea	`json:"location_area"`
	ConnectionType	string			`json:"connection_type"`
	StoryOnly		bool			`json:"story_only"`
	Notes			*string			`json:"notes"`
}


func (ac AreaConnection) ToHashFields() []any {
	return []any{
		ac.AreaID,
		ac.ConnectionType,
		ac.StoryOnly,
		derefOrNil(ac.Notes),
	}
}


func (ac AreaConnection) GetID() int32 {
	return ac.ID
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
			location.ID = dbLocation.ID

			err = l.seedSubLocations(qtx, location)
			if err != nil {
				return err
			}
		}

		return nil
	})
}


func (l *lookup) seedSubLocations(qtx *database.Queries, location Location) error {
	for _, subLocation := range location.SubLocations {
		subLocation.locationID = location.ID

		dbSubLocation, err := qtx.CreateSubLocation(context.Background(), database.CreateSubLocationParams{
			DataHash:      generateDataHash(subLocation),
			LocationID:    subLocation.locationID,
			Name:          subLocation.Name,
			Specification: getNullString(subLocation.Specification),
		})
		if err != nil {
			return fmt.Errorf("couldn't create Sub Location: %s - %s: %v", location.Name, subLocation.Name, err)
		}
		subLocation.ID = dbSubLocation.ID

		err = l.seedAreas(qtx, location, subLocation)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *lookup) seedAreas(qtx *database.Queries, location Location, subLocation SubLocation) error {
	for _, area := range subLocation.Areas {
		area.SubLocationID = subLocation.ID

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

		area.ID = dbArea.ID
		locationArea := LocationArea{
			Location:    location.Name,
			SubLocation: subLocation.Name,
			Area:        area.Name,
			Version:     area.Version,
		}
		key := createLookupKey(locationArea)
		l.areas[key] = area
	}

	return nil
}


func (l *lookup) createAreasRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/locations.json"

	var locations []Location
	err := loadJSONFile(string(srcPath), &locations)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, location := range locations {
			for _, subLocation := range location.SubLocations {
				for _, jsonArea := range subLocation.Areas {
					locationArea := LocationArea{
						Location: location.Name,
						SubLocation: subLocation.Name,
						Area: jsonArea.Name,
						Version: jsonArea.Version,
					}
					area, err := l.getArea(locationArea)
					if err != nil {
						return err
					}

					for _, connection := range area.ConnectedAreas {
						junction, err := createJunctionSeed(qtx, area, connection, l.seedAreaConnection)
						if err != nil {
							return fmt.Errorf("area: %s: %v", createLookupKey(locationArea), err)
						}

						err = qtx.CreateAreaConnectionJunction(context.Background(), database.CreateAreaConnectionJunctionParams{
							DataHash: 		generateDataHash(junction),
							AreaID: 		junction.ParentID,
							ConnectionID: 	junction.ChildID,
						})
						if err != nil {
							return fmt.Errorf("area: %s: couldn't create connected area junction: %v", createLookupKey(locationArea), err)
						}
					}
				}
			}
		}


		return nil
	})
}


func (l *lookup) seedAreaConnection(qtx *database.Queries, connection AreaConnection) (AreaConnection, error) {
	var err error

	connection.AreaID, err = assignFK(connection.LocationArea, l.getArea)
	if err != nil {
		return AreaConnection{}, err
	}

	dbConnection, err := qtx.CreateAreaConnection(context.Background(), database.CreateAreaConnectionParams{
		DataHash: 		generateDataHash(connection),
		AreaID: 		connection.AreaID,
		ConnectionType: database.AreaConnectionType(connection.ConnectionType),
		StoryOnly: 		connection.StoryOnly,
		Notes: 			getNullString(connection.Notes),
	})
	if err != nil {
		return AreaConnection{}, fmt.Errorf("couldn't create connection: %s: %v", createLookupKey(connection.LocationArea), err)
	}
	connection.ID = dbConnection.ID

	return connection, nil
}