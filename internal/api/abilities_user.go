package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type userAbility interface {
	canUseAbility(*Config, string) bool
	getBattleInteractions() []BattleInteraction
	fmt.Stringer
}

type unitRepl struct {
	unit        seedtypes.PlayerUnit
	bombWpn     bool
	isCelestial bool
	replVals    biReplacement
}

type biReplacement struct {
	Range          *int32
	ShatterRate    *int32
	Accuracy       *Accuracy
	DamageConstant *int32
	DamageFormula  *string
}

func applyUser[T seedtypes.Lookupable, R any, A APIResource](cfg *Config, r *http.Request, i handlerInput[T, R, A], ability userAbility, queryName typestrings.QueryParamName) ([]BattleInteraction, error) {
	repl, err := getUnitRepl(cfg, r, i, queryName)
	if queryIsEmpty(err) {
		return ability.getBattleInteractions(), nil
	}
	if err != nil {
		return nil, err
	}

	err = verifyAbilityUsage(cfg, ability, repl, queryName)
	if err != nil {
		return nil, err
	}

	battleInteractions := applyBiReplacement(ability.getBattleInteractions(), repl.replVals)

	return battleInteractions, nil
}

func getUnitRepl[T seedtypes.Lookupable, R any, A APIResource](cfg *Config, r *http.Request, i handlerInput[T, R, A], queryName typestrings.QueryParamName) (unitRepl, error) {
	queryParamUser := i.queryLookup[queryName]
	queryParamBomb := i.queryLookup[typestrings.QpnBombWpn]
	queryParamCelestialWpn := i.queryLookup[typestrings.QpnCelestialWeapon]

	unitID, err := queryParamUser.parseNameID(r, cfg.E.PlayerUnits.resTypeSingle, cfg.E.PlayerUnits.objLookup)
	if err != nil {
		return unitRepl{}, err
	}
	unit, _ := lookup.GetResourceByID(unitID, lookup.L.PlayerUnitsID)

	bombWpn, err := queryParamBomb.parseBool(r)
	if errExceptEmptyQuery(err) {
		return unitRepl{}, err
	}

	isCelestial, err := queryParamCelestialWpn.parseBool(r)
	if errExceptEmptyQuery(err) {
		return unitRepl{}, err
	}

	repl := unitRepl{
		unit:        unit,
		bombWpn:     bombWpn,
		isCelestial: isCelestial,
		replVals:    biReplacement{},
	}

	switch repl.unit.Type {
	case database.UnitTypeCharacter:
		repl, err = populateReplCharacter(repl, queryParamUser)
		if err != nil {
			return unitRepl{}, err
		}

	case database.UnitTypeAeon:
		repl, err = populateReplAeon(cfg, repl, queryParamUser)
		if err != nil {
			return unitRepl{}, err
		}
	}

	return repl, nil
}

func populateReplCharacter(repl unitRepl, queryParamUser QueryParam) (unitRepl, error) {
	charFormulas := map[string]database.DamageFormula{
		string(database.CelestialFormulaHpHigh): database.DamageFormulaClstlHpHigh,
		string(database.CelestialFormulaHpLow):  database.DamageFormulaClstlHpLow,
		string(database.CelestialFormulaMpHigh): database.DamageFormulaClstlMpHigh,
	}

	id, err := queryParamUser.checkNameID(repl.unit.Name, typestrings.ResTypeSingle(repl.unit.Type), lookup.L.Characters)
	if err != nil {
		return unitRepl{}, err
	}
	character, _ := lookup.GetResourceByID(id, lookup.L.CharactersID)

	repl.replVals.Range = &character.PhysAtkRange

	if repl.bombWpn {
		repl.replVals.DamageConstant = h.GetInt32Ptr(18)
	}

	if repl.isCelestial {
		clstlWpn, err := lookup.GetResourceByID(id, lookup.L.CelestialWeaponsID)
		if err != nil {
			return unitRepl{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't have a celestial weapon", character), err)
		}

		formula := string(charFormulas[clstlWpn.Formula])
		repl.replVals.DamageFormula = &formula
	}

	return repl, nil
}

func populateReplAeon(cfg *Config, repl unitRepl, queryParamUser QueryParam) (unitRepl, error) {
	if repl.isCelestial {
		return unitRepl{}, newHTTPError(http.StatusBadRequest, "aeons don't use celestial weapons", nil)
	}

	id, err := queryParamUser.checkNameID(repl.unit.Name, typestrings.ResTypeSingle(repl.unit.Type), lookup.L.Aeons)
	if err != nil {
		return unitRepl{}, err
	}
	aeon, _ := lookup.GetResourceByID(id, lookup.L.AeonsID)

	repl.replVals.Range = aeon.PhysAtkRange
	repl.replVals.ShatterRate = aeon.PhysAtkShatterRate
	repl.replVals.Accuracy = convertObjPtr(cfg, aeon.PhysAtkAccuracy, convertAccuracy)
	repl.replVals.DamageConstant = aeon.PhysAtkDmgConstant

	return repl, nil
}

func verifyAbilityUsage(cfg *Config, ability userAbility, repl unitRepl, queryName typestrings.QueryParamName) error {
	if !ability.canUseAbility(cfg, repl.unit.Name) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn %s", queryName, cfg.E.PlayerUnits.resTypeSingle, repl.unit.Name, ability), nil)
	}

	return nil
}

func applyBiReplacement(battleInteractions []BattleInteraction, replVals biReplacement) []BattleInteraction {
	for i, battleInteraction := range battleInteractions {
		if !battleInteraction.BasedOnUserAttack {
			continue
		}

		if replVals.Range != nil {
			battleInteraction.Range = replVals.Range
		}

		if replVals.ShatterRate != nil {
			battleInteraction.ShatterRate = *replVals.ShatterRate
		}

		if replVals.Accuracy != nil {
			battleInteraction.Accuracy = *replVals.Accuracy
		}

		if replVals.DamageConstant != nil {
			battleInteraction.Damage.HP.DamageConstant = *replVals.DamageConstant
		}

		if replVals.DamageFormula != nil {
			battleInteraction.Damage.HP.DamageFormula = *replVals.DamageFormula
		}

		battleInteractions[i] = battleInteraction
	}

	return battleInteractions
}
