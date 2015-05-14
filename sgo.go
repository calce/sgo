package sgo

import (
	"time"
)

var Token string = ""
var MerchantId string = ""
var AutoLoadNextPage bool = false

var defaultBackend *Backend

func GetDefaultBackend() (*Backend) {
	if defaultBackend == nil {
		defaultBackend = NewBackend(Token, "https://connect.squareup.com/v1", MerchantId, time.Second * 10, nil)
	}
	return defaultBackend
}

func SetDefaultBackend(backend *Backend) {
	defaultBackend = backend
}