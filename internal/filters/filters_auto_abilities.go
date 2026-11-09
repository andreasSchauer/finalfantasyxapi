package filters

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetAutoAbilitiesByMonster(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
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

func GetAutoAbilitiesByShop(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
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
