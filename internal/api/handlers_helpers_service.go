package api

import "net/http"


func serviceGet[P ServiceParams, R ServiceResponse](cfg *Config, r *http.Request, i handlerInputService[P, R]) (R, error) {
	var zero R
	params, payloadMap, err := getParamsStateURL[P](r, i.queryLookup)
	if err != nil {
		return zero, err
	}

	return compute(cfg, i, params, payloadMap)
}

func servicePost[P ServiceParams, R ServiceResponse](cfg *Config, r *http.Request, i handlerInputService[P, R]) (R, error) {
	var zero R
	params, payloadMap, err := getParamsJsonBody[P](r)
	if err != nil {
		return zero, err
	}

	return compute(cfg, i, params, payloadMap)
}


func compute[P ServiceParams, R ServiceResponse](cfg *Config, i handlerInputService[P, R], params P, payloadMap map[string]any) (R, error) {
	var zero R

	url, err := paramsToStateURL(cfg, i.endpoint, params)
	if err != nil {
		return zero, err
	}

	params, err = i.verifyFn(cfg, params, payloadMap)
	if err != nil {
		return zero, err
	}

	return i.executeFn(cfg, params, url)
}