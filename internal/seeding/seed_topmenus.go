package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Topmenu struct {
	ID			int32
	Name		string		`json:"name"`
}

func (t Topmenu) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", t),
		t.Name,
	}
}

func (t Topmenu) GetID() int32 {
	return t.ID
}

func (t Topmenu) Error() string {
	return fmt.Sprintf("topmenu %s", t.Name)
}

func (t Topmenu) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID: 	t.ID,
		Name:	t.Name,
	}
}


func (l *Lookup) loop1SeedTopmenus(qtx *database.Queries, ctx context.Context) error {
	topmenus := dedupeRows(l.json.topmenus, l.Hashes)

	params := database.CreateTopmenuBulkParams{
		DataHash: make([]string, len(topmenus)),
		Name:     make([]string, len(topmenus)),
	}

	for i, c := range topmenus {
		params.DataHash[i] = generateDataHash(c)
		params.Name[i] = c.Name
	}

	dbRows, err := qtx.CreateTopmenuBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create topmenu: %v", err)
	}

	for i, row := range dbRows {
		topmenus[i].ID = row.ID
		l.json.topmenus[i].ID = row.ID
		l.Topmenus[topmenus[i].Name] = topmenus[i]
		l.TopmenusID[row.ID] = topmenus[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}