package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAutoAbilitiesByMonster(server *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.AutoAbilities

	charIDPtr, err := nullNameID(r, E.Characters, ts.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetAutoAbilityIDsByMonster(ctx, database.GetAutoAbilityIDsByMonsterParams{
		MonsterID:   id,
		CharacterID: charIDPtr,
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, i.queryLookup[ts.QpnMonster], err)
	}

	return dbIDs, nil
}

func getAutoAbilitiesByShop(server *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.AutoAbilities

	charIDPtr, err := nullNameID(r, E.Characters, ts.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetAutoAbilityIDsByShop(ctx, database.GetAutoAbilityIDsByShopParams{
		ShopID:      id,
		CharacterID: charIDPtr,
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, i.queryLookup[ts.QpnMonster], err)
	}

	return dbIDs, nil
}
