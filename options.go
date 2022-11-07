package hertz_contrib_zerolog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
)

type (
	Options struct {
		context zerolog.Context
		level   zerolog.Level
	}

	Setter func(opts *Options)
)

func newOptions(log zerolog.Logger, setters []Setter) *Options {
	opts := &Options{
		context: log.With(),
		level:   log.GetLevel(),
	}

	for _, set := range setters {
		set(opts)
	}

	return opts
}

func WithLevel(level hlog.Level) Setter {
	lvl := MatchHlogLevel(level)
	return func(opts *Options) {
		opts.context = opts.context.Logger().Level(lvl).With()
		opts.level = lvl
	}
}

func WithField(name string, value interface{}) Setter {
	return func(opts *Options) {
		opts.context = opts.context.Interface(name, value)
	}
}

func WithFields(fields map[string]interface{}) Setter {
	return func(opts *Options) {
		opts.context = opts.context.Fields(fields)
	}
}

func WithTimestamp() Setter {
	return func(opts *Options) {
		opts.context = opts.context.Timestamp()
	}
}

func WithCaller() Setter {
	return func(opts *Options) {
		opts.context = opts.context.Caller()
	}
}

func WithHook(hook zerolog.Hook) Setter {
	return func(opts *Options) {
		opts.context = opts.context.Logger().Hook(hook).With()
	}
}

func WithHookFunc(hook zerolog.HookFunc) Setter {
	return func(opts *Options) {
		opts.context = opts.context.Logger().Hook(hook).With()
	}
}
