package api

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

type HandlerInputService[P ServiceParams, R ServiceResponse] struct {
	Endpoint  ts.EndpointName
	Usage     []string
	VerifyFn  func(paramsVerifyer, P) (P, error)
	ExecuteFn func(P) (R, error)
	DocsFn    func() ParamsDoc
	SamplesFn func() []SampleParams[P]
}
