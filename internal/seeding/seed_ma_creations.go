package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ArenaCreation struct {
	ID                        int32
	SubquestID                int32
	MonsterID                 *int32
	Name                      string  `json:"name"`
	Category                  string  `json:"category"`
	RequiredArea              *string `json:"required_area"`
	RequiredSpecies           *string `json:"required_species"`
	UnderwaterOnly            bool    `json:"underwater_only"`
	CreationsUnlockedCategory *string `json:"creations_unlocked_category"`
	Amount                    int32   `json:"amount"`
}

func (a ArenaCreation) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.SubquestID,
		h.DerefOrNil(a.MonsterID),
		a.Category,
		h.DerefOrNil(a.RequiredArea),
		h.DerefOrNil(a.RequiredSpecies),
		a.UnderwaterOnly,
		h.DerefOrNil(a.CreationsUnlockedCategory),
		a.Amount,
	}
}

func (a ArenaCreation) GetID() int32 {
	return a.ID
}

func (a ArenaCreation) Error() string {
	return fmt.Sprintf("monster arena creation %s", a.Name)
}

func (a ArenaCreation) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   a.ID,
		Name: a.Name,
	}
}

func (l *Lookup) loop7SeedArenaCreations(qtx *database.Queries, ctx context.Context) error {
	creations, err := l.extractArenaCreations()
	if err != nil {
		return err
	}

	params := database.CreateMonsterArenaCreationBulkParams{
		DataHash:   				make([]string, len(creations)),
		SubquestID: 				make([]int32, len(creations)),
		Category: 					make([]database.MaCreationCategory, len(creations)),
		RequiredArea: 				make([]database.NullMaCreationArea, len(creations)),
		RequiredSpecies: 			make([]database.NullMaCreationSpecies, len(creations)),
		UnderwaterOnly: 			make([]bool, len(creations)),
		CreationsUnlockedCategory: 	make([]database.NullCreationsUnlockedCategory, len(creations)),
		Amount: 					make([]int32, len(creations)),
		MonsterID: 					make([]sql.NullInt32, len(creations)),
	}

	for i, c := range creations {
		params.DataHash[i] = generateDataHash(c)
		params.SubquestID[i] = c.SubquestID
		params.Category[i] = database.MaCreationCategory(c.Category)
		params.RequiredArea[i] = database.ToNullMaCreationArea(c.RequiredArea)
		params.RequiredSpecies[i] = database.ToNullMaCreationSpecies(c.RequiredSpecies)
		params.UnderwaterOnly[i] = c.UnderwaterOnly
		params.CreationsUnlockedCategory[i] = database.ToNullCreationsUnlockedCategory(c.CreationsUnlockedCategory)
		params.Amount[i] = c.Amount
		params.MonsterID[i] = h.GetNullInt32(c.MonsterID)
	}

	dbRows, err := qtx.CreateMonsterArenaCreationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create arena creations: %v", err)
	}

	for i, row := range dbRows {
		creations[i].ID = row.ID
		l.json.monsterArenaCreations[i].ID = row.ID
		l.ArenaCreations[creations[i].Name] = creations[i]
		l.ArenaCreationsID[row.ID] = creations[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractArenaCreations() ([]ArenaCreation, error) {
	creations := []ArenaCreation{}
	var err error

	for i := range l.json.monsterArenaCreations {
		creation := &l.json.monsterArenaCreations[i]

		creation.SubquestID, err = assignFK(creation.Name, l.Subquests)
		if err != nil {
			return nil, err
		}

		obj := LookupObject{
			Name: creation.Name,
		}

		creation.MonsterID, err = assignFKPtr(&obj, l.Monsters)
		if err != nil {
			return nil, err
		}

		creations = append(creations, *creation)
	}

	return dedupeRows(creations, l.Hashes), nil
}