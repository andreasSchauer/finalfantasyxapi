package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type BaseStat struct {
	Stat  apiresource.NamedAPIResource `json:"stat"`
	Value int32                        `json:"value"`
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

func (bs BaseStat) ToResAmount() apiresource.ResourceAmount[apiresource.NamedAPIResource] {
	return apiresource.ResourceAmount[apiresource.NamedAPIResource]{
		Resource: bs.Stat,
		Amount:   bs.Value,
	}
}