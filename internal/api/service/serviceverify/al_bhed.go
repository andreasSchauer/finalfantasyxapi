package serviceverify

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func VerifyAlBhedParams(v ParamsVerifyer, params servicetypes.AlBhedParams) (servicetypes.AlBhedParams, error) {
	var err error

	params.Text, err = v.string(params.Text, ts.PfnText)
	if err != nil {
		return servicetypes.AlBhedParams{}, err
	}

	params.Direction, err = v.enum(params.Direction, ts.PfnDirection)
	if err != nil {
		return servicetypes.AlBhedParams{}, err
	}

	return params, nil
}
