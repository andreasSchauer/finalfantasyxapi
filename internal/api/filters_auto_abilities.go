package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getAutoAbilitiesByMonster(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.E.AutoAbilities

	charIDPtr, err := getQueryNameIDPtr(r, cfg.E.Characters, QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAutoAbilityIDsByMonster(ctx, database.GetAutoAbilityIDsByMonsterParams{
		MonsterID:   id,
		CharacterID: h.GetNullInt32(charIDPtr),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, i.queryLookup[QpnMonster], err)
	}

	return dbIDs, nil
}

func getAutoAbilitiesByShop(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.E.AutoAbilities

	charIDPtr, err := getQueryNameIDPtr(r, cfg.E.Characters, QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAutoAbilityIDsByShop(ctx, database.GetAutoAbilityIDsByShopParams{
		ShopID:      id,
		CharacterID: h.GetNullInt32(charIDPtr),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, i.queryLookup[QpnMonster], err)
	}

	return dbIDs, nil
}
