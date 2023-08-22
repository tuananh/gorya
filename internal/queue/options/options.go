package options

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Options struct {
	Name        string
	Addr        string
	PopInterval time.Duration
	Log         *logrus.Logger
}

type Option interface {
	Apply(*Options)
}

type fetchInterval time.Duration

func (o fetchInterval) Apply(i *Options) {
	if o > 0 {
		i.PopInterval = time.Duration(o)
	}
}

func WithFetchInterval(d time.Duration) Option {
	return fetchInterval(d)
}

type addr string

func (o addr) Apply(i *Options) {
	i.Addr = string(o)
}

func WithQueueAddr(d string) Option {
	return addr(d)
}

type name string

func (o name) Apply(i *Options) {
	i.Name = string(o)
}

func WithQueueName(d string) Option {
	return name(d)
}
