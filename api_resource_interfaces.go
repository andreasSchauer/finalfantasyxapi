package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type HasAPIResource interface {
	GetAPIResource() APIResource
}

type APIResource interface {
	GetID() int32
	GetURL() string
	seeding.Lookupable
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
