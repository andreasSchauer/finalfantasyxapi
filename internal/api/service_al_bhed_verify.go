package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

func verifyAlBhedParams(v paramsVerifyer, params AlBhedParams) (AlBhedParams, error) {
	var err error

	params.Text, err = v.field(params.Text, typestrings.PfnText)
	if err != nil {
		return AlBhedParams{}, err
	}

	params.Direction, err = v.enum(params.Direction, typestrings.PfnDirection)
	if err != nil {
		return AlBhedParams{}, err
	}

	return params, nil
}
