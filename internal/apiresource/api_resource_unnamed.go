package apiresource

import (
	"fmt"
)

type Unnamed struct {
	ID  int32  `json:"id"`
	URL string `json:"url"`
}

func (r Unnamed) IsZero() bool {
	return r.ID == 0
}

func (r Unnamed) GetID() int32 {
	return r.ID
}

func (r Unnamed) GetURL() string {
	return r.URL
}

func (r Unnamed) GetTestKey() string {
	return fmt.Sprint(r.ID)
}

func (r Unnamed) String() string {
	return fmt.Sprintf("unnamed api resource with url: %s", r.URL)
}

func (r Unnamed) GetAPIResource() APIResource {
	return r
}
