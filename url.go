package main

import (
	"fmt"
	"net/url"
)

func (cfg *Config) createResourceURL(endpoint string, id int32) string {
	return fmt.Sprintf("http://%s/api/%s/%d", cfg.host, endpoint, id)
}

func (cfg *Config) createResourceURLQuery(endpoint string, id int32, q url.Values) string {
	url := cfg.createResourceURL(endpoint, id)
	return fmt.Sprintf("%s?%s", url, q.Encode())
}

func (cfg *Config) createListURL(endpoint string) string {
	return fmt.Sprintf("http://%s/api/%s", cfg.host, endpoint)
}

func (cfg *Config) createSectionURL(endpoint, section string) string {
	url := cfg.createListURL(endpoint)
	return fmt.Sprintf("%s/%s", url, section)
}

func (cfg *Config) completeTestURL(path string) string {
	return fmt.Sprintf("http://%s/api%s", cfg.host, path)
}

func completeTestPath(endpoint string, id int32) string {
	return fmt.Sprintf("/%s/%d", endpoint, id)
}