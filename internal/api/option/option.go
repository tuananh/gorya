package option

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/nduyphuong/gorya/internal/logging"
)

func NewHandlerOption() connect.HandlerOption {

	return connect.WithHandlerOptions(
		connect.WithCodec(newJSONCodec("json")),
		connect.WithCodec(newJSONCodec("json; charset=utf-8")),
		connect.WithRecover(
			func(ctx context.Context, spec connect.Spec, header http.Header, r any) error {
				logging.LoggerFromContext(ctx).Log(log.ErrorLevel, takeStacktrace(defaultStackLength, 3))
				return connect.NewError(
					connect.CodeInternal, fmt.Errorf("panic: %v", r))
			}),
	)
}
