package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAeonCommand(r *http.Request, i handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList], id int32) (AeonCommand, error) {
	command, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return AeonCommand{}, err
	}

	var userName string

	if len(command.PossibleAbilities) == 3 {
		userName = "magus-sisters"
	} else {
		userName = command.PossibleAbilities[0].User
	}

	response := AeonCommand{
		ID:                command.ID,
		Name:              command.Name,
		Description:       command.Description,
		Effect:            command.Effect,
		Cursor:            command.Cursor,
		User:              nameToNamedAPIResource(cfg, cfg.e.characterClasses, userName, nil),
		Topmenu:           namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, command.Topmenu, nil),
		OpenSubmenu:       namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, command.OpenSubmenu, nil),
		PossibleAbilities: convertObjSlice(cfg, command.PossibleAbilities, convertPossibleAbilityList),
	}

	return response, nil
}

func (cfg *Config) retrieveAeonCommands(r *http.Request, i handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	return verifyParamsAndRetrieve(r, i)
}
