package backend

import (
	"context"
	"log"
	"github.com/nurislam03/postoffice/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"github.com/nurislam03/postoffice/api"
)

// Server ...
type Server struct {
	cfg *config.Config
	api *api.API
}

// NewServer ...
func NewServer(c *config.Config, a *api.API) *Server {
	return &Server{
		cfg:         c,
		api:         a,
	}
}

// Serve ...
func (s *Server) Serve() {

	portStr := viper.GetString("SERVER_PORT")

	srvr := &http.Server{
		ReadTimeout:  viper.GetDuration("READ_TIMEOUT") * time.Second,
		WriteTimeout: viper.GetDuration("WRITE_TIMEOUT") * time.Second,
		IdleTimeout:  viper.GetDuration("IDLE_TIMEOUT") * time.Second,
		Addr:         ":" + portStr,
		Handler:      s.api.Handler(),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		log.Println("Server Listening on :" + portStr)
		log.Fatal(srvr.ListenAndServe())
	}()

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srvr.Shutdown(ctx)

	log.Println("Server shut down gracefully")
}
