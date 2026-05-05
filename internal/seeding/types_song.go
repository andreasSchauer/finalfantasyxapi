package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Song struct {
	ID                   int32
	Name                 string            `json:"name"`
	StreamingName        *string           `json:"streaming_name"`
	InGameName           *string           `json:"in_game_name"`
	OSTName              *string           `json:"ost_name"`
	Translation          *string           `json:"translation"`
	StreamingTrackNumber *int32            `json:"streaming_track_number"`
	MusicSphereID        *int32            `json:"music_sphere_id"`
	OSTDisc              *int32            `json:"ost_disc"`
	OSTTrackNumber       *int32            `json:"ost_track_number"`
	Credits              *SongCredits      `json:"credits"`
	DurationInSeconds    int32             `json:"duration_in_seconds"`
	CanLoop              bool              `json:"can_loop"`
	SpecialUseCase       *string           `json:"special_use_case"`
	BackgroundMusic      []BackgroundMusic `json:"background_music"`
	Cues                 []Cue             `json:"cues"`
}

func (s Song) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Name,
		h.DerefOrNil(s.StreamingName),
		h.DerefOrNil(s.InGameName),
		h.DerefOrNil(s.OSTName),
		h.DerefOrNil(s.Translation),
		h.DerefOrNil(s.StreamingTrackNumber),
		h.DerefOrNil(s.MusicSphereID),
		h.DerefOrNil(s.OSTDisc),
		h.DerefOrNil(s.OSTTrackNumber),
		s.DurationInSeconds,
		s.CanLoop,
		h.DerefOrNil(s.SpecialUseCase),
		h.ObjPtrToID(s.Credits),
	}
}

func (s Song) ToKeyFields() []any {
	return []any{
		s.Name,
	}
}

func (s Song) GetID() int32 {
	return s.ID
}

func (s Song) Error() string {
	return fmt.Sprintf("song %s", s.Name)
}

func (s Song) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}

type SongCredits struct {
	ID        int32
	Composer  *string `json:"composer"`
	Arranger  *string `json:"arranger"`
	Performer *string `json:"performer"`
	Lyricist  *string `json:"lyricist"`
}

func (sc SongCredits) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", sc),
		h.DerefOrNil(sc.Composer),
		h.DerefOrNil(sc.Arranger),
		h.DerefOrNil(sc.Performer),
		h.DerefOrNil(sc.Lyricist),
	}
}

func (sc SongCredits) GetID() int32 {
	return sc.ID
}

func (sc SongCredits) Error() string {
	return fmt.Sprintf("song credits with composer: %v, arranger: %v, performer: %v, lyricist: %v", h.PtrToString(sc.Composer), h.PtrToString(sc.Arranger), h.PtrToString(sc.Performer), h.PtrToString(sc.Lyricist))
}
