package hertzZerolog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
)

type (
	Options struct {
		context zerolog.Context
		level   zerolog.Level
	}

	Opt func(opts *Options)
)

func newOptions(log zerolog.Logger, options []Opt) *Options {
	opts := &Options{
		context: log.With(),
		level:   log.GetLevel(),
	}

	for _, set := range options {
		set(opts)
	}

	return opts
}

func WithLevel(level hlog.Level) Opt {
	lvl := MatchHlogLevel(level)
	return func(opts *Options) {
		opts.context = opts.context.Logger().Level(lvl).With()
		opts.level = lvl
	}
}

func WithField(name string, value interface{}) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Interface(name, value)
	}
}

func WithFields(fields map[string]interface{}) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Fields(fields)
	}
}

func WithTimestamp() Opt {
	return func(opts *Options) {
		opts.context = opts.context.Timestamp()
	}
}

// WithFormattedTimestamp adds a timestamp field and sets the zerolog.TimeFieldFormat format for the zerolog logger
func WithFormattedTimestamp(format string) Opt {
	zerolog.TimeFieldFormat = format
	return func(opts *Options) {
		opts.context = opts.context.Timestamp()
	}
}

func WithCaller() Opt {
	return func(opts *Options) {
		opts.context = opts.context.Caller()
	}
}

func WithHook(hook zerolog.Hook) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Logger().Hook(hook).With()
	}
}

func WithHookFunc(hook zerolog.HookFunc) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Logger().Hook(hook).With()
	}
}
