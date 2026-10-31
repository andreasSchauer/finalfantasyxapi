package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOverdrive(r *http.Request, i handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedAPIResourceList], id int32) (Overdrive, error) {
	overdrive, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveOverdrives(r *http.Request, i handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		intListQuery(cfg, r, i, ids, QpnRank, cfg.db.GetOverdriveIDsByRank),
		nameIDQuery(r, i, ids, QpnUser, cfg.E.CharacterClasses.resTypeSingle, cfg.L.CharClasses, ToIntManyNull(cfg.db.GetOverdriveIDsByUser)),
	})
}
