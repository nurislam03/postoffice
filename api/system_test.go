package api

import (
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSystemCheck(t *testing.T) {
	api := NewAPI(nil, nil, nil)

	t.Run("system check", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/system/check", nil)
		res := httptest.NewRecorder()
		api.router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, `{"data":{"system_status":"ok"}}`, res.Body.String())
	})
}

func TestSystemPanic(t *testing.T) {
	api := NewAPI(nil, nil, nil)

	t.Run("system panic", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/system/panic", nil)
		res := httptest.NewRecorder()
		api.router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		jsonassert.New(t).Assertf(res.Body.String(), `{"errors":[{"id":"<<PRESENCE>>","code":"500001","status":500,"title":"System Panic"}]}`)
	})
}

func TestSystemErr(t *testing.T) {
	api := NewAPI(nil, nil, nil)

	t.Run("system err", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/system/err", nil)
		res := httptest.NewRecorder()
		api.router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		jsonassert.New(t).Assertf(res.Body.String(), `{"errors":[{"id":"<<PRESENCE>>","code":"500001","status":500,"title":"Internal Server error"}]}`)
	})
}

func TestSystemValidationErr(t *testing.T) {
	api := NewAPI(nil, nil, nil)

	t.Run("system validation err", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/system/verr", nil)
		res := httptest.NewRecorder()
		api.router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		jsonassert.New(t).Assertf(res.Body.String(), `{"errors":[{"id":"<<PRESENCE>>","code":"422001","status":422,"title":"System Validation Error"}]}`)
	})
}
