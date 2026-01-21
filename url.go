package main

import (
	"fmt"
	"net/url"
)

func createResourceURL(cfg *Config, endpoint string, id int32) string {
	return fmt.Sprintf("http://%s/api/%s/%d", cfg.host, endpoint, id)
}

func createResourceURLQuery(cfg *Config, endpoint string, id int32, q url.Values) string {
	url := createResourceURL(cfg, endpoint, id)
	return fmt.Sprintf("%s?%s", url, q.Encode())
}

func createListURL(cfg *Config, endpoint string) string {
	return fmt.Sprintf("http://%s/api/%s", cfg.host, endpoint)
}

func createSectionURL(cfg *Config, endpoint, section string) string {
	url := createListURL(cfg, endpoint)
	return fmt.Sprintf("%s/%s", url, section)
}

func completeTestURL(cfg *Config, path string) string {
	return fmt.Sprintf("http://%s/api%s", cfg.host, path)
}

func completeTestPath(endpoint string, id int32) string {
	return fmt.Sprintf("/%s/%d", endpoint, id)
}
