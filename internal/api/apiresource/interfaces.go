// Package apiresource jaisdüasd
package apiresource

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

type APIResource interface {
	HasAPIResource
	h.HasID
	HasURL
	GetTestKey() string
}

type HasAPIResource interface {
	GetAPIResource() APIResource
}

type HasURL interface {
	GetURL() string
}

type SimpleResource interface {
	HasURL
}