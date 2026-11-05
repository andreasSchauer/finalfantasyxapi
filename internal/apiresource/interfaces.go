// Package apiresource jaisdüasd
package apiresource

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type APIResource interface {
	HasAPIResource
	helpers.HasID
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