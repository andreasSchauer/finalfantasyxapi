package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EncounterArea struct {
	ID            int32
	LocationArea  LocationArea  `json:"location_area"`
	AreaID        int32
	Specification *string 		`json:"specification"`
	Availability  string		`json:"availability"`
}

func (ea EncounterArea) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", ea),
		ea.AreaID,
		h.DerefOrNil(ea.Specification),
		ea.Availability,
	}
}

func (ea EncounterArea) ToKeyFields() []any {
	return []any{
		Key(ea.LocationArea),
		ea.Specification,
	}
}

func (ea EncounterArea) GetID() int32 {
	return ea.ID
}

func (ea *EncounterArea) SetID(id int32) {
	ea.ID = id
}

func (ea EncounterArea) Error() string {
	return fmt.Sprintf("encounter location with %s, specification: %s", ea.LocationArea, h.PtrToString(ea.Specification))
}

func (ea EncounterArea) GetLocationArea() LocationArea {
	return ea.LocationArea
}
