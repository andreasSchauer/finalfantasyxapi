package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func SphereGrid(sg seedtypes.SphereGrid) apitypes.SphereGrid {
	return apitypes.SphereGrid{
		StatTable:  sg.StatTable,
		Lv1Locks:   sg.Lv1Locks,
		Lv2Locks:   sg.Lv2Locks,
		Lv3Locks:   sg.Lv3Locks,
		Lv4Locks:   sg.Lv4Locks,
		EmptyNodes: sg.EmptyNodes,
	}
}
