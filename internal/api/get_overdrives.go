package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getOverdrive(r *http.Request, i handlerInput[seedtypes.Overdrive, Overdrive, NamedAPIResource], id int32) (Overdrive, error) {
	overdrive, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Overdrive{}, err
	}

	response := Overdrive{
		ID:                 overdrive.ID,
		Name:               overdrive.Name,
		Version:            overdrive.Version,
		Specification:      overdrive.Specification,
		Description:        overdrive.Description,
		Effect:             overdrive.Effect,
		Rank:               overdrive.Rank,
		AppearsInHelpBar:   overdrive.AppearsInHelpBar,
		CanCopycat:         overdrive.CanCopycat,
		UnlockCondition:    overdrive.UnlockCondition,
		CountdownInSec:     overdrive.CountdownInSec,
		Cursor:             overdrive.Cursor,
		User:               nameToNamedAPIResource(cfg, cfg.E.CharacterClasses, overdrive.User, nil),
		OverdriveCommand:   namePtrToNamedAPIResPtr(cfg, cfg.E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		OverdriveAbilities: refsToNamedAPIResources(cfg, overdrive.OverdriveAbilities),
	}

	return response, nil
}

func (cfg *Config) retrieveOverdrives(r *http.Request, i handlerInput[seedtypes.Overdrive, Overdrive, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryIntList(cfg, r, ids, ts.QpnRank, cfg.db.GetOverdriveIDsByRank),
		i.queryNameID(r, ids, ts.QpnUser, cfg.E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, ToIntManyNull(cfg.db.GetOverdriveIDsByUser)),
	})
}
