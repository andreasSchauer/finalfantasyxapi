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
	Sublocations []Sublocation `json:"sublocations"`
}

func (l Location) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", l),
		l.Name,
	}
}

func (l Location) GetID() int32 {
	return l.ID
}

func (l Location) Error() string {
	return fmt.Sprintf("location %s", l.Name)
}

func (l Location) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   l.ID,
		Name: l.Name,
	}
}

type Sublocation struct {
	ID   int32
	Name string `json:"sublocation"`

	Areas    []Area `json:"areas"`
	Location Location
}

func (s Sublocation) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Location.ID,
		s.Name,
	}
}

func (s Sublocation) GetID() int32 {
	return s.ID
}

func (s *Sublocation) SetID(id int32) {
	s.ID = id
}

func (s Sublocation) Error() string {
	return fmt.Sprintf("sublocation %s", s.Name)
}

func (s Sublocation) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}

type Area struct {
	ID                   int32
	Name                 string           `json:"area"`
	Version              *int32           `json:"version"`
	Specification        *string          `json:"specification"`
	Availability         string           `json:"availability"`
	HasSaveSphere        bool             `json:"has_save_sphere"`
	AirshipDropOff       bool             `json:"airship_drop_off"`
	HasCompilationSphere bool             `json:"has_compilation_sphere"`
	CanRideChocobo       bool             `json:"can_ride_chocobo"`
	ConnectedAreas       []AreaConnection `json:"connected_areas"`
	Sublocation          Sublocation
}

func (a Area) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.Sublocation.ID,
		a.Name,
		h.DerefOrNil(a.Version),
		h.DerefOrNil(a.Specification),
		a.Availability,
		a.HasSaveSphere,
		a.AirshipDropOff,
		a.HasCompilationSphere,
		a.CanRideChocobo,
	}
}

func (a Area) GetID() int32 {
	return a.ID
}

func (a *Area) SetID(id int32) {
	a.ID = id
}

func (a Area) Error() string {
	return fmt.Sprintf("area '%s'", h.NameToString(a.Name, a.Version, a.Specification))
}

func (a Area) GetLocationArea() LocationArea {
	return LocationArea{
		Location:    a.Sublocation.Location.Name,
		Sublocation: a.Sublocation.Name,
		Area:        a.Name,
		Version:     a.Version,
	}
}

func (a Area) GetResParamsLocation() h.ResParamsLocation {
	return h.ResParamsLocation{
		AreaID:        a.ID,
		Location:      a.Sublocation.Location.Name,
		Sublocation:   a.Sublocation.Name,
		Area:          a.Name,
		Version:       a.Version,
		Specification: a.Specification,
	}
}

type LocationArea struct {
	ID          int32  `json:"-"`
	Location    string `json:"location"`
	Sublocation string `json:"sublocation"`
	Area        string `json:"area"`
	Version     *int32 `json:"version"`
}

func (la LocationArea) ToKeyFields() []any {
	return []any{
		la.Location,
		la.Sublocation,
		la.Area,
		h.DerefOrNil(la.Version),
	}
}

func (la LocationArea) GetID() int32 {
	return la.ID
}

func (la LocationArea) Error() string {
	return fmt.Sprintf("location area with location: %s, sublocation: %s, area: '%s'", la.Location, la.Sublocation, h.NameToString(la.Area, la.Version, nil))
}

func (la LocationArea) GetLocationArea() LocationArea {
	return la
}

type AreaConnection struct {
	ID             int32
	AreaID         int32
	LocationArea   LocationArea `json:"location_area"`
	ConnectionType string       `json:"connection_type"`
	IsStoryBased   bool         `json:"is_story_based"`
	Notes          *string      `json:"notes"`
}

func (ac AreaConnection) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", ac),
		ac.AreaID,
		ac.ConnectionType,
		ac.IsStoryBased,
		h.DerefOrNil(ac.Notes),
	}
}

func (ac AreaConnection) GetID() int32 {
	return ac.ID
}

func (ac *AreaConnection) SetID(id int32) {
	ac.ID = id
}

func (ac AreaConnection) Error() string {
	return fmt.Sprintf("area connection with %s", ac.LocationArea)
}

func (l *Lookup) seedLocations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/locations.json"

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
				return h.NewErr(location.Error(), err, "couldn't create location")
			}
			location.ID = dbLocation.ID
			l.Locations[location.Name] = location
			l.LocationsID[location.ID] = location

			err = l.seedSublocations(qtx, location)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (l *Lookup) seedSublocations(qtx *database.Queries, location Location) error {
	for _, sublocation := range location.Sublocations {
		sublocation.Location = location

		dbSublocation, err := qtx.CreateSublocation(context.Background(), database.CreateSublocationParams{
			DataHash:   generateDataHash(sublocation),
			LocationID: sublocation.Location.ID,
			Name:       sublocation.Name,
		})
		if err != nil {
			return h.NewErr(sublocation.Error(), err, "couldn't create sublocation")
		}
		sublocation.ID = dbSublocation.ID
		l.Sublocations[sublocation.Name] = sublocation
		l.SublocationsID[sublocation.ID] = sublocation

		err = l.seedAreas(qtx, sublocation)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedAreas(qtx *database.Queries, sublocation Sublocation) error {
	for _, area := range sublocation.Areas {
		area.Sublocation = sublocation

		dbArea, err := qtx.CreateArea(context.Background(), database.CreateAreaParams{
			DataHash:             generateDataHash(area),
			SublocationID:        area.Sublocation.ID,
			Name:                 area.Name,
			Version:              h.GetNullInt32(area.Version),
			Specification:        h.GetNullString(area.Specification),
			Availability:         database.AvailabilityType(area.Availability),
			HasSaveSphere:        area.HasSaveSphere,
			AirshipDropOff:       area.AirshipDropOff,
			HasCompilationSphere: area.HasCompilationSphere,
			CanRideChocobo:       area.CanRideChocobo,
		})
		if err != nil {
			return h.NewErr(area.Error(), err, "couldn't create area")
		}

		area.ID = dbArea.ID
		locationArea := area.GetLocationArea()

		key := Key(locationArea)
		l.Areas[key] = area
		l.AreasID[area.ID] = area
	}

	return nil
}

func (l *Lookup) seedAreasRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/locations.json"

	var locations []Location
	err := loadJSONFile(string(srcPath), &locations)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, location := range locations {
			for _, sublocation := range location.Sublocations {
				for _, jsonArea := range sublocation.Areas {
					locationArea := LocationArea{
						Location:    location.Name,
						Sublocation: sublocation.Name,
						Area:        jsonArea.Name,
						Version:     jsonArea.Version,
					}

					area, err := GetResource(locationArea, l.Areas)
					if err != nil {
						return h.NewErr(locationArea.Error(), err)
					}

					err = l.seedAreaConnections(qtx, area)
					if err != nil {
						return h.NewErr(locationArea.Error(), err)
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
			return h.NewErr(connection.Error(), err, "couldn't junction area connection")
		}
	}

	return nil
}

func (l *Lookup) seedAreaConnection(qtx *database.Queries, connection AreaConnection) (AreaConnection, error) {
	var err error

	connection.AreaID, err = assignFK(connection.LocationArea, l.Areas)
	if err != nil {
		return AreaConnection{}, h.NewErr(connection.Error(), err)
	}

	dbConnection, err := qtx.CreateAreaConnection(context.Background(), database.CreateAreaConnectionParams{
		DataHash:       generateDataHash(connection),
		AreaID:         connection.AreaID,
		ConnectionType: database.AreaConnectionType(connection.ConnectionType),
		IsStoryBased:   connection.IsStoryBased,
		Notes:          h.GetNullString(connection.Notes),
	})
	if err != nil {
		return AreaConnection{}, h.NewErr(connection.Error(), err, "couldn't create connection")
	}
	connection.ID = dbConnection.ID

	return connection, nil
}

func (l *Lookup) loop1SeedLocations(qtx *database.Queries, ctx context.Context) error {
	locations := dedupeRows(l.json.locations, l.Hashes)

	params := database.CreateLocationBulkParams{
		DataHash: make([]string, len(locations)),
		Name:     make([]string, len(locations)),
	}

	for i, mi := range locations {
		params.DataHash[i] = generateDataHash(mi)
		params.Name[i] = mi.Name
	}

	dbRows, err := qtx.CreateLocationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create locations: %v", err)
	}

	for i, row := range dbRows {
		locations[i].ID = row.ID
		l.json.locations[i].ID = row.ID
		l.Locations[locations[i].Name] = locations[i]
		l.LocationsID[row.ID] = locations[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop2SeedSublocations(qtx *database.Queries, ctx context.Context) error {
	sublocations := l.extractSublocations()

	params := database.CreateSublocationBulkParams{
		DataHash:   make([]string, len(sublocations)),
		LocationID: make([]int32, len(sublocations)),
		Name:       make([]string, len(sublocations)),
	}

	for i, s := range sublocations {
		params.DataHash[i] = generateDataHash(s)
		params.LocationID[i] = s.Location.ID
		params.Name[i] = s.Name
	}

	dbRows, err := qtx.CreateSublocationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create sublocations: %v", err)
	}

	for i, row := range dbRows {
		sublocations[i].ID = row.ID
		l.Sublocations[sublocations[i].Name] = sublocations[i]
		l.SublocationsID[row.ID] = sublocations[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSublocations() []Sublocation {
	sublocations := []Sublocation{}

	for i := range l.json.locations {
		location := &l.json.locations[i]

		for j := range location.Sublocations {
			sublocation := &location.Sublocations[j]
			sublocation.Location = *location
			sublocations = append(sublocations, *sublocation)
		}
	}

	return dedupeRows(sublocations, l.Hashes)
}

func (l *Lookup) loop3SeedAreas(qtx *database.Queries, ctx context.Context) error {
	areas, err := l.extractAreas()
	if err != nil {
		return err
	}

	params := database.CreateAreaBulkParams{
		DataHash:             make([]string, len(areas)),
		SublocationID:        make([]int32, len(areas)),
		Name:                 make([]string, len(areas)),
		Version:              make([]sql.NullInt32, len(areas)),
		Specification:        make([]sql.NullString, len(areas)),
		Availability:         make([]database.AvailabilityType, len(areas)),
		HasSaveSphere:        make([]bool, len(areas)),
		AirshipDropOff:       make([]bool, len(areas)),
		HasCompilationSphere: make([]bool, len(areas)),
		CanRideChocobo:       make([]bool, len(areas)),
	}

	for i, a := range areas {
		params.DataHash[i] = generateDataHash(a)
		params.SublocationID[i] = a.Sublocation.ID
		params.Name[i] = a.Name
		params.Version[i] = h.GetNullInt32(a.Version)
		params.Specification[i] = h.GetNullString(a.Specification)
		params.Availability[i] = database.AvailabilityType(a.Availability)
		params.HasSaveSphere[i] = a.HasSaveSphere
		params.AirshipDropOff[i] = a.AirshipDropOff
		params.HasCompilationSphere[i] = a.HasCompilationSphere
		params.CanRideChocobo[i] = a.CanRideChocobo
	}

	dbRows, err := qtx.CreateAreaBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create areas: %v", err)
	}

	for i, row := range dbRows {
		areas[i].ID = row.ID
		key := Key(areas[i].GetLocationArea())
		l.Areas[key] = areas[i]
		l.AreasID[row.ID] = areas[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAreas() ([]Area, error) {
	areas := []Area{}
	var err error

	// why does area.Sublocation.ID get updated, but not sublocation.ID and location.Sublocations[i].ID?
	for i := range l.json.locations {
		location := &l.json.locations[i]

		for j := range location.Sublocations {
			sublocation := &location.Sublocations[j]

			for k := range sublocation.Areas {
				area := &sublocation.Areas[k]
				area.Sublocation = *sublocation

				area.Sublocation.ID, err = l.getHashID(area.Sublocation)
				if err != nil {
					return nil, err
				}

				areas = append(areas, *area)
			}
		}
	}

	return dedupeRows(areas, l.Hashes), nil
}

func (l *Lookup) loop4SeedAreaConnections(qtx *database.Queries, ctx context.Context) error {
	areas, err := l.extractAreaConnections()
	if err != nil {
		return err
	}

	params := database.CreateAreaConnectionBulkParams{
		DataHash:       make([]string, len(areas)),
		AreaID:         make([]int32, len(areas)),
		ConnectionType: make([]database.AreaConnectionType, len(areas)),
		IsStoryBased:   make([]bool, len(areas)),
		Notes:          make([]sql.NullString, len(areas)),
	}

	for i, a := range areas {
		params.DataHash[i] = generateDataHash(a)
		params.AreaID[i] = a.AreaID
		params.ConnectionType[i] = database.AreaConnectionType(a.ConnectionType)
		params.IsStoryBased[i] = a.IsStoryBased
		params.Notes[i] = h.GetNullString(a.Notes)
	}

	dbRows, err := qtx.CreateAreaConnectionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create area connections: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAreaConnections() ([]AreaConnection, error) {
	areas := []AreaConnection{}

	for i := range l.json.locations {
		location := &l.json.locations[i]

		for j := range location.Sublocations {
			sublocation := &location.Sublocations[j]

			connAreas, err := l.prepareAreaConnections(sublocation.Areas)
			if err != nil {
				return nil, err
			}

			areas = append(areas, connAreas...)
		}
	}

	return dedupeRows(areas, l.Hashes), nil
}

func (l *Lookup) prepareAreaConnections(areas []Area) ([]AreaConnection, error) {
	connectedAreas := []AreaConnection{}
	var err error

	for i := range areas {
		area := &areas[i]

		for j := range area.ConnectedAreas {
			connArea := &area.ConnectedAreas[j]

			connArea.AreaID, err = assignFK(connArea.LocationArea, l.Areas)
			if err != nil {
				return nil, err
			}

			connectedAreas = append(connectedAreas, *connArea)
		}
	}

	return connectedAreas, nil
}
