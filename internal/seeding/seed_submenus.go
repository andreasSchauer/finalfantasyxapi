package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Submenu struct {
	//id 		int32
	//dataHash	string
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Effect      string  `json:"effect"`
	Topmenu     *string `json:"topmenu"`
}

func (s Submenu) ToHashFields() []any {
	return []any{
		s.Name,
		s.Description,
		s.Effect,
		derefOrNil(s.Topmenu),
	}
}

func (l *lookup) seedSubmenus(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/submenus.json"

	var submenus []Submenu
	err := loadJSONFile(string(srcPath), &submenus)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, menu := range submenus {
			err = qtx.CreateSubmenu(context.Background(), database.CreateSubmenuParams{
				DataHash:    generateDataHash(menu),
				Name:        menu.Name,
				Description: menu.Description,
				Effect:      menu.Effect,
				Topmenu:     nullTopmenuType(menu.Topmenu),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Menu Command: %s: %v", menu.Name, err)
			}
		}
		return nil
	})
}
