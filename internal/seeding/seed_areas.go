package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
		l.Areas[Key(areas[i])] = areas[i]
		l.AreasID[row.ID] = areas[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAreas() ([]Area, error) {
	areas := []Area{}
	var err error

	for i := range l.json.locations {
		location := &l.json.locations[i]

		for j := range location.Sublocations {
			sublocation := &location.Sublocations[j]

			sublocation.ID, err = l.getHashID(sublocation)
			if err != nil {
				return nil, err
			}

			for k := range sublocation.Areas {
				area := &sublocation.Areas[k]
				area.Sublocation = *sublocation
				areas = append(areas, *area)
			}
		}
	}

	return dedupeRows(areas, l.Hashes), nil
}

func (l *Lookup) completeAreas(areas []Area) error {
	for i := range areas {
		area := &areas[i]

		err := assignIDs(l, area.ConnectedAreas)
		if err != nil {
			return err
		}

		area.ID, err = l.getHashID(area)
		if err != nil {
			return err
		}

		l.Areas[Key(area)] = *area
		l.AreasID[area.ID] = *area
	}

	return nil
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

func (l *Lookup) getAreas() []Area {
	areas := []Area{}

	for _, location := range l.json.locations {
		for _, sublocation := range location.Sublocations {
			areas = append(areas, sublocation.Areas...)
		}
	}

	return areas
}

func (l *Lookup) getAreaConnectedAreas(a Area) ([]AreaConnection, error) {
	return a.ConnectedAreas, nil
}

func (l *Lookup) seedJuncAreaConnectedAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "area + area connections"
	jParams, err := processJunctions(l, desc, l.getAreas(), l.getAreaConnectedAreas)
	if err != nil {
		return err
	}

	return qtx.CreateAreaConnectedAreasJunctionBulk(ctx, database.CreateAreaConnectedAreasJunctionBulkParams{
		DataHash:     jParams.DataHashes,
		AreaID:       jParams.ParentIDs,
		ConnectionID: jParams.ChildIDs,
	})
}
