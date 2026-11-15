package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedAgilityTiers(qtx *database.Queries, ctx context.Context) error {
	agilityTiers := dedupeRows(sh.RawJSON.AgilityTiers, sh.Hashes)

	params := database.CreateAgilityTierBulkParams{
		DataHash:        make([]string, len(agilityTiers)),
		MinAgility:      make([]int32, len(agilityTiers)),
		MaxAgility:      make([]int32, len(agilityTiers)),
		TickSpeed:       make([]int32, len(agilityTiers)),
		MonsterMinIcv:   make([]sql.NullInt32, len(agilityTiers)),
		MonsterMaxIcv:   make([]sql.NullInt32, len(agilityTiers)),
		CharacterMaxIcv: make([]sql.NullInt32, len(agilityTiers)),
	}

	for i, a := range agilityTiers {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.MinAgility[i] = a.MinAgility
		params.MaxAgility[i] = a.MaxAgility
		params.TickSpeed[i] = a.TickSpeed
		params.MonsterMinIcv[i] = h.GetNullInt32(a.MonsterMinICV)
		params.MonsterMaxIcv[i] = h.GetNullInt32(a.MonsterMaxICV)
		params.CharacterMaxIcv[i] = h.GetNullInt32(a.CharacterMaxICV)
	}

	dbRows, err := qtx.CreateAgilityTierBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create agility tiers: %v", err)
	}

	for i, row := range dbRows {
		agilityTiers[i].ID = row.ID
		sh.RawJSON.AgilityTiers[i].ID = row.ID
		lookup.L.AgilityTiersID[row.ID] = agilityTiers[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func Loop2SeedAgilitySubtiers(qtx *database.Queries, ctx context.Context) error {
	subtiers := extractAgilitySubtiers()

	params := database.CreateAgilitySubtierBulkParams{
		DataHash:        make([]string, len(subtiers)),
		AgilityTierID:   make([]int32, len(subtiers)),
		MinAgility:      make([]int32, len(subtiers)),
		MaxAgility:      make([]int32, len(subtiers)),
		CharacterMinIcv: make([]sql.NullInt32, len(subtiers)),
	}

	for i, s := range subtiers {
		params.DataHash[i] = sh.GenerateDataHash(s)
		params.AgilityTierID[i] = s.AgilityTierID
		params.MinAgility[i] = s.MinAgility
		params.MaxAgility[i] = s.MaxAgility
		params.CharacterMinIcv[i] = h.GetNullInt32(s.CharacterMinICV)
	}

	dbRows, err := qtx.CreateAgilitySubtierBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create agility subtiers: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAgilitySubtiers() []seedtypes.AgilitySubtier {
	subtiers := []seedtypes.AgilitySubtier{}

	for i := range sh.RawJSON.AgilityTiers {
		tier := &sh.RawJSON.AgilityTiers[i]

		for j := range tier.CharacterMinICVs {
			subtier := &tier.CharacterMinICVs[j]
			subtier.AgilityTierID = tier.ID
			subtiers = append(subtiers, *subtier)
		}
	}

	return dedupeRows(subtiers, sh.Hashes)
}

func completeAgilityTiers() error {
	for i := range sh.RawJSON.AgilityTiers {
		tier := &sh.RawJSON.AgilityTiers[i]

		err := assignIDs(tier.CharacterMinICVs)
		if err != nil {
			return err
		}
	}

	return nil
}
