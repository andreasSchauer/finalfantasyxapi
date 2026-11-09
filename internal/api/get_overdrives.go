package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getOverdrive(r *http.Request, i handlerinput.Main[seedtypes.Overdrive, apitypes.Overdrive, apiresource.Named], id int32) (apitypes.Overdrive, error) {
	overdrive, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Overdrive{}, err
	}

	response := apitypes.Overdrive{
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
		User:               handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, overdrive.User, nil),
		OverdriveCommand:   handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		OverdriveAbilities: RefsToNamedAPIResources(overdrive.OverdriveAbilities),
	}

	return response, nil
}

func retrieveOverdrives(r *http.Request, i handlerinput.Main[seedtypes.Overdrive, apitypes.Overdrive, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryIntList(r, ids, ts.QpnRank, config.Cfg.DB.GetOverdriveIDsByRank),
		i.QueryNameID(r, ids, ts.QpnUser, apilookups.E.CharacterClasses.ResTypeSingle, lookup.L.CharClasses, handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveIDsByUser)),
	})
}
