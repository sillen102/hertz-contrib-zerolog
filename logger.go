package hertz_contrib_zerolog

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
)

var _ hlog.FullLogger = (*Logger)(nil)

type Logger struct {
	log   zerolog.Logger
	out   io.Writer
	level zerolog.Level
}

func NewLogger(out io.Writer, level zerolog.Level) *Logger {

	switch l := out.(type) {
	case zerolog.Logger:
		return newLogger(l, level)
	default:
		return newLogger(zerolog.New(out), level)
	}
}

func newLogger(log zerolog.Logger, level zerolog.Level) *Logger {
	opts := newOptions(log)

	return &Logger{
		log:   opts.context.Logger(),
		out:   nil,
		level: level,
	}
}

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

func (l *Logger) CtxLogf(level hlog.Level, ctx context.Context, format string, kvs ...interface{}) {
	ctx = l.log.WithContext(ctx)
	switch level {
	case hlog.LevelTrace, hlog.LevelDebug:
		log.Ctx(ctx).Debug().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelInfo:
		log.Ctx(ctx).Info().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelNotice, hlog.LevelWarn:
		log.Ctx(ctx).Warn().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelError:
		log.Ctx(ctx).Error().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelFatal:
		log.Ctx(ctx).Fatal().Msg(fmt.Sprintf(format, kvs...))
	default:
		log.Ctx(ctx).Warn().Msg(fmt.Sprintf(format, kvs...))
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
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) SetLevel(level hlog.Level) {
	var lvl zerolog.Level
	switch level {
	case hlog.LevelTrace, hlog.LevelDebug:
		lvl = zerolog.DebugLevel
	case hlog.LevelInfo:
		lvl = zerolog.InfoLevel
	case hlog.LevelWarn, hlog.LevelNotice:
		lvl = zerolog.WarnLevel
	case hlog.LevelError:
		lvl = zerolog.ErrorLevel
	case hlog.LevelFatal:
		lvl = zerolog.FatalLevel
	default:
		lvl = zerolog.WarnLevel
	}
	l.log.Level(lvl)
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.log.Output(writer)
}
