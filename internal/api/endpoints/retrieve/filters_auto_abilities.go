package retrieve

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func getAutoAbilitiesByMonster(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.AutoAbilities

	charIDPtr, err := handlerinput.NullNameID(r, apilookups.E.Characters, ts.QpnCharacter, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetAutoAbilityIDsByMonster(ctx, database.GetAutoAbilityIDsByMonsterParams{
		MonsterID:   id,
		CharacterID: charIDPtr,
	})
	if err != nil {
		return nil, i.ErrorDBFilter(i.QueryLookup[ts.QpnMonster], err)
	}

	return dbIDs, nil
}

func getAutoAbilitiesByShop(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.AutoAbilities

	charIDPtr, err := handlerinput.NullNameID(r, apilookups.E.Characters, ts.QpnCharacter, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetAutoAbilityIDsByShop(ctx, database.GetAutoAbilityIDsByShopParams{
		ShopID:      id,
		CharacterID: charIDPtr,
	})
	if err != nil {
		return nil, i.ErrorDBFilter(i.QueryLookup[ts.QpnMonster], err)
	}

	return dbIDs, nil
}
