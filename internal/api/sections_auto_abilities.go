package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func createAutoAbilitySimple(r *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.AutoAbilities
	autoAbility, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	autoAbilitySimple := sectiontypes.AutoAbilitySimple{
		ID:                  autoAbility.ID,
		URL:                 ah.CreateResourceURL(i.Endpoint, id),
		Name:                autoAbility.Name,
		RequiredItem:        ah.ConvertObjPtr(autoAbility.RequiredItem, convertItemAmountSimple),
		ActivationCondition: autoAbility.ActivationCondition,
		Counter:             autoAbility.Counter,
		GradualRecovery:     autoAbility.GradualRecovery,
		AutoItemUse:         h.SliceOrNil(autoAbility.AutoItemUse),
		OnHitElement:        autoAbility.OnHitElement,
		AddedElemResist:     ah.ConvertObjPtr(autoAbility.AddedElemResist, convertElemResistSimple),
		OnHitStatus:         ah.ConvertObjPtr(autoAbility.OnHitStatus, convertInflictedStatusSimple),
		AddedStatusses:      h.SliceOrNil(autoAbility.AddedStatusses),
		AddedStatusResists:  ah.ConvertObjSliceOrNil(autoAbility.AddedStatusResists, convertStatusResistSimple),
		AddedProperty:       autoAbility.AddedProperty,
		ConversionTo:        autoAbility.ConversionTo,
		StatChanges:         ah.ConvertObjSliceOrNil(autoAbility.StatChanges, convertStatChangeSimple),
		ModifierChanges:     ah.ConvertObjSliceOrNil(autoAbility.ModifierChanges, convertModChangeSimple),
	}

	return autoAbilitySimple, nil
}

func convertElemResistSimple(er seedtypes.ElementalResist) string {
	return fmt.Sprintf("%s: %s", er.Element, er.Affinity)
}

func convertStatusResistSimple(sr seedtypes.StatusResist) string {
	if sr.Resistance == 254 {
		return fmt.Sprintf("%s (immune)", sr.StatusCondition)
	}

	return fmt.Sprintf("%s (%d", sr.StatusCondition, sr.Resistance) + "%)"
}
