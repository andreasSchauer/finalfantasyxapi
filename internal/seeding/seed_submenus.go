package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Submenu struct {
	ID          int32
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Effect      string   `json:"effect"`
	Topmenu     *string  `json:"topmenu"`
	Users       []string `json:"users"`
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

func (s Submenu) Error() string {
	return fmt.Sprintf("submenu %s", s.Name)
}

func (l *Lookup) seedSubmenus(db *database.Queries, dbConn *sql.DB) error {
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
				return getErr(submenu.Error(), err, "couldn't create submenu")
			}
			submenu.ID = dbSubmenu.ID

			l.submenus[submenu.Name] = submenu
		}
		return nil
	})
}

func (l *Lookup) seedSubmenusRelationships(db *database.Queries, dbConn *sql.DB) error {
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
					return getErr(submenu.Error(), err)
				}

				err = qtx.CreateSubmenusUsersJunction(context.Background(), database.CreateSubmenusUsersJunctionParams{
					DataHash:         generateDataHash(junction),
					SubmenuID:        junction.ParentID,
					CharacterClassID: junction.ChildID,
				})
				if err != nil {
					subjects := joinSubjects(submenu.Error(), jsonCharClass)
					return getErr(subjects, err, "couldn't junction user")
				}
			}
		}

		return nil
	})
}
