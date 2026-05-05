package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop4SeedStats(qtx *database.Queries, ctx context.Context) error {
	stats, err := l.extractStats()
	if err != nil {
		return err
	}

	params := database.CreateStatBulkParams{
		DataHash: make([]string, len(stats)),
		Name:     make([]string, len(stats)),
		Effect:   make([]string, len(stats)),
		MinVal:   make([]int32, len(stats)),
		MaxVal:   make([]int32, len(stats)),
		MaxVal2:  make([]sql.NullInt32, len(stats)),
		SphereID: make([]sql.NullInt32, len(stats)),
	}

	for i, s := range stats {
		params.DataHash[i] = generateDataHash(s)
		params.Name[i] = s.Name
		params.Effect[i] = s.Effect
		params.MinVal[i] = s.MinVal
		params.MaxVal[i] = s.MaxVal
		params.MaxVal2[i] = h.GetNullInt32(s.MaxVal2)
		params.SphereID[i] = h.GetNullInt32(s.SphereID)
	}

	dbRows, err := qtx.CreateStatBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create stats: %v", err)
	}

	for i, row := range dbRows {
		stats[i].ID = row.ID
		l.json.stats[i].ID = row.ID
		l.Stats[Key(stats[i])] = stats[i]
		l.StatsID[row.ID] = stats[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractStats() ([]Stat, error) {
	stats := []Stat{}
	var err error

	for i := range l.json.stats {
		stat := &l.json.stats[i]

		stat.SphereID, err = assignFKPtr(&stat.ActivationSphere, l.Spheres)
		if err != nil {
			return nil, err
		}

		stats = append(stats, *stat)
	}

	return dedupeRows(stats, l.Hashes), nil
}
