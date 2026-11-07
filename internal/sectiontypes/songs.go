package sectiontypes

type SongSimple struct {
	ID                int32   `json:"id"`
	URL               string  `json:"url"`
	Name              string  `json:"name"`
	Composer          *string `json:"composer,omitempty"`
	Arranger          *string `json:"arranger,omitempty"`
	Performer         *string `json:"performer,omitempty"`
	Lyricist          *string `json:"lyricist,omitempty"`
	DurationInSeconds int32   `json:"duration_in_seconds"`
	CanLoop           bool    `json:"can_loop"`
}

func (s SongSimple) GetURL() string {
	return s.URL
}
