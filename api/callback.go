package api

import (
	"log"
	"net/http"
)

type callbackPld struct {
	ObjectIds []int64 `json:"object_ids"`
}

func (c *callbackPld) validate() *validationError {
	errV := validationError{}

	if len(c.ObjectIds) <= 0 {
		errV.add("object_ids", "are required") //Todo: ask - minimum length of object id
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
		log.Print(v) // todo remove
	}

	resp := response{
		code: http.StatusOK,
		Data: object{},
	}
	resp.serveJSON(w)
}
