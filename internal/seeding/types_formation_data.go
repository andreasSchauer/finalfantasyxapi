package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type FormationData struct {
	ID             int32              `json:"formation_data_id"`
	Category       string             `json:"category"`
	Availability   string             `json:"availability"`
	IsForcedAmbush bool               `json:"is_forced_ambush"`
	CanEscape      bool               `json:"can_escape"`
	BossMusic      *FormationBossSong `json:"boss_music"`
	Notes          *string            `json:"notes"`
}

func (fd FormationData) GetID() int32 {
	return fd.ID
}

func (fd FormationData) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", fd),
		fd.Category,
		fd.Availability,
		fd.IsForcedAmbush,
		fd.CanEscape,
		h.ObjPtrToID(fd.BossMusic),
		h.DerefOrNil(fd.Notes),
	}
}

func (fd FormationData) Error() string {
	return fmt.Sprintf("formation data with category: %s, forced ambush: %t, can escape: %t, boss music id: %v, notes: %v", fd.Category, fd.IsForcedAmbush, fd.CanEscape, h.ObjPtrToID(fd.BossMusic), h.PtrToString(fd.Notes))
}

type FormationBossSong struct {
	ID               int32  `json:"formation_boss_song_id"`
	SongID           int32  `json:"song_id"`
	Song             string `json:"music"`
	CelebrateVictory bool   `json:"celebrate_victory"`
}

func (s FormationBossSong) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.SongID,
		s.CelebrateVictory,
	}
}

func (s FormationBossSong) GetID() int32 {
	return s.ID
}

func (s FormationBossSong) Error() string {
	return fmt.Sprintf("formation boss song %s, celebrate victory: %t", s.Song, s.CelebrateVictory)
}
