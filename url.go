package php

import (
	"github.com/syyongx/php2go"
	"net/url"
	"strings"
)

// ParseURL parse_url()
// Parse a URL and return its components
// -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment
var ParseURL = php2go.ParseURL

// URLEncode urlencode()
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

// URLDecode urldecode()
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

// Rawurlencode rawurlencode()
func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

// Rawurldecode rawurldecode()
func Rawurldecode(str string) (string, error) {
	return url.QueryUnescape(strings.Replace(str, "%20", "+", -1))
}

// HTTPBuildQuery http_build_query()
func HTTPBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}
