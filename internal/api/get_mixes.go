package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getMix(r *http.Request, i handlerInput[seedtypes.Mix, Mix, NamedAPIResource], id int32) (Mix, error) {
	mix, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Mix{}, err
	}

	overdrive, _ := lookup.GetResourceByID(mix.OverdriveID, lookup.L.OverdrivesID)

	combinations, err := getMixCombinations(srv, r, mix)
	if err != nil {
		return Mix{}, err
	}

	response := Mix{
		ID:           mix.ID,
		Name:         mix.Name,
		Category:     mix.Category,
		Overdrive:    nameToNamedAPIResource(srv, E.Overdrives, overdrive.Name, nil),
		Description:  overdrive.Description,
		Effect:       overdrive.Effect,
		Combinations: combinations,
	}

	return response, nil
}

func (srv *Server) retrieveMixes(r *http.Request, i handlerInput[seedtypes.Mix, Mix, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.MixCategory, ids, ts.QpnCategory, srv.Cfg.DB.GetMixIDsByCategory),
		i.queryIDWrapper(srv, r, ids, ts.QpnReqItem, lookup.L.Items, getMixesByItem),
	})
}
