package api

import (
	"net/http"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func servicePost[P serviceParams, R serviceResponse](server *Server, r *http.Request, i handlerInputService[P, R]) (R, error) {
	var zero R
	params, userValues, err := getParamsJSONBody[P](r)
	if err != nil {
		return zero, err
	}

	return compute(server, i, params, userValues)
}

func compute[P serviceParams, R serviceResponse](server *Server, i handlerInputService[P, R], params P, userValues map[ts.FieldName]any) (R, error) {
	var zero R

	params, err := verifyJSONParams(server, i, params, userValues)
	if err != nil {
		return zero, err
	}

	return i.executeFn(server, params)
}

func verifyJSONParams[P serviceParams, R serviceResponse](server *Server, i handlerInputService[P, R], params P, userValues map[ts.FieldName]any) (P, error) {
	var zero P
	docs := i.docsFn(server)
	paramsRules := compileParamsRules(docs.Fields)

	v := createParamsVerifyer(server, paramsRules, userValues)

	err := v.checkFields()
	if err != nil {
		return zero, err
	}

	return i.verifyFn(v, params)

}
