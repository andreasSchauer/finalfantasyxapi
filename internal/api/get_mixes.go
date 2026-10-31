package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMix(r *http.Request, i handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedAPIResourceList], id int32) (Mix, error) {
	mix, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Mix{}, err
	}

	overdrive, _ := seeding.GetResourceByID(mix.OverdriveID, cfg.L.OverdrivesID)

	combinations, err := getMixCombinations(cfg, r, mix)
	if err != nil {
		return Mix{}, err
	}

	response := Mix{
		ID:           mix.ID,
		Name:         mix.Name,
		Category:     mix.Category,
		Overdrive:    nameToNamedAPIResource(cfg, cfg.E.Overdrives, overdrive.Name, nil),
		Description:  overdrive.Description,
		Effect:       overdrive.Effect,
		Combinations: combinations,
	}

	return response, nil
}

func (cfg *Config) retrieveMixes(r *http.Request, i handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumListQuery(cfg, r, i, cfg.T.MixCategory, ids, QpnCategory, cfg.db.GetMixIDsByCategory),
		idQueryWrapper(cfg, r, i, ids, QpnReqItem, cfg.L.Items, getMixesByItem),
	})
}
