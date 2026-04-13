package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMix(r *http.Request, i handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedApiResourceList], id int32) (Mix, error) {
	mix, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Mix{}, err
	}

	overdrive, _ := seeding.GetResourceByID(mix.OverdriveID, cfg.l.OverdrivesID)

	combinations, err := getMixCombinations(cfg, r, mix)
	if err != nil {
		return Mix{}, err
	}

	response := Mix{
		ID:           mix.ID,
		Name:         mix.Name,
		Category:     enumToNamedAPIResource(cfg, cfg.e.mixCategory.endpoint, mix.Category, cfg.t.MixCategory),
		Overdrive:    nameToNamedAPIResource(cfg, cfg.e.overdrives, overdrive.Name, nil),
		Description:  overdrive.Description,
		Effect:       overdrive.Effect,
		Combinations: combinations,
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
		frl(idQueryWrapper(cfg, r, i, resources, "req_item", len(cfg.l.Items), getMixesByItem)),
	})
}
