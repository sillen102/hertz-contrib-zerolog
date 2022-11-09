package zerolog

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
)

var _ hlog.FullLogger = (*Logger)(nil)

// Logger is a wrapper around `zerolog.Logger` that provides an implementation of `hlog.FullLogger` interface
type Logger struct {
	log     zerolog.Logger
	out     io.Writer
	level   zerolog.Level
	options []Opt
}

// New returns a new Logger instance
func New(options ...Opt) *Logger {
	return newLogger(zerolog.New(os.Stdout), options)
}

// From returns a new Logger instance using existing zerolog log.
func From(log *zerolog.Logger, options ...Opt) *Logger {
	return newLogger(*log, options)
}

// SetLevel setting logging level for logger
func (l *Logger) SetLevel(level hlog.Level) {
	lvl := matchHlogLevel(level)
	l.level = lvl
	l.log = l.log.Level(lvl)
}

// SetOutput setting output for logger
func (l *Logger) SetOutput(writer io.Writer) {
	l.out = writer
	l.log = l.log.Output(writer)
}

// Unwrap returns the underlying zerolog logger
func (l *Logger) Unwrap() *zerolog.Logger {
	return &l.log
}

func newLogger(log zerolog.Logger, options []Opt) *Logger {
	opts := newOptions(log, options)

	return &Logger{
		log:     opts.context.Logger(),
		out:     nil,
		level:   opts.level,
		options: options,
	}
}

// Log log using zerolog logger with specified level
func (l *Logger) Log(level hlog.Level, kvs ...interface{}) {
	switch level {
	case hlog.LevelTrace, hlog.LevelDebug:
		l.log.Debug().Msg(fmt.Sprint(kvs...))
	case hlog.LevelInfo:
		l.log.Info().Msg(fmt.Sprint(kvs...))
	case hlog.LevelNotice, hlog.LevelWarn:
		l.log.Warn().Msg(fmt.Sprint(kvs...))
	case hlog.LevelError:
		l.log.Error().Msg(fmt.Sprint(kvs...))
	case hlog.LevelFatal:
		l.log.Fatal().Msg(fmt.Sprint(kvs...))
	default:
		l.log.Warn().Msg(fmt.Sprint(kvs...))
	}
}

// Logf log using zerolog logger with specified level and formatting
func (l *Logger) Logf(level hlog.Level, format string, kvs ...interface{}) {
	switch level {
	case hlog.LevelTrace, hlog.LevelDebug:
		l.log.Debug().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelInfo:
		l.log.Info().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelNotice, hlog.LevelWarn:
		l.log.Warn().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelError:
		l.log.Error().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelFatal:
		l.log.Fatal().Msg(fmt.Sprintf(format, kvs...))
	default:
		l.log.Warn().Msg(fmt.Sprintf(format, kvs...))
	}
}

// CtxLogf log using zerolog logger with specified level, formatting and context
func (l *Logger) CtxLogf(level hlog.Level, ctx context.Context, format string, kvs ...interface{}) {
	ctx = l.log.WithContext(ctx)
	switch level {
	case hlog.LevelTrace, hlog.LevelDebug:
		zerolog.Ctx(ctx).Debug().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelInfo:
		zerolog.Ctx(ctx).Info().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelNotice, hlog.LevelWarn:
		zerolog.Ctx(ctx).Warn().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelError:
		zerolog.Ctx(ctx).Error().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelFatal:
		zerolog.Ctx(ctx).Fatal().Msg(fmt.Sprintf(format, kvs...))
	default:
		zerolog.Ctx(ctx).Warn().Msg(fmt.Sprintf(format, kvs...))
	}
}

func (l *Logger) Trace(v ...interface{}) {
	l.Log(hlog.LevelTrace, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Log(hlog.LevelDebug, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Log(hlog.LevelInfo, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.Log(hlog.LevelNotice, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.Log(hlog.LevelWarn, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Log(hlog.LevelError, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Log(hlog.LevelFatal, v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Logf(hlog.LevelTrace, format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(hlog.LevelDebug, format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(hlog.LevelInfo, format, v...)
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	l.Logf(hlog.LevelWarn, format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logf(hlog.LevelWarn, format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(hlog.LevelError, format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logf(hlog.LevelError, format, v...)
}

func (l *Logger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelTrace, ctx, format, v...)
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelDebug, ctx, format, v...)
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelInfo, ctx, format, v...)
}

func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelNotice, ctx, format, v...)
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelWarn, ctx, format, v...)
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelFatal, ctx, format, v...)
}
