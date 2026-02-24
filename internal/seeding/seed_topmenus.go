package seeding

import (
	"context"
	"database/sql"
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
		t.Name,
	}
}

func (t Topmenu) GetID() int32 {
	return t.ID
}

func (t Topmenu) Error() string {
	return fmt.Sprintf("topmenu %s", t.Name)
}


func (l *Lookup) seedTopmenus(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/topmenus.json"

	var topmenus []Topmenu
	err := loadJSONFile(string(srcPath), &topmenus)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, topmenu := range topmenus {
			dbTopmenu, err := qtx.CreateTopmenu(context.Background(), database.CreateTopmenuParams{
				DataHash:    generateDataHash(topmenu),
				Name:        topmenu.Name,
			})
			if err != nil {
				return h.NewErr(topmenu.Error(), err, "couldn't create topmenu")
			}

			topmenu.ID = dbTopmenu.ID
			l.Topmenus[topmenu.Name] = topmenu
			l.TopmenusID[topmenu.ID] = topmenu
		}
		return nil
	})
}