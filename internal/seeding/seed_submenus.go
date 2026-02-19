package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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
		h.DerefOrNil(s.Topmenu),
	}
}

func (s Submenu) GetID() int32 {
	return s.ID
}

func (s Submenu) Error() string {
	return fmt.Sprintf("submenu %s", s.Name)
}

func (s Submenu) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID: s.ID,
		Name: s.Name,
	}
}

func (l *Lookup) seedSubmenus(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/submenus.json"

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
				Topmenu:     h.NullTopmenuType(submenu.Topmenu),
			})
			if err != nil {
				return h.NewErr(submenu.Error(), err, "couldn't create submenu")
			}

			submenu.ID = dbSubmenu.ID
			l.Submenus[submenu.Name] = submenu
			l.SubmenusID[submenu.ID] = submenu
		}
		return nil
	})
}

func (l *Lookup) seedSubmenusRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/submenus.json"

	var submenus []Submenu
	err := loadJSONFile(string(srcPath), &submenus)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonSubmenu := range submenus {
			submenu, err := GetResource(jsonSubmenu.Name, l.Submenus)
			if err != nil {
				return err
			}

			for _, jsonCharClass := range jsonSubmenu.Users {
				junction, err := createJunction(submenu, jsonCharClass, l.CharClasses)
				if err != nil {
					return h.NewErr(submenu.Error(), err)
				}

				err = qtx.CreateSubmenusUsersJunction(context.Background(), database.CreateSubmenusUsersJunctionParams{
					DataHash:         generateDataHash(junction),
					SubmenuID:        junction.ParentID,
					CharacterClassID: junction.ChildID,
				})
				if err != nil {
					subjects := h.JoinErrSubjects(submenu.Error(), jsonCharClass)
					return h.NewErr(subjects, err, "couldn't junction user")
				}
			}
		}

		return nil
	})
}
