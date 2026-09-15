package api

import "net/http"

func servicePost[P ServiceParams, R any](cfg *Config, r *http.Request, i handlerInputService[P, R]) (R, error) {
	var zero R
	params, valueMap, err := getParamsJsonBody[P](r)
	if err != nil {
		return zero, err
	}

	return compute(cfg, i, params, valueMap)
}

func compute[P ServiceParams, R any](cfg *Config, i handlerInputService[P, R], params P, valueMap map[FieldName]any) (R, error) {
	var zero R

	params, err := i.verifyFn(cfg, params, valueMap)
	if err != nil {
		return zero, err
	}

	return i.executeFn(cfg, params)
}
