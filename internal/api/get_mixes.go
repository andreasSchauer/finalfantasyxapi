package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMix(r *http.Request, i handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedApiResourceList], id int32) (Mix, error) {
	mix, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Mix{}, err
	}

	overdrive, _ := seeding.GetResourceByID(mix.OverdriveID, cfg.l.OverdrivesID)
	combinations := mix.PossibleCombinations

	best, err := parseBooleanQuery(r, cfg.q.mixes["best"])
	if err != nil && !errors.Is(err, errEmptyQuery) {
		return Mix{}, err
	}

	if best {
		combinations = mix.BestCombinations
	}

	itemID, err := parseNameIdQuery(r, cfg.q.mixes["contains_item"], cfg.e.items.resourceType, cfg.e.items.objLookup)
	if err != nil && !errors.Is(err, errEmptyQuery) {
		return Mix{}, err
	}
	if !errors.Is(err, errEmptyQuery) {
		item, _ := seeding.GetResourceByID(itemID, cfg.l.ItemsID)

		combinations = h.Filter(combinations, func(c seeding.MixCombination) bool {
			return item.Name == c.FirstItem || item.Name == c.SecondItem
		})
	}

	response := Mix{
		ID:           mix.ID,
		Name:         mix.Name,
		Category:     newNamedAPIResourceFromEnum(cfg, cfg.e.mixCategory.endpoint, mix.Category, cfg.t.MixCategory),
		Overdrive:    nameToNamedAPIResource(cfg, cfg.e.overdrives, overdrive.Name, nil),
		Description:  overdrive.Description,
		Effect:       overdrive.Effect,
		Combinations: convertObjSlice(cfg, combinations, convertMixCombination),
	}

	return response, nil
}

func (cfg *Config) retrieveMixes(r *http.Request, i handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.MixCategory, resources, "category", cfg.db.GetMixIDsByCategory)),
		frl(nameOrIdQueryWrapper(cfg, r, i, resources, "req_item", cfg.e.items.resourceType, cfg.l.Items, getMixesByItem)),
	})
}

func getMixesByItem(cfg *Config, r *http.Request, firstItemId int32) ([]int32, error) {
	i := cfg.e.mixes

	secondItemIdPtr, err := getQueryNameIdPtr(r, cfg.e.items, "second_item", i.queryLookup)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMixIDsByItems(r.Context(), database.GetMixIDsByItemsParams{
		FirstItemID: 	firstItemId,
		SecondItemID: 	h.GetNullInt32(secondItemIdPtr),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by items.", i.resourceType), err)
	}

	return dbIDs, nil
}
