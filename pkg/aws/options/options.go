package options

type Options struct {
}

type Option interface {
	Apply(*Options)
}
