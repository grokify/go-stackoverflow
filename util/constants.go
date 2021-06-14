package util

import (
	stackoverflow "github.com/grokify/go-stackoverflow/client"
)

const (
	SiteStackOverflow = "stackoverflow"
)

func APIErrorBody(err error) []byte {
	if err != nil {
		openAPIErr := err.(stackoverflow.GenericOpenAPIError)
		return openAPIErr.Body()
	}
	return []byte{}
}
