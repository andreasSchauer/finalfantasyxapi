package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"



type BlitzballPrize struct {
	ID       int32          `json:"id"`
	Category string         `json:"category"`
	Slot     string         `json:"slot"`
	Items    []PossibleItem `json:"items"`
}

func convertBlitzballItem(cfg *Config, bi seeding.BlitzballItem) PossibleItem {
	return convertPossibleItem(cfg, bi.PossibleItem)
}