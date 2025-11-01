package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Submenu struct {
	ID	 		int32
	Name        string  	`json:"name"`
	Description string  	`json:"description"`
	Effect      string  	`json:"effect"`
	Topmenu     *string 	`json:"topmenu"`
	Users		[]string	`json:"users"`
}

func (s Submenu) ToHashFields() []any {
	return []any{
		s.Name,
		s.Description,
		s.Effect,
		derefOrNil(s.Topmenu),
	}
}

func (s Submenu) GetID() int32 {
	return s.ID
}



func (l *lookup) seedSubmenus(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/submenus.json"

	var submenus []Submenu
	err := loadJSONFile(string(srcPath), &submenus)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, submenu := range submenus {
			dbSubmenu, err := qtx.CreateSubmenu(context.Background(), database.CreateSubmenuParams{
				DataHash:    generateDataHash(submenu),
				Name:        submenu.Name,
				Description: submenu.Description,
				Effect:      submenu.Effect,
				Topmenu:     nullTopmenuType(submenu.Topmenu),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Submenu: %s: %v", submenu.Name, err)
			}
			submenu.ID = dbSubmenu.ID

			l.submenus[submenu.Name] = submenu
		}
		return nil
	})
}


func (l *lookup) seedSubmenusRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/submenus.json"

	var submenus []Submenu
	err := loadJSONFile(string(srcPath), &submenus)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonSubmenu := range submenus {
			submenu, err := l.getSubmenu(jsonSubmenu.Name)
			if err != nil {
				return err
			}

			for _, jsonCharClass := range jsonSubmenu.Users {
				junction, err := createJunction(submenu, jsonCharClass, l.getCharacterClass)
				if err != nil {
					return err
				}

				err = qtx.CreateSubmenusUsersJunction(context.Background(), database.CreateSubmenusUsersJunctionParams{
					DataHash: generateDataHash(junction),
					SubmenuID: 			junction.ParentID,
					CharacterClassID: 	junction.ChildID,
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}