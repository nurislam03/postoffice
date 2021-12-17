package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/nurislam03/postoffice/config"
	"github.com/streadway/amqp"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// API ...
type API struct {
	router   chi.Router
	cfg      *config.Config
	amqpServer *amqp.Connection
}

// NewAPI ...
func NewAPI(cfg *config.Config, amqp *amqp.Connection) *API {
	api := &API{
		router:   chi.NewRouter(),
		cfg:      cfg,
		amqpServer: amqp,
	}

	api.register()
	api.registerRouter()
	return api
}

var logger = logrus.New()

func init() {
	SetLogLevel()
}

// Handler ...
func (a *API) Handler() http.Handler {
	return a.router
}

func (a *API) register() {

	a.router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	a.router.Use(recoverer)

	a.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		err := newAPIError("Not Found", errURINotFound, nil)
		panic(err)
	})

	a.router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		err := newAPIError("Method Not Allowed", errInvalidMethod, nil)
		resp := response{
			code:   http.StatusMethodNotAllowed,
			Errors: []apiError{*err},
		}
		resp.serveJSON(w)
	})
}

func (a *API) registerRouter() {
	a.router.Route("/", func(r chi.Router) {
		r.Mount("/callback", a.callbackHandlers())
		r.Mount("/system", a.systemHandlers())
	})
}

func (a *API) callbackHandlers() http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Post("/", a.Callback)
	})

	return h
}

func (a *API) systemHandlers() http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Get("/check", a.systemCheck)
		r.Get("/panic", a.systemPanic)
		r.Get("/err", a.systemErr)
		r.Get("/verr", a.systemValidationErr)
	})

	return h
}
