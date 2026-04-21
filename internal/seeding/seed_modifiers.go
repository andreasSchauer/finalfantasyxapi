package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Modifier struct {
	ID           int32
	Name         string   `json:"name"`
	Effect       string   `json:"effect"`
	Category     string   `json:"type"`
	DefaultValue *float32 `json:"default_value"`
}

func (m Modifier) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.Name,
		m.Effect,
		m.Category,
		h.DerefOrNil(m.DefaultValue),
	}
}

func (m Modifier) GetID() int32 {
	return m.ID
}

func (m Modifier) Error() string {
	return fmt.Sprintf("modifier %s", m.Name)
}

func (m Modifier) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   m.ID,
		Name: m.Name,
	}
}

func (l *Lookup) seedModifiers(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/modifiers.json"

	var modifiers []Modifier
	err := loadJSONFile(string(srcPath), &modifiers)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, modifier := range modifiers {
			dbModifier, err := qtx.CreateModifier(context.Background(), database.CreateModifierParams{
				DataHash:     generateDataHash(modifier),
				Name:         modifier.Name,
				Effect:       modifier.Effect,
				Category:     database.ModifierCategory(modifier.Category),
				DefaultValue: h.GetNullFloat64(modifier.DefaultValue),
			})
			if err != nil {
				return h.NewErr(modifier.Error(), err, "couldn't create modifier")
			}

			modifier.ID = dbModifier.ID
			l.Modifiers[modifier.Name] = modifier
			l.ModifiersID[modifier.ID] = modifier
		}
		return nil
	})
}
