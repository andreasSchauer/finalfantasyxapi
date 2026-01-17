package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type APIResource interface {
	HasAPIResource
	seeding.LookupableID
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
