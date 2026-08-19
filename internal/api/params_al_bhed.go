package api

type AlBhedParams struct {
	Text      string    `json:"text"`
	Direction string 	`json:"direction"`
}

func verifyAlBhedParams(cfg *Config, params AlBhedParams, payloadMap map[string]any) (AlBhedParams, error) {
	var err error
	lookup := fieldDocSliceToMap(cfg.getAlBhedFieldDoc())
	
	params.Text, err = verifyParamField(params.Text, "text", payloadMap, lookup)
	if err != nil {
		return AlBhedParams{}, err
	}

	params.Direction, err = verifyParamField(params.Direction, "direction", payloadMap, lookup)
	if err != nil {
		return AlBhedParams{}, err
	}
	
	return params, nil
}