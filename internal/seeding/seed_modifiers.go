package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Modifier struct {
	ID				int32
	Name       		string		`json:"name"`
	Effect			string		`json:"effect"`
	Type			string		`json:"type"`
	DefaultValue	*float32	`json:"default_value"`
}

func (m Modifier) ToHashFields() []any {
	return []any{
		m.Name,
		m.Effect,
		m.Type,
		derefOrNil(m.DefaultValue),
	}
}



func (l *lookup) seedModifiers(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/modifiers.json"

	var modifiers []Modifier
	err := loadJSONFile(string(srcPath), &modifiers)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, modifier := range modifiers {
			dbModifier, err := qtx.CreateModifier(context.Background(), database.CreateModifierParams{
				DataHash:     	generateDataHash(modifier),
				Name:         	modifier.Name,
				Effect: 		modifier.Effect,
				Type: 			database.ModifierType(modifier.Type),
				DefaultValue: 	getNullFloat64(modifier.DefaultValue),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Modifier: %s: %v", modifier.Name, err)
			}

			modifier.ID = dbModifier.ID
			l.modifiers[modifier.Name] = modifier
		}
		return nil
	})
}
