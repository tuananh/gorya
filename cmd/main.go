package main

import (
	"os"

	"github.com/nduyphuong/gorya/internal/signals"

	"github.com/nduyphuong/gorya/internal/logging"
)

func main() {
	ctx := signals.SetupSignalHandler()
	if err := Execute(ctx); err != nil {
		logging.LoggerFromContext(ctx).Error(err)
		os.Exit(1)
	}
}
