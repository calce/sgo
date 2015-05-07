package sgo

import (
	"time"
)

var Token string = ""
var MerchantId string = ""
var AutoLoadNextPage bool = false

var defaultBackend *Backend

func GetDefaultBackend() (*Backend, error) {
	if defaultBackend == nil {
		var err error
		defaultBackend, err = NewBackend(Token, "https://connect.squareup.com/v1", MerchantId, time.Second * 10, nil)
		if err != nil { return nil, err }
	}
	return defaultBackend, nil
}

func SetDefaultBackend(backend *Backend) {
	defaultBackend = backend
}