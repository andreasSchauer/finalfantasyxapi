package api


type FMV struct {
	ID                  int32             `json:"id"`
	Name                string            `json:"name"`
	Translation         *string           `json:"translation,omitempty"`
	CutsceneDescription string            `json:"cutscene_description"`
	Area                AreaAPIResource   `json:"area"`
	Song                *NamedAPIResource `json:"song"`
}