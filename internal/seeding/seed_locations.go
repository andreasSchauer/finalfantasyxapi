package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Location struct {
	ID           int32
	Name         string        `json:"location"`
	SubLocations []SubLocation `json:"sub_locations"`
}

func (l Location) ToHashFields() []any {
	return []any{
		l.Name,
	}
}

func (l Location) Error() string {
	return fmt.Sprintf("location %s", l.Name)
}

type SubLocation struct {
	ID            int32
	Name          string  `json:"sub_location"`
	Specification *string `json:"specification"`
	Areas         []Area  `json:"areas"`
	Location      Location
}

func (s SubLocation) ToHashFields() []any {
	return []any{
		s.Location.ID,
		s.Name,
		h.DerefOrNil(s.Specification),
	}
}

func (s SubLocation) Error() string {
	return fmt.Sprintf("sublocation %s", s.Name)
}

type Area struct {
	ID                   int32
	Name                 string           `json:"area"`
	Version              *int32           `json:"version"`
	Specification        *string          `json:"specification"`
	StoryOnly            bool             `json:"story_only"`
	HasSaveSphere        bool             `json:"has_save_sphere"`
	AirshipDropOff       bool             `json:"airship_drop_off"`
	HasCompilationSphere bool             `json:"has_compilation_sphere"`
	CanRideChocobo       bool             `json:"can_ride_chocobo"`
	ConnectedAreas       []AreaConnection `json:"connected_areas"`
	SubLocation          SubLocation
}

func (a Area) ToHashFields() []any {
	return []any{
		a.SubLocation.ID,
		a.Name,
		h.DerefOrNil(a.Version),
		h.DerefOrNil(a.Specification),
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

func (a Area) Error() string {
	return fmt.Sprintf("area %s, version %v", a.Name, h.DerefOrNil(a.Version))
}

func (a Area) GetLocationArea() LocationArea {
	return LocationArea{
		Location:    a.SubLocation.Location.Name,
		SubLocation: a.SubLocation.Name,
		Area:        a.Name,
		Version:     a.Version,
	}
}

type LocationArea struct {
	ID          int32
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
		h.DerefOrNil(la.Version),
	}
}

func (la LocationArea) GetID() int32 {
	return la.ID
}

func (la LocationArea) Error() string {
	return fmt.Sprintf("location area with location: %s, sublocation: %s, area: %s, version: %v", la.Location, la.SubLocation, la.Area, h.DerefOrNil(la.Version))
}

type AreaConnection struct {
	ID             int32
	AreaID         int32
	LocationArea   LocationArea `json:"location_area"`
	ConnectionType string       `json:"connection_type"`
	StoryOnly      bool         `json:"story_only"`
	Notes          *string      `json:"notes"`
}

func (ac AreaConnection) ToHashFields() []any {
	return []any{
		ac.AreaID,
		ac.ConnectionType,
		ac.StoryOnly,
		h.DerefOrNil(ac.Notes),
	}
}

func (ac AreaConnection) GetID() int32 {
	return ac.ID
}

func (ac AreaConnection) Error() string {
	return fmt.Sprintf("area connection with %s", ac.LocationArea)
}

func (l *Lookup) seedLocations(db *database.Queries, dbConn *sql.DB) error {
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
				return h.GetErr(location.Error(), err, "couldn't create location")
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

func (l *Lookup) seedSubLocations(qtx *database.Queries, location Location) error {
	for _, subLocation := range location.SubLocations {
		subLocation.Location = location

		dbSubLocation, err := qtx.CreateSubLocation(context.Background(), database.CreateSubLocationParams{
			DataHash:      generateDataHash(subLocation),
			LocationID:    subLocation.Location.ID,
			Name:          subLocation.Name,
			Specification: h.GetNullString(subLocation.Specification),
		})
		if err != nil {
			return h.GetErr(subLocation.Error(), err, "couldn't create sublocation")
		}
		subLocation.ID = dbSubLocation.ID

		err = l.seedAreas(qtx, subLocation)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedAreas(qtx *database.Queries, subLocation SubLocation) error {
	for _, area := range subLocation.Areas {
		area.SubLocation = subLocation

		dbArea, err := qtx.CreateArea(context.Background(), database.CreateAreaParams{
			DataHash:             generateDataHash(area),
			SubLocationID:        area.SubLocation.ID,
			Name:                 area.Name,
			Version:              h.GetNullInt32(area.Version),
			Specification:        h.GetNullString(area.Specification),
			StoryOnly:            area.StoryOnly,
			HasSaveSphere:        area.HasSaveSphere,
			AirshipDropOff:       area.AirshipDropOff,
			HasCompilationSphere: area.HasCompilationSphere,
			CanRideChocobo:       area.CanRideChocobo,
		})
		if err != nil {
			return h.GetErr(area.Error(), err, "couldn't create area")
		}

		area.ID = dbArea.ID
		locationArea := area.GetLocationArea()

		key := createLookupKey(locationArea)
		l.Areas[key] = area
	}

	return nil
}

func (l *Lookup) seedAreasRelationships(db *database.Queries, dbConn *sql.DB) error {
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
						Location:    location.Name,
						SubLocation: subLocation.Name,
						Area:        jsonArea.Name,
						Version:     jsonArea.Version,
					}

					area, err := GetResource(locationArea, l.Areas)
					if err != nil {
						return h.GetErr(locationArea.Error(), err)
					}

					err = l.seedAreaConnections(qtx, area)
					if err != nil {
						return h.GetErr(locationArea.Error(), err)
					}
				}
			}
		}

		return nil
	})
}

func (l *Lookup) seedAreaConnections(qtx *database.Queries, area Area) error {
	for _, connection := range area.ConnectedAreas {
		junction, err := createJunctionSeed(qtx, area, connection, l.seedAreaConnection)
		if err != nil {
			return err
		}

		err = qtx.CreateAreaConnectedAreasJunction(context.Background(), database.CreateAreaConnectedAreasJunctionParams{
			DataHash:     generateDataHash(junction),
			AreaID:       junction.ParentID,
			ConnectionID: junction.ChildID,
		})
		if err != nil {
			return h.GetErr(connection.Error(), err, "couldn't junction area connection")
		}
	}

	return nil
}

func (l *Lookup) seedAreaConnection(qtx *database.Queries, connection AreaConnection) (AreaConnection, error) {
	var err error

	connection.AreaID, err = assignFK(connection.LocationArea, l.Areas)
	if err != nil {
		return AreaConnection{}, h.GetErr(connection.Error(), err)
	}

	dbConnection, err := qtx.CreateAreaConnection(context.Background(), database.CreateAreaConnectionParams{
		DataHash:       generateDataHash(connection),
		AreaID:         connection.AreaID,
		ConnectionType: database.AreaConnectionType(connection.ConnectionType),
		StoryOnly:      connection.StoryOnly,
		Notes:          h.GetNullString(connection.Notes),
	})
	if err != nil {
		return AreaConnection{}, h.GetErr(connection.Error(), err, "couldn't create connection")
	}
	connection.ID = dbConnection.ID

	return connection, nil
}
