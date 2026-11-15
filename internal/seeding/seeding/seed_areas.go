package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3SeedAreas(qtx *database.Queries, ctx context.Context) error {
	areas, err := extractAreas()
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
		params.DataHash[i] = seedhelpers.GenerateDataHash(a)
		params.SublocationID[i] = a.SublocationID
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
		lookup.L.Areas[seedtypes.Key(areas[i])] = areas[i]
		lookup.L.AreasID[row.ID] = areas[i]
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAreas() ([]seedtypes.Area, error) {
	areas := []seedtypes.Area{}
	var err error

	for i := range seedhelpers.RawJSON.Locations {
		location := &seedhelpers.RawJSON.Locations[i]

		for j := range location.Sublocations {
			sublocation := &location.Sublocations[j]

			sublocation.ID, err = seedhelpers.GetHashID(sublocation)
			if err != nil {
				return nil, err
			}

			for k := range sublocation.Areas {
				area := &sublocation.Areas[k]
				area.Location = location.Name
				area.LocationID = location.ID
				area.Sublocation = sublocation.Name
				area.SublocationID = sublocation.ID
				areas = append(areas, *area)
			}
		}
	}

	return dedupeRows(areas, seedhelpers.Hashes), nil
}

func completeAreas(areas []seedtypes.Area) error {
	for i := range areas {
		area := &areas[i]

		err := assignIDs(area.ConnectedAreas)
		if err != nil {
			return err
		}

		area.ID, err = seedhelpers.GetHashID(area)
		if err != nil {
			return err
		}

		lookup.L.Areas[seedtypes.Key(area)] = *area
		lookup.L.AreasID[area.ID] = *area
	}

	return nil
}

func Loop4SeedAreaConnections(qtx *database.Queries, ctx context.Context) error {
	areas, err := extractAreaConnections()
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
		params.DataHash[i] = seedhelpers.GenerateDataHash(a)
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
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAreaConnections() ([]seedtypes.AreaConnection, error) {
	areas := []seedtypes.AreaConnection{}

	for i := range seedhelpers.RawJSON.Locations {
		location := &seedhelpers.RawJSON.Locations[i]

		for j := range location.Sublocations {
			sublocation := &location.Sublocations[j]

			connAreas, err := prepareAreaConnections(sublocation.Areas)
			if err != nil {
				return nil, err
			}

			areas = append(areas, connAreas...)
		}
	}

	return dedupeRows(areas, seedhelpers.Hashes), nil
}

func prepareAreaConnections(areas []seedtypes.Area) ([]seedtypes.AreaConnection, error) {
	connectedAreas := []seedtypes.AreaConnection{}
	var err error

	for i := range areas {
		area := &areas[i]

		for j := range area.ConnectedAreas {
			connArea := &area.ConnectedAreas[j]

			connArea.AreaID, err = seedhelpers.AssignFK(connArea.LocationArea, lookup.L.Areas)
			if err != nil {
				return nil, err
			}

			connectedAreas = append(connectedAreas, *connArea)
		}
	}

	return connectedAreas, nil
}
