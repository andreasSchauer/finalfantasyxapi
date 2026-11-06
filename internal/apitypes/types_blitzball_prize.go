package apitypes

type BlitzballPrize struct {
	ID       int32          `json:"id"`
	Name     string         `json:"name"`
	Category string         `json:"category"`
	Slot     string         `json:"slot"`
	Items    []PossibleItem `json:"items"`
}