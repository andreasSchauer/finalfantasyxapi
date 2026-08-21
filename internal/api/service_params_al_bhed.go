package api

type AlBhedParams struct {
	Text      string `json:"text,omitempty"`
	Direction string `json:"direction,omitempty"`
}

func (p AlBhedParams) GetDoc(cfg *Config) ParamsDoc {
	return cfg.getAlBhedParamsDoc()
}

func verifyAlBhedParams(cfg *Config, params AlBhedParams, payloadMap map[string]any) (AlBhedParams, error) {
	var err error

	lookup := paramsDocToFieldMap(params.GetDoc(cfg))

	params.Text, err = verifyParamField(params.Text, "text", payloadMap, lookup, nil)
	if err != nil {
		return AlBhedParams{}, err
	}

	params.Direction, err = verifyParamField(params.Direction, "direction", payloadMap, lookup, vfEnum)
	if err != nil {
		return AlBhedParams{}, err
	}

	return params, nil
}
