package seeding

import (
	"context"
	"fmt"
	
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop1SeedSphereGrids(qtx *database.Queries, ctx context.Context) error {
	grids := l.extractSphereGrids()

	params := database.CreateSphereGridBulkParams{
		DataHash:       make([]string, len(grids)),
		Type: 			make([]database.SphereGridType, len(grids)),
		Hp: 			make([]int32, len(grids)),
		Mp: 			make([]int32, len(grids)),
		Strength: 		make([]int32, len(grids)),
		Defense: 		make([]int32, len(grids)),
		Magic: 			make([]int32, len(grids)),
		MagicDefense: 	make([]int32, len(grids)),
		Agility: 		make([]int32, len(grids)),
		Luck: 			make([]int32, len(grids)),
		Evasion: 		make([]int32, len(grids)),
		Accuracy: 		make([]int32, len(grids)),
		Lv1Locks: 		make([]int32, len(grids)),
		Lv2Locks: 		make([]int32, len(grids)),
		Lv3Locks: 		make([]int32, len(grids)),
		Lv4Locks: 		make([]int32, len(grids)),
		EmptyNodes: 	make([]int32, len(grids)),
	}

	for i, g := range grids {
		params.DataHash[i] = generateDataHash(g)
		params.Type[i] = g.Type
		params.Hp[i] = g.HP
		params.Mp[i] = g.MP
		params.Strength[i] = g.Strength
		params.Defense[i] = g.Defense
		params.Magic[i] = g.Magic
		params.MagicDefense[i] = g.MagicDefense
		params.Agility[i] = g.Agility
		params.Luck[i] = g.Luck
		params.Evasion[i] = g.Evasion
		params.Accuracy[i] = g.Accuracy
		params.Lv1Locks[i] = g.Lv1Locks
		params.Lv2Locks[i] = g.Lv2Locks
		params.Lv3Locks[i] = g.Lv3Locks
		params.Lv4Locks[i] = g.Lv4Locks
		params.EmptyNodes[i] = g.EmptyNodes
	}

	dbRows, err := qtx.CreateSphereGridBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create sphere grid: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSphereGrids() []SphereGrid {
	grids := []SphereGrid{}

	for i := range l.json.characters {
		char := &l.json.characters[i]

		if char.IsStoryBased {
			continue
		}

		char.StdSphereGrid.Type = database.SphereGridTypeStandard
		char.ExpSphereGrid.Type = database.SphereGridTypeExpert
		grids = append(grids, *char.StdSphereGrid)
		grids = append(grids, *char.ExpSphereGrid)
	}

	return dedupeRows(grids, l.Hashes)
}