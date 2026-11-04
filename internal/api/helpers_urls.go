package api

import (
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getBaseURL(cfg *Config) string {
	return fmt.Sprintf("http://%s/api", cfg.host)
}

func createListURL(cfg *Config, endpoint typestrings.EndpointName) string {
	return appendToURL(getBaseURL(cfg), string(endpoint))
}

func createEnumURL(cfg *Config, e typestrings.EnumName) string {
	return fmt.Sprintf("%s/%s/%s", getBaseURL(cfg), typestrings.EpEnums, typestrings.GetEnumKey(e))
}

func URLToPath(cfg *Config, url string) string {
	prefix := getBaseURL(cfg)
	return strings.TrimPrefix(url, prefix)
}

func CreateResourceURL(cfg *Config, endpoint typestrings.EndpointName, id int32) string {
	url := createListURL(cfg, endpoint)
	return fmt.Sprintf("%s/%d", url, id)
}

func CreateResourceURLQuery(cfg *Config, endpoint typestrings.EndpointName, id int32, q url.Values) string {
	url := CreateResourceURL(cfg, endpoint, id)
	return fmt.Sprintf("%s?%s", url, q.Encode())
}

func CompleteTestURL(cfg *Config, path string) string {
	return fmt.Sprintf("%s%s", getBaseURL(cfg), path)
}

func CompleteTestPath(endpoint typestrings.EndpointName, id int32) string {
	return fmt.Sprintf("/%s/%d", endpoint, id)
}

func getIDFromURL(url string) int32 {
	urlTrimmed := strings.TrimSuffix(url, "/")
	segments := strings.Split(urlTrimmed, "/")
	idStr := segments[len(segments)-1]

	id, _ := strconv.Atoi(idStr)
	return int32(id)
}

func getSectionURL(cfg *Config, endpoint typestrings.EndpointName, section typestrings.SectionName) string {
	return fmt.Sprintf("%s/%s/%s", getBaseURL(cfg), endpoint, section)
}

func appendToURL(url, segment string) string {
	return fmt.Sprintf("%s/%s", url, segment)
}

func endpointsToURLs(cfg *Config, source []typestrings.EndpointName) []string {
	urls := make([]string, 0, len(source))

	for _, ep := range source {
		url := createListURL(cfg, ep)
		urls = append(urls, url)
	}

	slices.Sort(urls)
	return urls
}
