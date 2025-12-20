package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type HasAPIResource interface {
	getAPIResource() IsAPIResource
}

type IsAPIResource interface {
	getID() int32
	getURL() string
	seeding.Lookupable
}

type IsAPIResourceList interface {
	getListParams() ListParams
	getResults() []HasAPIResource
}

type ResourceAmount interface {
	HasAPIResource
	GetName() string
	GetVal() int32
}