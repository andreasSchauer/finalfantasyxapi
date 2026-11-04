package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getMix(r *http.Request, i handlerInput[seedtypes.Mix, Mix, NamedAPIResource], id int32) (Mix, error) {
	mix, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Mix{}, err
	}

	overdrive, _ := lookup.GetResourceByID(mix.OverdriveID, lookup.L.OverdrivesID)

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

func (cfg *Config) retrieveMixes(r *http.Request, i handlerInput[seedtypes.Mix, Mix, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, enums.T.MixCategory, ids, ts.QpnCategory, cfg.db.GetMixIDsByCategory),
		i.queryIDWrapper(cfg, r, ids, ts.QpnReqItem, lookup.L.Items, getMixesByItem),
	})
}
