package api

import (
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getBaseURL() string {
	return fmt.Sprintf("http://%s/api", config.Cfg.Host)
}

func createListURL(endpoint ts.EndpointName) string {
	return appendToURL(getBaseURL(), string(endpoint))
}

func createEnumURL(e ts.EnumName) string {
	return fmt.Sprintf("%s/%s/%s", getBaseURL(), ts.EpEnums, ts.GetEnumKey(e))
}

func URLToPath(url string) string {
	prefix := getBaseURL()
	return strings.TrimPrefix(url, prefix)
}

func CreateResourceURL(endpoint ts.EndpointName, id int32) string {
	url := createListURL(endpoint)
	return fmt.Sprintf("%s/%d", url, id)
}

func CreateResourceURLQuery(endpoint ts.EndpointName, id int32, q url.Values) string {
	url := CreateResourceURL(endpoint, id)
	return fmt.Sprintf("%s?%s", url, q.Encode())
}

func CompleteTestURL(path string) string {
	return fmt.Sprintf("%s%s", getBaseURL(), path)
}

func CompleteTestPath(endpoint ts.EndpointName, id int32) string {
	return fmt.Sprintf("/%s/%d", endpoint, id)
}

func getIDFromURL(url string) int32 {
	urlTrimmed := strings.TrimSuffix(url, "/")
	segments := strings.Split(urlTrimmed, "/")
	idStr := segments[len(segments)-1]

	id, _ := strconv.Atoi(idStr)
	return int32(id)
}

func getSectionURL(endpoint ts.EndpointName, section ts.SectionName) string {
	return fmt.Sprintf("%s/%s/%s", getBaseURL(), endpoint, section)
}

func appendToURL(url, segment string) string {
	return fmt.Sprintf("%s/%s", url, segment)
}

func endpointsToURLs(source []ts.EndpointName) []string {
	urls := make([]string, 0, len(source))

	for _, ep := range source {
		url := createListURL(ep)
		urls = append(urls, url)
	}

	slices.Sort(urls)
	return urls
}
