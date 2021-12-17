package api

import (
	"errors"
	"net"
	"net/http"
	"runtime/debug"
	"syscall"

	raven "github.com/getsentry/raven-go"
)

func logFatalError(err *apiError) {
	if err.source == nil || err.Status < 500 {
		return
	}

	raven.CaptureError(err.source, err.buildTags())
	logger.Errorf("Panic: %+v\n", err.source)
	logger.Println(string(debug.Stack()))
}

func logFatalErrors(errs apiErrors) {
	for _, er := range errs {
		logFatalError(&er)
	}
}

func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				raven.ClearContext()
				raven.SetHttpContext(raven.NewHttp(r))

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
		}()

		next.ServeHTTP(w, r)
	})
}
