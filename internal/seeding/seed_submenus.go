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
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Effect      string  `json:"effect"`
	Topmenu     *string `json:"topmenu"`
	TopmenuID   *int32
	Users       []string `json:"users"`
}

func (s Submenu) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
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
		ID:   s.ID,
		Name: s.Name,
	}
}

func (l *Lookup) loop2SeedSubmenus(qtx *database.Queries, ctx context.Context) error {
	submenus, err := l.extractSubmenus()
	if err != nil {
		return err
	}

	params := database.CreateSubmenuBulkParams{
		DataHash:    make([]string, len(submenus)),
		Name:        make([]string, len(submenus)),
		TopmenuID:   make([]sql.NullInt32, len(submenus)),
		Description: make([]sql.NullString, len(submenus)),
		Effect:      make([]string, len(submenus)),
	}

	for i, s := range submenus {
		params.DataHash[i] = generateDataHash(s)
		params.Name[i] = s.Name
		params.TopmenuID[i] = h.GetNullInt32(s.TopmenuID)
		params.Description[i] = h.GetNullString(s.Description)
		params.Effect[i] = s.Effect
	}

	dbRows, err := qtx.CreateSubmenuBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create submenus: %v", err)
	}

	for i, row := range dbRows {
		submenus[i].ID = row.ID
		l.json.submenus[i].ID = row.ID
		l.Submenus[submenus[i].Name] = submenus[i]
		l.SubmenusID[row.ID] = submenus[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSubmenus() ([]Submenu, error) {
	submenus := []Submenu{}
	var err error

	for i := range l.json.submenus {
		submenu := &l.json.submenus[i]

		if submenu.Topmenu != nil {
			submenu.TopmenuID, err = assignFKPtr(submenu.Topmenu, l.Topmenus)
			if err != nil {
				return nil, err
			}
		}
		submenus = append(submenus, *submenu)
	}

	return dedupeRows(submenus, l.Hashes), nil
}

func (l *Lookup) getSubmenuUsers(s Submenu) ([]CharacterClass, error) {
	return getResources(s.Users, l.CharClasses)
}

func (l *Lookup) seedJuncSubmenusUsers(qtx *database.Queries, ctx context.Context) error {
	const desc string = "submenus + users"
	jParams, err := processJunctions(l, desc, l.json.submenus, l.getSubmenuUsers)
	if err != nil {
		return err
	}

	return qtx.CreateSubmenusUsersJunctionBulk(ctx, database.CreateSubmenusUsersJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		SubmenuID:        jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
