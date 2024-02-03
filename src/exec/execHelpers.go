// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/exec/execHelpers.go
// Original timestamp: 2024/01/07 13:17

package exec

import (
	"fmt"
	"net/url"
	"nxrmuploader/helpers"
	"path"
)

var IndexNumber int8

func parseURL(fullURL string) (string, string, error) {
	//var url URL
	parsed, err := url.Parse(fullURL)
	if err != nil {
		return "", "", helpers.CustomError{fmt.Sprintf("Unable to parse URL: %s", err)}
	}

	fqdn := parsed.Scheme + "://" + parsed.Host
	if parsed.Port() != "" {
		fqdn += ":" + parsed.Port()
	}
	endpoint := path.Base(parsed.Path)

	return fqdn, endpoint, nil
}
