package main

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type APIResource interface {
	HasAPIResource
	h.HasID
	seeding.Lookupable
	GetURL() string
}

type HasAPIResource interface {
	GetAPIResource() APIResource
}

type APIResourceList interface {
	getListParams() ListParams
	getResults() []HasAPIResource
}

type ResourceAmount interface {
	HasAPIResource
	GetName() string
	GetVal() int32
}
