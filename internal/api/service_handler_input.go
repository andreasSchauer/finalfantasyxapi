package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/serviceverify"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type HandlerInputService[P servicetypes.ServiceParams, R servicetypes.ServiceResponse] struct {
	Endpoint  ts.EndpointName
	Usage     []string
	VerifyFn  func(serviceverify.ParamsVerifyer, P) (P, error)
	ExecuteFn func(P) (R, error)
	DocsFn    func() servicetypes.ParamsDoc
	Samples   []servicetypes.SampleParams[P]
}
