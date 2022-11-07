package hertz_contrib_zerolog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
)

var (
	zerologLevels = map[hlog.Level]zerolog.Level{
		hlog.LevelTrace:  zerolog.TraceLevel,
		hlog.LevelDebug:  zerolog.DebugLevel,
		hlog.LevelInfo:   zerolog.InfoLevel,
		hlog.LevelWarn:   zerolog.WarnLevel,
		hlog.LevelNotice: zerolog.WarnLevel,
		hlog.LevelError:  zerolog.ErrorLevel,
		hlog.LevelFatal:  zerolog.FatalLevel,
	}

	hlogLevels = map[zerolog.Level]hlog.Level{
		zerolog.TraceLevel: hlog.LevelTrace,
		zerolog.DebugLevel: hlog.LevelDebug,
		zerolog.InfoLevel:  hlog.LevelInfo,
		zerolog.WarnLevel:  hlog.LevelWarn,
		zerolog.ErrorLevel: hlog.LevelError,
		zerolog.FatalLevel: hlog.LevelFatal,
	}
)

// MatchHlogLevel map hlog.Level to zerolog.Level
func MatchHlogLevel(level hlog.Level) zerolog.Level {
	zlvl, found := zerologLevels[level]

	if found {
		return zlvl
	}

	return zerolog.WarnLevel // Default level
}

// MatchZerologLevel map zerolog.Level to hlog.Level
func MatchZerologLevel(level zerolog.Level) hlog.Level {
	hlvl, found := hlogLevels[level]

	if found {
		return hlvl
	}

	return hlog.LevelWarn // Default level
}
