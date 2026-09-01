package api


func verifyAlBhedParams(cfg *Config, params AlBhedParams, valueMap map[FieldName]any) (AlBhedParams, error) {
	var err error
	valTree := compileValidationTree(params.GetDoc(cfg).Fields)

	err = vfExistingFields(valueMap, valTree)
	if err != nil {
		return AlBhedParams{}, err
	}

	params.Text, err = verifyParamField(cfg, params.Text, pfnText, valueMap, valTree, nil)
	if err != nil {
		return AlBhedParams{}, err
	}

	params.Direction, err = verifyParamField(cfg, params.Direction, pfnDirection, valueMap, valTree, vfEnum)
	if err != nil {
		return AlBhedParams{}, err
	}

	return params, nil
}
