package api

func verifyAlBhedParams(v paramsVerifyer, params AlBhedParams) (AlBhedParams, error) {
	var err error

	params.Text, err = v.field(params.Text, PfnText)
	if err != nil {
		return AlBhedParams{}, err
	}

	params.Direction, err = v.enum(params.Direction, PfnDirection)
	if err != nil {
		return AlBhedParams{}, err
	}

	return params, nil
}
