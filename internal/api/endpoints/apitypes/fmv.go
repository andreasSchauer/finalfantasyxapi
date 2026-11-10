package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type FMV struct {
	ID                  int32              `json:"id"`
	Name                string             `json:"name"`
	Translation         *string            `json:"translation,omitempty"`
	CutsceneDescription string             `json:"cutscene_description"`
	Area                apiresource.Area   `json:"area"`
	Song                *apiresource.Named `json:"song"`
}
