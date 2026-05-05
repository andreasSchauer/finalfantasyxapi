package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop3SeedPrimers(qtx *database.Queries, ctx context.Context) error {
	primers, err := l.extractPrimers()
	if err != nil {
		return err
	}

	params := database.CreatePrimerBulkParams{
		DataHash:      make([]string, len(primers)),
		KeyItemID:     make([]int32, len(primers)),
		AlBhedLetter:  make([]string, len(primers)),
		EnglishLetter: make([]string, len(primers)),
	}

	for i, p := range primers {
		params.DataHash[i] = generateDataHash(p)
		params.KeyItemID[i] = p.KeyItemID
		params.AlBhedLetter[i] = p.AlBhedLetter
		params.EnglishLetter[i] = p.EnglishLetter
	}

	dbRows, err := qtx.CreatePrimerBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create primers: %v", err)
	}

	for i, row := range dbRows {
		primers[i].ID = row.ID
		l.json.primers[i].ID = row.ID
		l.Primers[Key(primers[i])] = primers[i]
		l.PrimersID[row.ID] = primers[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractPrimers() ([]Primer, error) {
	primers := []Primer{}
	var err error

	for i := range l.json.primers {
		primer := &l.json.primers[i]

		primer.KeyItemID, err = assignFK(primer.Name, l.KeyItems)
		if err != nil {
			return nil, err
		}

		primers = append(primers, *primer)
	}

	return dedupeRows(primers, l.Hashes), nil
}
