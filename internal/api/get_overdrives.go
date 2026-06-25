package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOverdrive(r *http.Request, i handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList], id int32) (Overdrive, error) {
	overdrive, err := verifyParamsAndGet(cfg, r, i, id)
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
		User:               nameToNamedAPIResource(cfg, cfg.e.characterClasses, overdrive.User, nil),
		OverdriveCommand:   namePtrToNamedAPIResPtr(cfg, cfg.e.overdriveCommands, overdrive.OverdriveCommand, nil),
		OverdriveAbilities: refsToNamedAPIResources(cfg, overdrive.OverdriveAbilities),
	}

	return response, nil
}

func (cfg *Config) retrieveOverdrives(r *http.Request, i handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(intListQuery(cfg, r, i, ids, qpnRank, cfg.db.GetOverdriveIDsByRank)),
		fidl(nameIdQuery(r, i, ids, qpnUser, cfg.e.characterClasses.resTypeSing, cfg.l.CharClasses, ToIntManyNull(cfg.db.GetOverdriveIDsByUser))),
	})
}
