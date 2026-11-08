package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/serviceverify"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type ServiceInput[P servicetypes.ServiceParams, R servicetypes.ServiceResponse] struct {
	Endpoint  ts.EndpointName
	Usage     []string
	VerifyFn  func(serviceverify.ParamsVerifyer, P) (P, error)
	ExecuteFn func(P) (R, error)
	DocsFn    func() servicetypes.ParamsDoc
	Samples   []servicetypes.SampleParams[P]
}

func (i ServiceInput[P, R]) ServicePost(r *http.Request) (R, error) {
	var zero R
	params, userValues, err := servicetypes.GetParamsJSONBody[P](r)
	if err != nil {
		return zero, err
	}

	return i.compute(params, userValues)
}

func (i ServiceInput[P, R]) compute(params P, userValues map[ts.FieldName]any) (R, error) {
	var zero R

	params, err := i.verifyJSONParams(params, userValues)
	if err != nil {
		return zero, err
	}

	return i.ExecuteFn(params)
}

func (i ServiceInput[P, R]) verifyJSONParams(params P, userValues map[ts.FieldName]any) (P, error) {
	var zero P
	docs := i.DocsFn()

	v := serviceverify.NewParamsVerifyer(docs.CompileRules(), userValues)

	err := v.CheckFields()
	if err != nil {
		return zero, err
	}

	return i.VerifyFn(v, params)

}

// VerifyQueries verifies, if no query params are in use for an endpoint that receives a post request
func (i ServiceInput[P, R]) VerifyQueries(r *http.Request) error {
	if len(r.URL.Query()) > 0 {
		return errs.NewHTTPError(http.StatusBadRequest, "computational endpoints don't have any query parameters", nil)
	}

	return nil
}
