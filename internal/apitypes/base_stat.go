package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type BaseStat struct {
	Stat  apiresource.Named `json:"stat"`
	Value int32             `json:"value"`
}

func (bs BaseStat) GetAPIResource() apiresource.APIResource {
	return bs.Stat
}

func (bs BaseStat) GetName() string {
	return bs.Stat.Name
}

func (bs BaseStat) GetVersion() *int32 {
	return nil
}

func (bs BaseStat) GetVal() int32 {
	return bs.Value
}

func (bs BaseStat) ToResAmount() apiresource.ResourceAmount[apiresource.Named] {
	return apiresource.ResourceAmount[apiresource.Named]{
		Resource: bs.Stat,
		Amount:   bs.Value,
	}
}
