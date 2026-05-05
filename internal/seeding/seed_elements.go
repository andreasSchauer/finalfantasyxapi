package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop1SeedElements(qtx *database.Queries, ctx context.Context) error {
	elements := dedupeRows(l.json.elements, l.Hashes)

	params := database.CreateElementBulkParams{
		DataHash: make([]string, len(elements)),
		Name:     make([]string, len(elements)),
	}

	for i, e := range elements {
		params.DataHash[i] = generateDataHash(e)
		params.Name[i] = e.Name
	}

	dbRows, err := qtx.CreateElementBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create elements: %v", err)
	}

	for i, row := range dbRows {
		elements[i].ID = row.ID
		l.json.elements[i].ID = row.ID
		l.Elements[Key(elements[i])] = elements[i]
		l.ElementsID[row.ID] = elements[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop2UpdateElements(qtx *database.Queries, ctx context.Context) error {
	elements, err := l.extractUpdatedElements()
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
		params.DataHash[i] = generateDataHash(e)
		params.OppositeElementID[i] = h.GetNullInt32(e.OppositeElementID)
	}

	dbRows, err := qtx.UpdateElementBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't update elements: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractUpdatedElements() ([]Element, error) {
	elements := []Element{}
	var err error

	for i := range l.json.elements {
		element := &l.json.elements[i]

		if element.OppositeElement == nil {
			continue
		}

		delete(l.Hashes, generateDataHash(element))

		element.OppositeElementID, err = assignFKPtr(element.OppositeElement, l.Elements)
		if err != nil {
			return nil, err
		}

		elements = append(elements, *element)
	}

	return dedupeRows(elements, l.Hashes), nil
}
