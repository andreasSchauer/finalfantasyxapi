package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAutoAbilitiesByMonster(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.AutoAbilities

	charIDPtr, err := nullNameID(r, E.Characters, ts.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetAutoAbilityIDsByMonster(ctx, database.GetAutoAbilityIDsByMonsterParams{
		MonsterID:   id,
		CharacterID: charIDPtr,
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, i.queryLookup[ts.QpnMonster], err)
	}

	return dbIDs, nil
}

func getAutoAbilitiesByShop(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.AutoAbilities

	charIDPtr, err := nullNameID(r, E.Characters, ts.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetAutoAbilityIDsByShop(ctx, database.GetAutoAbilityIDsByShopParams{
		ShopID:      id,
		CharacterID: charIDPtr,
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, i.queryLookup[ts.QpnMonster], err)
	}

	return dbIDs, nil
}
