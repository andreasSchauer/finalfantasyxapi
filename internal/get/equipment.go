package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetEquipment(r *http.Request, i handlerinput.Main[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.Named], id int32) (apitypes.EquipmentName, error) {
	equipment, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.EquipmentName{}, err
	}

	rel, err := assemble.GetEquipmentRelationships(r, equipment)
	if err != nil {
		return apitypes.EquipmentName{}, err
	}

	response := apitypes.EquipmentName{
		ID:                      equipment.ID,
		Name:                    equipment.Name,
		Character:               rel.Character,
		EquipmentTable:          rel.EquipmentTable,
		Type:                    rel.Type,
		Classification:          rel.Classification,
		Priority:                rel.Priority,
		CelestialWeapon:         rel.CelestialWeapon,
		RequiredAutoAbilities:   rel.RequiredAutoAbilities,
		SelectableAutoAbilities: rel.SelectableAutoAbilities,
		RequiredSlots:           rel.RequiredSlots,
		Treasures:               rel.Treasures,
		Shops:                   rel.Shops,
	}

	return response, nil
}

func RetrieveEquipment(r *http.Request, i handlerinput.Main[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryIDList(r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, config.Cfg.DB.GetEquipmentIDsByAutoAbility),
		i.QueryNameID(r, ids, ts.QpnCharacter, apilookups.E.Characters.ResTypeSingle, apilookups.E.Characters.ObjLookup, config.Cfg.DB.GetEquipmentIDsByCharacter),
		i.QueryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetEquipmentIDsByEquipType),
		i.QueryBool2(r, ids, ts.QpnCelestialWeapon, config.Cfg.DB.GetEquipmentIDsCelestialWeapon),
	})
}
