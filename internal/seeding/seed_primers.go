package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Primer struct {
	ID 			  int32
	Name          string `json:"name"`
	AlBhedLetter  string `json:"al_bhed_letter"`
	EnglishLetter string `json:"english_letter"`
	KeyItemID     int32
}

func (p Primer) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", p),
		p.AlBhedLetter,
		p.EnglishLetter,
		p.KeyItemID,
	}
}

func (p Primer) GetID() int32 {
	return p.ID
}

func (p Primer) Error() string {
	return fmt.Sprintf("primer %s", p.Name)
}

func (p Primer) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID: 	p.ID,
		Name: 	p.Name,
	}
}

func (l *Lookup) loop3SeedPrimers(qtx *database.Queries, ctx context.Context) error {
	primers, err := l.extractPrimers()
	if err != nil {
		return err
	}

	params := database.CreatePrimerBulkParams{
		DataHash:   	make([]string, len(primers)),
		KeyItemID: 		make([]int32, len(primers)),
		AlBhedLetter: 	make([]string, len(primers)),
		EnglishLetter: 	make([]string, len(primers)),
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
		l.Primers[primers[i].Name] = primers[i]
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