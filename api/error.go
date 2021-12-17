package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// errorCode maps custom code to http status code
type errorCode struct {
	Code    string
	Status  int
	Message string
}

type validationError map[string][]string

func (v *validationError) Error() string {
	b, _ := json.Marshal(v)
	return string(b)
}

func (v validationError) add(key string, val string) {
	v[key] = append(v[key], val)
}

func (v validationError) extend(prefix string, err *validationError) {
	if err == nil {
		return
	}
	for k, e := range *err {
		k = prefix + k
		v[k] = append(v[k], e...)
	}
}


var (
	errBadRequest            = &errorCode{Code: "400001", Status: http.StatusBadRequest}
	errURINotFound           = &errorCode{Code: "404001", Status: http.StatusNotFound}
	errInvalidMethod         = &errorCode{Code: "405001", Status: http.StatusMethodNotAllowed}
	errEntityNotUnique       = &errorCode{Code: "409001", Status: http.StatusConflict}
	errInvalidData           = &errorCode{Code: "422001", Status: http.StatusUnprocessableEntity}
	errInternalServer        = &errorCode{Code: "500001", Status: http.StatusInternalServerError}
	errInternalServerNetPipe = &errorCode{Code: "500002", Status: http.StatusInternalServerError}
)

type apiErrors []apiError

// apiError represents an error object of response
type apiError struct {
	ID     string          `json:"id"`
	Code   string          `json:"code"`
	Detail json.RawMessage `json:"detail,omitempty"`
	Status int             `json:"status"`
	Title  string          `json:"title"`
	source error
	tags   map[string]string
}

func (err *apiError) buildTags() map[string]string {
	tags := err.tags
	if tags == nil {
		tags = map[string]string{}
	}
	tags["error_id"] = err.ID
	tags["title"] = err.Title
	tags["code"] = err.Code
	return tags
}

func (err *apiError) Error() string {
	return fmt.Sprintf("api: ID: %s | Code: %s | Status: %d | Title: %s", err.ID, err.Code, err.Status, err.Title)
}

// newAPIError returns an apiError object
func newAPIError(title string, eC *errorCode, src error) *apiError {
	err := &apiError{
		ID:     uuid.NewV4().String(),
		Code:   eC.Code,
		Status: eC.Status,
		Title:  title,
		source: src,
	}
	if src != nil && eC.Status < 500 {
		if _, ok := src.(*validationError); ok {
			err.Detail = json.RawMessage(src.Error())
		} else {
			err.Detail, _ = json.Marshal(src.Error())
		}
	}
	return err
}

// newAPIErrorWithTags returns an apiError object with tags
func newAPIErrorWithTags(title string, eC *errorCode, src error, tags map[string]string) *apiError {
	err := newAPIError(title, eC, src)
	err.tags = tags
	return err
}
