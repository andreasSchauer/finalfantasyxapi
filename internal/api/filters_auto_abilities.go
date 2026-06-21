package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func getAutoAbilitiesByMonster(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.autoAbilities

	charIdPtr, err := getQueryNameIdPtr(r, cfg.e.characters, "character", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAutoAbilityIDsByMonster(r.Context(), database.GetAutoAbilityIDsByMonsterParams{
		MonsterID:   id,
		CharacterID: h.GetNullInt32(charIdPtr),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["monster"], err)
	}

	return dbIDs, nil
}

func getAutoAbilitiesByShop(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.autoAbilities

	charIdPtr, err := getQueryNameIdPtr(r, cfg.e.characters, "character", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAutoAbilityIDsByShop(r.Context(), database.GetAutoAbilityIDsByShopParams{
		ShopID: 	 id,
		CharacterID: h.GetNullInt32(charIdPtr),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["monster"], err)
	}

	return dbIDs, nil
}