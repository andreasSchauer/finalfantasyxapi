package api

func verifyAlBhedParams(cfg *Config, params AlBhedParams, paramsRules ParamsRules, userValues map[FieldName]any) (AlBhedParams, error) {
	var err error
	
	params.Text, err = verifyParamField(cfg, params.Text, pfnText, userValues, paramsRules, nil)
	if err != nil {
		return AlBhedParams{}, err
	}

	params.Direction, err = verifyParamField(cfg, params.Direction, pfnDirection, userValues, paramsRules, vfEnum)
	if err != nil {
		return AlBhedParams{}, err
	}

	return params, nil
}
