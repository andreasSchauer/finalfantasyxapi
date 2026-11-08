package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type OverdriveMode struct {
	ID          int32                                           `json:"id"`
	Name        string                                          `json:"name"`
	Description string                                          `json:"description"`
	Effect      string                                          `json:"effect"`
	Type        string                                          `json:"type"`
	FillRate    *float32                                        `json:"fill_rate,omitempty"`
	Actions     []apiresource.ResourceAmount[apiresource.Named] `json:"actions"`
}
