package api

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func createListURL(cfg *Config, endpoint EndpointName) string {
	return fmt.Sprintf("http://%s/api/%s", cfg.host, endpoint)
}

func urlToPath(cfg *Config, url string) string {
	prefix := fmt.Sprintf("http://%s/api", cfg.host)
	return strings.TrimPrefix(url, prefix)
}

func createResourceURL(cfg *Config, endpoint EndpointName, id int32) string {
	url := createListURL(cfg, endpoint)
	return fmt.Sprintf("%s/%d", url, id)
}

func createResourceURLQuery(cfg *Config, endpoint EndpointName, id int32, q url.Values) string {
	url := createResourceURL(cfg, endpoint, id)
	return fmt.Sprintf("%s?%s", url, q.Encode())
}

func completeTestURL(cfg *Config, path string) string {
	return fmt.Sprintf("http://%s/api%s", cfg.host, path)
}

func completeTestPath(endpoint EndpointName, id int32) string {
	return fmt.Sprintf("/%s/%d", endpoint, id)
}

func getIdFromURL(url string) int32 {
	urlTrimmed := strings.TrimSuffix(url, "/")
	segments := strings.Split(urlTrimmed, "/")
	idStr := segments[len(segments)-1]

	id, _ := strconv.Atoi(idStr)
	return int32(id)
}
