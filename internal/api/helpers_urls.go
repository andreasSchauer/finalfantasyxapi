package api

import (
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getBaseURL(server *Server) string {
	return fmt.Sprintf("http://%s/api", server.Cfg.Host)
}

func createListURL(server *Server, endpoint ts.EndpointName) string {
	return appendToURL(getBaseURL(server), string(endpoint))
}

func createEnumURL(server *Server, e ts.EnumName) string {
	return fmt.Sprintf("%s/%s/%s", getBaseURL(server), ts.EpEnums, ts.GetEnumKey(e))
}

func URLToPath(server *Server, url string) string {
	prefix := getBaseURL(server)
	return strings.TrimPrefix(url, prefix)
}

func CreateResourceURL(server *Server, endpoint ts.EndpointName, id int32) string {
	url := createListURL(server, endpoint)
	return fmt.Sprintf("%s/%d", url, id)
}

func CreateResourceURLQuery(server *Server, endpoint ts.EndpointName, id int32, q url.Values) string {
	url := CreateResourceURL(server, endpoint, id)
	return fmt.Sprintf("%s?%s", url, q.Encode())
}

func CompleteTestURL(server *Server, path string) string {
	return fmt.Sprintf("%s%s", getBaseURL(server), path)
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

func getSectionURL(server *Server, endpoint ts.EndpointName, section ts.SectionName) string {
	return fmt.Sprintf("%s/%s/%s", getBaseURL(server), endpoint, section)
}

func appendToURL(url, segment string) string {
	return fmt.Sprintf("%s/%s", url, segment)
}

func endpointsToURLs(server *Server, source []ts.EndpointName) []string {
	urls := make([]string, 0, len(source))

	for _, ep := range source {
		url := createListURL(server, ep)
		urls = append(urls, url)
	}

	slices.Sort(urls)
	return urls
}
