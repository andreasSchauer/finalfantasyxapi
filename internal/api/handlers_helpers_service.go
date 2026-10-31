package api

import "net/http"

func servicePost[P, R any](cfg *Config, r *http.Request, i handlerInputService[P, R]) (R, error) {
	var zero R
	params, userValues, err := getParamsJSONBody[P](r)
	if err != nil {
		return zero, err
	}

	return compute(cfg, i, params, userValues)
}

func compute[P, R any](cfg *Config, i handlerInputService[P, R], params P, userValues map[FieldName]any) (R, error) {
	var zero R

	params, err := verifyJSONParams(cfg, i, params, userValues)
	if err != nil {
		return zero, err
	}

	return i.executeFn(cfg, params)
}

func verifyJSONParams[P, R any](cfg *Config, i handlerInputService[P, R], params P, userValues map[FieldName]any) (P, error) {
	var zero P
	docs := i.docsFn(cfg)
	paramsRules := compileParamsRules(docs.Fields)

	v := createParamsVerifyer(cfg, paramsRules, userValues)

	err := v.vfExistingFields()
	if err != nil {
		return zero, err
	}
	return i.verifyFn(v, params)

}
