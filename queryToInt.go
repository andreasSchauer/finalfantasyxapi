package main

import "strconv"

func queryStrToInt(s string, defaultVal int) (int, error) {
	if s == "" {
		return defaultVal, nil
	} 

	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return val, nil
	
}