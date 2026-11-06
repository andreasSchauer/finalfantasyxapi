package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertSphereGrid(sg seedtypes.SphereGrid) apitypes.SphereGrid {
	return apitypes.SphereGrid{
		StatTable:  sg.StatTable,
		Lv1Locks:   sg.Lv1Locks,
		Lv2Locks:   sg.Lv2Locks,
		Lv3Locks:   sg.Lv3Locks,
		Lv4Locks:   sg.Lv4Locks,
		EmptyNodes: sg.EmptyNodes,
	}
}
