package api

import (
	"net/http"
)

// systemCheck ...
func (a *API) systemCheck(w http.ResponseWriter, r *http.Request) {
	resp := response{
		code: http.StatusOK,
		Data: object{
			"system_status": "ok",
		},
	}
	resp.serveJSON(w)
}

// systemPanic ...
func (a *API) systemPanic(w http.ResponseWriter, r *http.Request) {
	panic(newAPIError("System Panic", errInternalServer, nil))
}

// systemErr ...
func (a *API) systemErr(w http.ResponseWriter, r *http.Request) {
	x, y := 1, 0
	_ = x / y
}

// systemValidationErr ...
func (a *API) systemValidationErr(w http.ResponseWriter, r *http.Request) {
	err := validationError{}
	err.add("Hello", "World")
	err.add("Error", "Validation error")
	panic(newAPIError("System Validation Error", errInvalidData, nil))
}
