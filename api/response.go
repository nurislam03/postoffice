package api

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"syscall"
)

type (
	// object represents an object
	object map[string]interface{}

	// response represents a response
	response struct {
		code   int
		Data   interface{} `json:"data,omitempty"`
		Meta   *pager      `json:"meta,omitempty"`
		Errors []apiError  `json:"errors,omitempty"`
	}
)

// serveJSON serves the response to writer as JSON
func (resp *response) serveJSON(w http.ResponseWriter) {
	if resp.code == 0 {
		panic(errors.New("response status not defined"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.code)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

// parseBody parses request body to given data struct
func parseBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func handleAPIError(w http.ResponseWriter, err interface{}) {
	switch err := err.(type) {
	case apiErrors:
		st := http.StatusOK
		for _, e := range err {
			if e.Status > st {
				st = e.Status
			}
		}
		resp := response{
			code:   st,
			Errors: err,
		}
		resp.serveJSON(w)
		logFatalErrors(err)

	case *apiError:
		resp := response{
			code:   err.Status,
			Errors: []apiError{*err},
		}
		resp.serveJSON(w)
		logFatalError(err)

	case *net.OpError:
		if err.Err == syscall.EPIPE || err.Err == syscall.ECONNRESET {
			break
		}

		resp := response{
			code: http.StatusInternalServerError,
			Errors: []apiError{
				*newAPIError("Internal Server error", errInternalServerNetPipe, err),
			},
		}
		resp.serveJSON(w)
		logFatalError(&resp.Errors[0])

	case error:
		resp := response{
			code: http.StatusInternalServerError,
			Errors: []apiError{
				*newAPIError("Internal Server error", errInternalServer, err),
			},
		}
		resp.serveJSON(w)
		logFatalError(&resp.Errors[0])

	case string:
		resp := response{
			code: http.StatusInternalServerError,
			Errors: []apiError{
				*newAPIError("Internal Server error", errInternalServer, errors.New(err)),
			},
		}
		resp.serveJSON(w)
		logFatalError(&resp.Errors[0])

	default:
		panic(err)
	}

}
