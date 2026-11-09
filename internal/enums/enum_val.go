package enums

type EnumVal struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (e EnumVal) IsZero() bool {
	return e.ID == 0
}

func (e EnumVal) GetID() int32 {
	return e.ID
}