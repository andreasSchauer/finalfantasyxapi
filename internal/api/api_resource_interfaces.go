package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type APIResource interface {
	HasAPIResource
	helpers.HasID
	HasURL
}

type HasAPIResource interface {
	GetAPIResource() APIResource
}

type HasURL interface {
	GetURL() string
}

type APIResourceList interface {
	getListParams() ListParams
	getResults() []HasAPIResource
}

type SimpleResource interface {
	HasURL
}

type NameAmount interface {
	GetName() string
	GetVersion() *int32
	GetVal() int32
}

type ResourceAmount interface {
	HasAPIResource
	NameAmount
}
