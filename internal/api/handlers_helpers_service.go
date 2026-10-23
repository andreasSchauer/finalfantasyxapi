package api

import "net/http"

func servicePost[P, R any](cfg *Config, r *http.Request, i handlerInputService[P, R]) (R, error) {
	var zero R
	params, userValues, err := getParamsJsonBody[P](r)
	if err != nil {
		return zero, err
	}

	return compute(cfg, i, params, userValues)
}

func compute[P, R any](cfg *Config, i handlerInputService[P, R], params P, userValues map[FieldName]any) (R, error) {
	var zero R

	params, err := i.verifyFn(cfg, params, i.docsFn(cfg), userValues)
	if err != nil {
		return zero, err
	}

	return i.executeFn(cfg, params)
}
