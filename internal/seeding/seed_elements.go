package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Element struct {
	ID                int32
	Name              string  `json:"name"`
	OppositeElement   *string `json:"opposite_element"`
	OppositeElementID *int32
}

func (e Element) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.Name,
		h.DerefOrNil(e.OppositeElementID),
	}
}

func (e Element) GetID() int32 {
	return e.ID
}

func (e Element) Error() string {
	return fmt.Sprintf("element %s", e.Name)
}

func (e Element) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   e.ID,
		Name: e.Name,
	}
}

func (l *Lookup) seedElements(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/elements.json"

	var elements []Element
	err := loadJSONFile(string(srcPath), &elements)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, element := range elements {
			dbElement, err := qtx.CreateElement(context.Background(), database.CreateElementParams{
				DataHash: generateDataHash(element),
				Name:     element.Name,
			})
			if err != nil {
				return h.NewErr(element.Error(), err, "couldn't create element")
			}

			element.ID = dbElement.ID
			l.Elements[element.Name] = element
			l.ElementsID[element.ID] = element
		}
		return nil
	})
}

func (l *Lookup) seedElementsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/elements.json"

	var elements []Element
	err := loadJSONFile(string(srcPath), &elements)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonElement := range elements {
			element, err := GetResource(jsonElement.Name, l.Elements)
			if err != nil {
				return err
			}

			element.OppositeElementID, err = assignFKPtr(element.OppositeElement, l.Elements)
			if err != nil {
				return h.NewErr(element.Error(), err)
			}

			err = qtx.UpdateElement(context.Background(), database.UpdateElementParams{
				DataHash:          generateDataHash(element),
				OppositeElementID: h.GetNullInt32(element.OppositeElementID),
				ID:                element.ID,
			})
			if err != nil {
				return h.NewErr(element.Error(), err, "couldn't update element")
			}

			l.Elements[element.Name] = element
		}
		return nil
	})
}

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
		l.Elements[elements[i].Name] = elements[i]
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