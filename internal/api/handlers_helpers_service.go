package api

import (
	"net/http"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// most if not all of these should be methods

func servicePost[P ServiceParams, R ServiceResponse](r *http.Request, i HandlerInputService[P, R]) (R, error) {
	var zero R
	params, userValues, err := getParamsJSONBody[P](r)
	if err != nil {
		return zero, err
	}

	return compute(i, params, userValues)
}

func compute[P ServiceParams, R ServiceResponse](i HandlerInputService[P, R], params P, userValues map[ts.FieldName]any) (R, error) {
	var zero R

	params, err := verifyJSONParams(i, params, userValues)
	if err != nil {
		return zero, err
	}

	return i.executeFn(params)
}

func verifyJSONParams[P ServiceParams, R ServiceResponse](i HandlerInputService[P, R], params P, userValues map[ts.FieldName]any) (P, error) {
	var zero P
	docs := i.docsFn()
	paramsRules := compileParamsRules(docs.Fields)

	v := createParamsVerifyer(paramsRules, userValues)

	err := v.checkFields()
	if err != nil {
		return zero, err
	}

	return i.verifyFn(v, params)

}
