package sectiontypes

type AeonStatSimple struct {
	ID    int32         `json:"id"`
	URL   string        `json:"url"`
	Name  string        `json:"name"`
	AVals AeonStatTable `json:"a_vals"`
	BVals AeonStatTable `json:"b_vals"`
	XVals []Xval        `json:"x_vals"`
}

func (a AeonStatSimple) GetURL() string {
	return a.URL
}

type Xval struct {
	Battles int32         `json:"battles"`
	Stats   AeonStatTable `json:"stats"`
}

type AeonStatTable struct {
	HP           int32  `json:"hp"`
	MP           int32  `json:"mp"`
	Strength     int32  `json:"strength"`
	Defense      int32  `json:"defense"`
	Magic        int32  `json:"magic"`
	MagicDefense int32  `json:"magic_defense"`
	Agility      int32  `json:"agility"`
	Luck         *int32 `json:"luck,omitempty"`
	Evasion      int32  `json:"evasion"`
	Accuracy     int32  `json:"accuracy"`
}
