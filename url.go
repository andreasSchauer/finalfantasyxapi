package main

import "fmt"

func (cfg *apiConfig) createURL(endpoint string, id int32) string {
	return fmt.Sprintf("http://%s/api/%s/%d", cfg.host, endpoint, id)
}
