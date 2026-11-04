package api

import (
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getBaseURL(srv *Server) string {
	return fmt.Sprintf("http://%s/api", srv.Cfg.Host)
}

func createListURL(srv *Server, endpoint ts.EndpointName) string {
	return appendToURL(getBaseURL(srv), string(endpoint))
}

func createEnumURL(srv *Server, e ts.EnumName) string {
	return fmt.Sprintf("%s/%s/%s", getBaseURL(srv), ts.EpEnums, ts.GetEnumKey(e))
}

func URLToPath(srv *Server, url string) string {
	prefix := getBaseURL(srv)
	return strings.TrimPrefix(url, prefix)
}

func CreateResourceURL(srv *Server, endpoint ts.EndpointName, id int32) string {
	url := createListURL(srv, endpoint)
	return fmt.Sprintf("%s/%d", url, id)
}

func CreateResourceURLQuery(srv *Server, endpoint ts.EndpointName, id int32, q url.Values) string {
	url := CreateResourceURL(srv, endpoint, id)
	return fmt.Sprintf("%s?%s", url, q.Encode())
}

func CompleteTestURL(srv *Server, path string) string {
	return fmt.Sprintf("%s%s", getBaseURL(srv), path)
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

func getSectionURL(srv *Server, endpoint ts.EndpointName, section ts.SectionName) string {
	return fmt.Sprintf("%s/%s/%s", getBaseURL(srv), endpoint, section)
}

func appendToURL(url, segment string) string {
	return fmt.Sprintf("%s/%s", url, segment)
}

func endpointsToURLs(srv *Server, source []ts.EndpointName) []string {
	urls := make([]string, 0, len(source))

	for _, ep := range source {
		url := createListURL(srv, ep)
		urls = append(urls, url)
	}

	slices.Sort(urls)
	return urls
}
