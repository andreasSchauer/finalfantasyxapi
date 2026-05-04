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

func (l *Lookup) loop1SeedModifiers(qtx *database.Queries, ctx context.Context) error {
	modifiers := dedupeRows(l.json.modifiers, l.Hashes)

	params := database.CreateModifierBulkParams{
		DataHash:     make([]string, len(modifiers)),
		Name:         make([]string, len(modifiers)),
		Effect:       make([]string, len(modifiers)),
		Category:     make([]database.ModifierCategory, len(modifiers)),
		DefaultValue: make([]sql.NullFloat64, len(modifiers)),
	}

	for i, m := range modifiers {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Effect[i] = m.Effect
		params.Category[i] = database.ModifierCategory(m.Category)
		params.DefaultValue[i] = h.GetNullFloat64(m.DefaultValue)
	}

	dbRows, err := qtx.CreateModifierBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create modifiers: %v", err)
	}

	for i, row := range dbRows {
		modifiers[i].ID = row.ID
		l.json.modifiers[i].ID = row.ID
		l.Modifiers[modifiers[i].Name] = modifiers[i]
		l.ModifiersID[row.ID] = modifiers[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}