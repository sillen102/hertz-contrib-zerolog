package zerolog

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestMatchHlogLevel(t *testing.T) {
	assert.Equal(t, zerolog.TraceLevel, matchHlogLevel(hlog.LevelTrace))
	assert.Equal(t, zerolog.DebugLevel, matchHlogLevel(hlog.LevelDebug))
	assert.Equal(t, zerolog.InfoLevel, matchHlogLevel(hlog.LevelInfo))
	assert.Equal(t, zerolog.WarnLevel, matchHlogLevel(hlog.LevelWarn))
	assert.Equal(t, zerolog.ErrorLevel, matchHlogLevel(hlog.LevelError))
	assert.Equal(t, zerolog.FatalLevel, matchHlogLevel(hlog.LevelFatal))
}

func TestMatchZerologLevel(t *testing.T) {
	assert.Equal(t, hlog.LevelTrace, matchZerologLevel(zerolog.TraceLevel))
	assert.Equal(t, hlog.LevelDebug, matchZerologLevel(zerolog.DebugLevel))
	assert.Equal(t, hlog.LevelInfo, matchZerologLevel(zerolog.InfoLevel))
	assert.Equal(t, hlog.LevelWarn, matchZerologLevel(zerolog.WarnLevel))
	assert.Equal(t, hlog.LevelError, matchZerologLevel(zerolog.ErrorLevel))
	assert.Equal(t, hlog.LevelFatal, matchZerologLevel(zerolog.FatalLevel))
}
