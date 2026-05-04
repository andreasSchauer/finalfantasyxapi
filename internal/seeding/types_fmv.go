package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type FMV struct {
	ID                  int32
	Name                string       `json:"name"`
	Translation         *string      `json:"translation"`
	CutsceneDescription string       `json:"cutscene_description"`
	SongName            *string      `json:"music"`
	LocationArea        LocationArea `json:"location_area"`
	SongID              *int32
	AreaID              int32
}

func (f FMV) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", f),
		f.Name,
		h.DerefOrNil(f.Translation),
		f.CutsceneDescription,
		h.DerefOrNil(f.SongName),
		h.DerefOrNil(f.SongID),
		f.AreaID,
	}
}

func (f FMV) GetID() int32 {
	return f.ID
}

func (f FMV) Error() string {
	return fmt.Sprintf("fmv %s", f.Name)
}

func (f FMV) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   f.ID,
		Name: f.Name,
	}
}
