package utils

import (
	"time"
)

var squareTimeLayout = time.RFC3339Nano //"2006-01-02T15:04:05Z"

func ValidTime(t string, acceptBlank bool) bool {
	if t == "" {
			if acceptBlank { return true }
			return false
	}

	_, err := time.Parse(squareTimeLayout, t)
	return err == nil
}