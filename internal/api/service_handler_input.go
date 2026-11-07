package api

import 	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

type HandlerInputService[P ServiceParams, R ServiceResponse] struct {
	Endpoint  ts.EndpointName
	usage     []string
	verifyFn  func(paramsVerifyer, P) (P, error)
	executeFn func(P) (R, error)
	docsFn    func() ParamsDoc
	samplesFn func() []SampleParams[P]
}
