package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func createMonsterSimple(r *http.Request, id int32, subsection handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Monsters
	mon, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	areaIDs := subsection.Relations[id][handlerinput.RelationAreas]

	monSimple := sectiontypes.MonsterSimple{
		ID:             mon.ID,
		URL:            ah.CreateResourceURL(i.Endpoint, id),
		Name:           mon.Name,
		Version:        mon.Version,
		Specification:  mon.Specification,
		HP:             getMonsterSimpleHP(mon),
		AP:             getMonsterSimpleAP(mon),
		Gil:            mon.Gil,
		MaxBribeAmount: getMonsterSimpleBribeAmount(mon, getMonsterHP(mon)),
		Availability:   mon.Availability,
		IsRepeatable:   mon.IsRepeatable,
		CanBeCaptured:  mon.CanBeCaptured,
		RonsoRages:     h.SliceOrNil(mon.RonsoRages),
		Areas:          ah.ConvertObjSlice(areaIDs, idToLocAreaString),
		Items:          ah.ConvertObjPtr(mon.Items, convertMonsterItemsSimple),
		Equipment:      ah.ConvertObjPtr(mon.Equipment, convertMonsterEquipmentSimple),
	}

	return monSimple, nil
}

func getMonsterHP(mon seedtypes.Monster) int32 {
	for _, stat := range mon.BaseStats {
		if stat.StatName == "hp" {
			return stat.Value
		}
	}

	return 0
}

func getMonsterSimpleHP(mon seedtypes.Monster) string {
	hp := getMonsterHP(mon)
	return fmt.Sprintf("%d (%d)", hp, mon.OverkillDamage)
}

func getMonsterSimpleAP(mon seedtypes.Monster) string {
	return fmt.Sprintf("%d (%d)", mon.AP, mon.APOverkill)
}

func getMonsterSimpleBribeAmount(mon seedtypes.Monster, hp int32) *int32 {
	if mon.Items == nil || mon.Items.Bribe == nil {
		return nil
	}

	bribeAmount := hp * 25
	return &bribeAmount
}

func getMonsterSectionRelations(r *http.Request, monIDs []int32) (map[int32]map[handlerinput.Relation][]int32, error) {
	i := apilookups.E.Monsters
	relations := make(map[int32]map[handlerinput.Relation][]int32)

	monsterJunctions, err := i.GetDBJunctions(r.Context(), monIDs, apilookups.E.Areas.ResTypeSingle, config.Cfg.DB.GetMonsterAreaIdPairs, juncMonsterArea)
	if err != nil {
		return nil, err
	}

	for _, monID := range monIDs {
		relationMap := make(map[handlerinput.Relation][]int32)

		relationMap[handlerinput.RelationAreas], monsterJunctions = GetJunctionIDs(monID, monsterJunctions)

		relations[monID] = relationMap
	}

	return relations, nil
}

func convertMonsterItemsSimple(items seedtypes.MonsterItems) sectiontypes.MonsterItemsSimple {
	return sectiontypes.MonsterItemsSimple{
		StealCommon:         ah.ConvertObjPtr(items.StealCommon, convertItemAmountSimple),
		StealRare:           ah.ConvertObjPtr(items.StealRare, convertItemAmountSimple),
		DropCommon:          ah.ConvertObjPtr(items.DropCommon, convertItemAmountSimple),
		DropRare:            ah.ConvertObjPtr(items.DropRare, convertItemAmountSimple),
		SecondaryDropCommon: ah.ConvertObjPtr(items.SecondaryDropCommon, convertItemAmountSimple),
		SecondaryDropRare:   ah.ConvertObjPtr(items.SecondaryDropRare, convertItemAmountSimple),
		Bribe:               ah.ConvertObjPtr(items.Bribe, convertItemAmountSimple),
		OtherItems:          ah.ConvertObjSliceOrNil(items.OtherItems, posItemToItemAmtSimple),
	}
}

func convertMonsterEquipmentSimple(equipment seedtypes.MonsterEquipment) sectiontypes.MonsterEquipmentSimple {
	return sectiontypes.MonsterEquipmentSimple{
		WeaponAbilities: ah.ConvertObjSlice(equipment.WeaponAbilities, monsterAutoAbilityString),
		ArmorAbilities:  ah.ConvertObjSlice(equipment.ArmorAbilities, monsterAutoAbilityString),
	}
}
