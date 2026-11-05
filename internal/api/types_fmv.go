package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type FMV struct {
	ID                  int32                         `json:"id"`
	Name                string                        `json:"name"`
	Translation         *string                       `json:"translation,omitempty"`
	CutsceneDescription string                        `json:"cutscene_description"`
	Area                apiresource.AreaAPIResource   `json:"area"`
	Song                *apiresource.NamedAPIResource `json:"song"`
}
