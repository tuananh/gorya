package main

import (
	"errors"
	"fmt"
	"github.com/nduyphuong/gorya/internal/api"
	"github.com/nduyphuong/gorya/internal/api/config"
	"github.com/nduyphuong/gorya/internal/os"
	versionpkg "github.com/nduyphuong/gorya/internal/version"
	pkgerrors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net"
	"sync"
)

func newServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "api",
		DisableAutoGenTag: true,
		SilenceErrors:     true,
		SilenceUsage:      true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			version := versionpkg.GetVersion()
			log.WithFields(log.Fields{
				"version": version.Version,
				"commit":  version.GitCommit,
			}).Info("Starting Gorya API Server")
			var wg sync.WaitGroup
			errCh := make(chan error, 1)
			cfg := config.ServerConfigFromEnv()

			srv, err := api.NewServer(cfg)
			if err != nil {
				return pkgerrors.Wrap(err, "error creating API server")
			}
			l, err := net.Listen(
				"tcp",
				fmt.Sprintf(
					"%s:%s",
					os.GetEnv("HOST", "0.0.0.0"),
					os.GetEnv("PORT", "8080"),
				),
			)
			if err != nil {
				return pkgerrors.Wrap(err, "error creating listener")
			}
			defer func() {
				_ = l.Close()
			}()
			wg.Add(1)
			go func() {
				srvErr := srv.Serve(ctx, l)
				errCh <- pkgerrors.Wrap(srvErr, "serve")
				wg.Done()
			}()
			wg.Wait()
			close(errCh)
			var resErr error
			for err := range errCh {
				resErr = errors.Join(resErr, err)
			}
			return resErr
		},
	}
}
