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
		ID: e.ID,
		Name: e.Name,
	}
}

func (l *Lookup) seedElements(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/elements.json"

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
	const srcPath = "./data/elements.json"

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
