package tgram

import "go.uber.org/zap"

type Options struct {
	Logger  *zap.Logger
	Storage Storage
}

func (opts *Options) setDefaults() {
	if opts.Logger == nil {
		opts.Logger = zap.NewNop()
	}
}
