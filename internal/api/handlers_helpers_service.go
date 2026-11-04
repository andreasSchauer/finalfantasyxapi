package api

import (
	"net/http"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func servicePost[P ServiceParams, R ServiceResponse](srv *Server, r *http.Request, i handlerInputService[P, R]) (R, error) {
	var zero R
	params, userValues, err := getParamsJSONBody[P](r)
	if err != nil {
		return zero, err
	}

	return compute(srv, i, params, userValues)
}

func compute[P ServiceParams, R ServiceResponse](srv *Server, i handlerInputService[P, R], params P, userValues map[ts.FieldName]any) (R, error) {
	var zero R

	params, err := verifyJSONParams(srv, i, params, userValues)
	if err != nil {
		return zero, err
	}

	return i.executeFn(srv, params)
}

func verifyJSONParams[P ServiceParams, R ServiceResponse](srv *Server, i handlerInputService[P, R], params P, userValues map[ts.FieldName]any) (P, error) {
	var zero P
	docs := i.docsFn(srv)
	paramsRules := compileParamsRules(docs.Fields)

	v := createParamsVerifyer(srv, paramsRules, userValues)

	err := v.checkFields()
	if err != nil {
		return zero, err
	}

	return i.verifyFn(v, params)

}
