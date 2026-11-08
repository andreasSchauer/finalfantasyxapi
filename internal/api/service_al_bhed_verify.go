package api

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

func verifyAlBhedParams(v paramsVerifyer, params AlBhedParams) (AlBhedParams, error) {
	var err error

	params.Text, err = v.string(params.Text, ts.PfnText)
	if err != nil {
		return AlBhedParams{}, err
	}

	params.Direction, err = v.enum(params.Direction, ts.PfnDirection)
	if err != nil {
		return AlBhedParams{}, err
	}

	return params, nil
}
