package api

import (
	"net/http"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func servicePost[P serviceParams, R serviceResponse](cfg *Config, r *http.Request, i handlerInputService[P, R]) (R, error) {
	var zero R
	params, userValues, err := getParamsJSONBody[P](r)
	if err != nil {
		return zero, err
	}

	return compute(cfg, i, params, userValues)
}

func compute[P serviceParams, R serviceResponse](cfg *Config, i handlerInputService[P, R], params P, userValues map[ts.FieldName]any) (R, error) {
	var zero R

	params, err := verifyJSONParams(cfg, i, params, userValues)
	if err != nil {
		return zero, err
	}

	return i.executeFn(cfg, params)
}

func verifyJSONParams[P serviceParams, R serviceResponse](cfg *Config, i handlerInputService[P, R], params P, userValues map[ts.FieldName]any) (P, error) {
	var zero P
	docs := i.docsFn(cfg)
	paramsRules := compileParamsRules(docs.Fields)

	v := createParamsVerifyer(cfg, paramsRules, userValues)

	err := v.checkFields()
	if err != nil {
		return zero, err
	}

	return i.verifyFn(v, params)

}
