package api

import "net/url"

// I could sum the functions that are present in both verify functions into one, but this is already quite good and flexible


// checks the usage of an endpoint that uses ids as its primary key for single resources, like /locations/{id}.
func verifyQueryUsageIdEp(q url.Values, queryParam QueryParam, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, id *int32, segment *string) error {
	err := verifyExclusiveParam(q, queryParam, queryLookup)
	if err != nil {
		return err
	}

	err = verifySegmentOnlyParam(queryParam, segment, endpoint)
	if err != nil {
		return err
	}

	err = verifySingleResourceParamID(queryParam, id)
	if err != nil {
		return err
	}

	err = verifyListResourceParamID(queryParam, id)
	if err != nil {
		return err
	}

	err = verifyRequiredParams(q, queryParam)
	if err != nil {
		return err
	}

	err = verifyForbiddenParams(q, queryParam)
	if err != nil {
		return err
	}

	err = verifyUsableWith(q, queryParam)
	if err != nil {
		return err
	}

	return nil
}


// checks the usage of an endpoint that uses keys as its primary key for single resources, like /enums/{EnumName}.
func verifyQueryUsageKeyEp(q url.Values, queryParam QueryParam, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, key, segment *string) error {
	err := verifyExclusiveParam(q, queryParam, queryLookup)
	if err != nil {
		return err
	}

	err = verifySegmentOnlyParam(queryParam, segment, endpoint)
	if err != nil {
		return err
	}

	err = verifySingleResourceParamKey(queryParam, key)
	if err != nil {
		return err
	}

	err = verifyListResourceParamKey(queryParam, key)
	if err != nil {
		return err
	}

	err = verifyRequiredParams(q, queryParam)
	if err != nil {
		return err
	}

	err = verifyForbiddenParams(q, queryParam)
	if err != nil {
		return err
	}

	err = verifyUsableWith(q, queryParam)
	if err != nil {
		return err
	}

	return nil
}