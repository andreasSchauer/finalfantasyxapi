package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedElements(qtx *database.Queries, ctx context.Context) error {
	elements := dedupeRows(sh.RawJSON.Elements, sh.Hashes)

	params := database.CreateElementBulkParams{
		DataHash: make([]string, len(elements)),
		Name:     make([]string, len(elements)),
	}

	for i, e := range elements {
		params.DataHash[i] = sh.GenerateDataHash(e)
		params.Name[i] = e.Name
	}

	dbRows, err := qtx.CreateElementBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create elements: %v", err)
	}

	for i, row := range dbRows {
		elements[i].ID = row.ID
		sh.RawJSON.Elements[i].ID = row.ID
		lookup.L.Elements[seedtypes.Key(elements[i])] = elements[i]
		lookup.L.ElementsID[row.ID] = elements[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func loop2UpdateElements(qtx *database.Queries, ctx context.Context) error {
	elements, err := extractUpdatedElements()
	if err != nil {
		return err
	}

	params := database.UpdateElementBulkParams{
		ID:                make([]int32, len(elements)),
		DataHash:          make([]string, len(elements)),
		OppositeElementID: make([]sql.NullInt32, len(elements)),
	}

	for i, e := range elements {
		params.ID[i] = e.ID
		params.DataHash[i] = sh.GenerateDataHash(e)
		params.OppositeElementID[i] = h.GetNullInt32(e.OppositeElementID)
	}

	dbRows, err := qtx.UpdateElementBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't update elements: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractUpdatedElements() ([]seedtypes.Element, error) {
	elements := []seedtypes.Element{}
	var err error

	for i := range sh.RawJSON.Elements {
		element := &sh.RawJSON.Elements[i]

		if element.OppositeElement == nil {
			continue
		}

		delete(sh.Hashes, sh.GenerateDataHash(element))

		element.OppositeElementID, err = sh.AssignFKPtr(element.OppositeElement, lookup.L.Elements)
		if err != nil {
			return nil, err
		}

		elements = append(elements, *element)
	}

	return dedupeRows(elements, sh.Hashes), nil
}
