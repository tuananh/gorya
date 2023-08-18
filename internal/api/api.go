package api

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/nduyphuong/gorya/internal/api/config"
	"github.com/nduyphuong/gorya/internal/api/handler"
	"github.com/nduyphuong/gorya/internal/api/option"
	"github.com/nduyphuong/gorya/internal/logging"
	"github.com/nduyphuong/gorya/internal/store"
	"github.com/nduyphuong/gorya/internal/version"
	svcv1alpha1 "github.com/nduyphuong/gorya/pkg/api/service/v1alpha1"
	"github.com/pkg/errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server interface {
	Serve(ctx context.Context, l net.Listener) error
}

type server struct {
	cfg config.ServerConfig
}

func NewServer(cfg config.ServerConfig) (Server, error) {
	return &server{
		cfg: cfg,
	}, nil
}

func (s *server) Serve(ctx context.Context, l net.Listener) error {
	errCh := make(chan error)
	log := logging.LoggerFromContext(ctx)
	log.Infof("Server is listening on %q", l.Addr().String())
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	opts := option.NewHandlerOption()
	store, err := store.GetSingleton()
	if err != nil {
		return err
	}
	path, svcHandler := svcv1alpha1.NewGoryaServiceHandler(ctx, store, s, opts)
	mux.Handle(path, svcHandler)
	srv := &http.Server{
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: time.Minute,
	}
	go func() { errCh <- srv.Serve(l) }()

	select {
	case <-ctx.Done():
		log.Info("Gracefully stopping server...")
		time.Sleep(s.cfg.GracefulShutdownTimeout)
		return srv.Shutdown(context.Background())
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}

func (s *server) GetTimeZone() http.Handler {
	return handler.GetTimeZoneV1Alpha1()
}

func (s *server) GetVersionInfo() http.Handler {
	return handler.GetVersionInfoV1Alpha1(version.GetVersion())
}

func (s *server) AddSchedule(ctx context.Context, store store.Interface) http.Handler {
	return handler.AddScheduleV1Alpha1(ctx, store)
}
