package errors

import (
	"google.golang.org/api/googleapi"
	"net/http"
)

// IsNotFound reports whether err is the result of the server replying with http.StatusNotFound.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	ae, ok := err.(*googleapi.Error)
	return ok && ae.Code == http.StatusNotFound
}
