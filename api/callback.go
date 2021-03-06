package api

import (
	"net/http"
	"strconv"
)

type callbackPld struct {
	ObjectIds []int64 `json:"object_ids"`
}

func (c *callbackPld) validate() *validationError {
	errV := validationError{}

	if len(c.ObjectIds) <= 0 {
		errV.add("object_ids", "are required")
	}

	if len(errV) > 0 {
		return &errV
	}

	return nil
}

// Callback ...
func (a *API) Callback(w http.ResponseWriter, r *http.Request) {
	body := callbackPld{}
	if err := parseBody(r, &body); err != nil {
		handleAPIError(w, newAPIError("Unable to parse body", errBadRequest, err))
		return
	}

	if err := body.validate(); err != nil {
		handleAPIError(w, newAPIError("Invalid Data", errInvalidData, err))
		return
	}

	for _, v := range body.ObjectIds {
		a.PushToPublisher("get-status", strconv.FormatInt(v, 10))
	}

	resp := response{
		code: http.StatusOK,
		Data: object{},
	}
	resp.serveJSON(w)
}
