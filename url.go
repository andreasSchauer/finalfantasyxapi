package main

import "fmt"

func (cfg *Config) createURL(endpoint string, id int32) string {
	return fmt.Sprintf("http://%s/api/%s/%d", cfg.host, endpoint, id)
}

func (cfg *Config) completeURL(path string) string {
	return fmt.Sprintf("http://%s/api%s", cfg.host, path)
}
