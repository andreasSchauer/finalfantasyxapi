package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type unitRepl struct {
	unit        seedtypes.PlayerUnit
	bombWpn     bool
	isCelestial bool
	replVals    biReplacement
}

type biReplacement struct {
	Range          *int32
	ShatterRate    *int32
	Accuracy       *apitypes.Accuracy
	DamageConstant *int32
	DamageFormula  *string
}

func applyUser[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i HandlerInput[T, R, A], ability apitypes.UserAbility, queryName ts.QueryParamName) ([]apitypes.BattleInteraction, error) {
	repl, err := getUnitRepl(r, i, queryName)
	if errs.QueryIsEmpty(err) {
		return ability.GetBattleInteractions(), nil
	}
	if err != nil {
		return nil, err
	}

	err = verifyAbilityUsage(ability, repl, queryName)
	if err != nil {
		return nil, err
	}

	battleInteractions := applyBiReplacement(ability.GetBattleInteractions(), repl.replVals)

	return battleInteractions, nil
}

func getUnitRepl[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i HandlerInput[T, R, A], queryName ts.QueryParamName) (unitRepl, error) {
	queryParamUser := i.queryLookup[queryName]
	queryParamBomb := i.queryLookup[ts.QpnBombWpn]
	queryParamCelestialWpn := i.queryLookup[ts.QpnCelestialWeapon]

	unitID, err := queryParamUser.ParseNameID(r, E.PlayerUnits.resTypeSingle, E.PlayerUnits.objLookup)
	if err != nil {
		return unitRepl{}, err
	}
	unit, _ := lookup.GetResourceByID(unitID, lookup.L.PlayerUnitsID)

	bombWpn, err := queryParamBomb.ParseBool(r)
	if errs.ErrExceptEmptyQuery(err) {
		return unitRepl{}, err
	}

	isCelestial, err := queryParamCelestialWpn.ParseBool(r)
	if errs.ErrExceptEmptyQuery(err) {
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
		repl, err = populateReplAeon(repl, queryParamUser)
		if err != nil {
			return unitRepl{}, err
		}
	}

	return repl, nil
}

func populateReplCharacter(repl unitRepl, queryParamUser queryparam.QueryParam) (unitRepl, error) {
	charFormulas := map[string]database.DamageFormula{
		string(database.CelestialFormulaHpHigh): database.DamageFormulaClstlHpHigh,
		string(database.CelestialFormulaHpLow):  database.DamageFormulaClstlHpLow,
		string(database.CelestialFormulaMpHigh): database.DamageFormulaClstlMpHigh,
	}

	id, err := queryParamUser.CheckNameID(repl.unit.Name, ts.ResTypeSingle(repl.unit.Type), lookup.L.Characters)
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
			return unitRepl{}, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't have a celestial weapon", character), err)
		}

		formula := string(charFormulas[clstlWpn.Formula])
		repl.replVals.DamageFormula = &formula
	}

	return repl, nil
}

func populateReplAeon(repl unitRepl, queryParamUser queryparam.QueryParam) (unitRepl, error) {
	if repl.isCelestial {
		return unitRepl{}, errs.NewHTTPError(http.StatusBadRequest, "aeons don't use celestial weapons", nil)
	}

	id, err := queryParamUser.CheckNameID(repl.unit.Name, ts.ResTypeSingle(repl.unit.Type), lookup.L.Aeons)
	if err != nil {
		return unitRepl{}, err
	}
	aeon, _ := lookup.GetResourceByID(id, lookup.L.AeonsID)

	repl.replVals.Range = aeon.PhysAtkRange
	repl.replVals.ShatterRate = aeon.PhysAtkShatterRate
	repl.replVals.Accuracy = ah.ConvertObjPtr(aeon.PhysAtkAccuracy, convertAccuracy)
	repl.replVals.DamageConstant = aeon.PhysAtkDmgConstant

	return repl, nil
}

func verifyAbilityUsage(ability apitypes.UserAbility, repl unitRepl, queryName ts.QueryParamName) error {
	if !ability.CanUseAbility(repl.unit.Name) {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn %s", queryName, E.PlayerUnits.resTypeSingle, repl.unit.Name, ability), nil)
	}

	return nil
}

func applyBiReplacement(battleInteractions []apitypes.BattleInteraction, replVals biReplacement) []apitypes.BattleInteraction {
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
