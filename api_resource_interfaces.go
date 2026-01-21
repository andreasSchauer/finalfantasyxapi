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
	getResults() 	[]HasAPIResource
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
