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
		OverdriveAbilities: refsToAbilityResources(cfg, overdrive.OverdriveAbilities),
	}

	return response, nil
}

func (cfg *Config) retrieveOverdrives(r *http.Request, i handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(intQueryNullable(cfg, r, i, resources, "rank", cfg.db.GetOverdriveIDsByRank)),
		frl(nameOrIdQuery(cfg, r, i, resources, "user", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetOverdriveIDsByUser)),
	})
}
