package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type AutoAbilitySimple struct {
	ID             			int32       `json:"id"`
	URL            			string      `json:"url"`
	Name           			string      `json:"name"`
	RequiredItem        	*string 	`json:"required_item"`
	ActivationCondition    	string     	`json:"activation_condition"`
	Counter                	*string     `json:"counter,omitempty"`
	GradualRecovery        	*string     `json:"gradual_recovery,omitempty"`
	AutoItemUse            	[]string    `json:"auto_item_use,omitempty"`
	OnHitElement           	*string     `json:"on_hit_element,omitempty"`
	AddedElemResist        	*string     `json:"added_elem_resist,omitempty"`
	OnHitStatus            	*string     `json:"on_hit_status,omitempty"`
	AddedStatusses         	[]string    `json:"added_statusses,omitempty"`
	AddedStatusResists     	[]string    `json:"added_status_resists,omitempty"`
	AddedProperty          	*string     `json:"added_property,omitempty"`
	ConversionTo           	*string     `json:"conversion_to,omitempty"`
	StatChanges            	[]string    `json:"stat_changes,omitempty"`
	ModifierChanges        	[]string    `json:"modifier_changes,omitempty"`
}

func (a AutoAbilitySimple) GetURL() string {
	return a.URL
}


func createAutoAbilitySimple(cfg *Config, r *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := cfg.e.autoAbilities
	autoAbility, _ := seeding.GetResourceByID(id, i.objLookupID)


	autoAbilitySimple := AutoAbilitySimple{
		ID: 					autoAbility.ID,
		URL: 					createResourceURL(cfg, i.endpoint, id),
		Name: 					autoAbility.Name,
		RequiredItem: 			convertObjPtr(cfg, autoAbility.RequiredItem, convertItemAmountSimple),
		ActivationCondition: 	autoAbility.ActivationCondition,
		Counter: 				autoAbility.Counter,
		GradualRecovery: 		autoAbility.GradualRecovery,
		AutoItemUse: 			h.SliceOrNil(autoAbility.AutoItemUse),
		OnHitElement: 			autoAbility.OnHitElement,
		AddedElemResist: 		convertObjPtr(cfg, autoAbility.AddedElemResist, convertElemResistSimple),
		OnHitStatus: 			convertObjPtr(cfg, autoAbility.OnHitStatus, convertInflictedStatusSimple),
		AddedStatusses: 		h.SliceOrNil(autoAbility.AddedStatusses),
		AddedStatusResists: 	convertObjSliceOrNil(cfg, autoAbility.AddedStatusResists, convertStatusResistSimple),
		AddedProperty: 			autoAbility.AddedProperty,
		ConversionTo: 			autoAbility.ConversionTo,
		StatChanges: 			convertObjSliceOrNil(cfg, autoAbility.StatChanges, convertStatChangeSimple),
		ModifierChanges: 		convertObjSliceOrNil(cfg, autoAbility.ModifierChanges, convertModChangeSimple),
	}

	return autoAbilitySimple, nil
}


func convertElemResistSimple(_ *Config, er seeding.ElementalResist) string {
	return fmt.Sprintf("%s: %s", er.Element, er.Affinity)
}

func convertStatusResistSimple(_ *Config, sr seeding.StatusResist) string {
	if sr.Resistance == 254 {
		return fmt.Sprintf("%s (immune)", sr.StatusCondition)
	}

	return fmt.Sprintf("%s (%d", sr.StatusCondition, sr.Resistance) + "%)"
}