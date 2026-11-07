package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i HandlerInputService[P, R]) ServicePost(r *http.Request) (R, error) {
	var zero R
	params, userValues, err := getParamsJSONBody[P](r)
	if err != nil {
		return zero, err
	}

	return i.compute(params, userValues)
}

func (i HandlerInputService[P, R]) compute(params P, userValues map[ts.FieldName]any) (R, error) {
	var zero R

	params, err := i.verifyJSONParams(params, userValues)
	if err != nil {
		return zero, err
	}

	return i.ExecuteFn(params)
}

func (i HandlerInputService[P, R]) verifyJSONParams(params P, userValues map[ts.FieldName]any) (P, error) {
	var zero P
	docs := i.DocsFn()

	v := createParamsVerifyer(docs.CompileRules(), userValues)

	err := v.checkFields()
	if err != nil {
		return zero, err
	}

	return i.VerifyFn(v, params)

}

// VerifyQueries verifies, if no query params are in use for an endpoint that receives a post request
func (i HandlerInputService[P, R]) VerifyQueries(r *http.Request) error {
	if len(r.URL.Query()) > 0 {
		return errs.NewHTTPError(http.StatusBadRequest, "computational endpoints don't have any query parameters", nil)
	}

	return nil
}
