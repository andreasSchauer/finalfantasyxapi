package queryparam

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"

func GetQueryParam(params []QueryParam, target ts.QueryParamName) QueryParam {
	for _, param := range params {
		if param.Name == target {
			return param
		}
	}

	return QueryParam{}
}
