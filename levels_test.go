package hertzZerolog

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestMatchHlogLevel(t *testing.T) {
	assert.Equal(t, zerolog.TraceLevel, MatchHlogLevel(hlog.LevelTrace))
	assert.Equal(t, zerolog.DebugLevel, MatchHlogLevel(hlog.LevelDebug))
	assert.Equal(t, zerolog.InfoLevel, MatchHlogLevel(hlog.LevelInfo))
	assert.Equal(t, zerolog.WarnLevel, MatchHlogLevel(hlog.LevelWarn))
	assert.Equal(t, zerolog.ErrorLevel, MatchHlogLevel(hlog.LevelError))
	assert.Equal(t, zerolog.FatalLevel, MatchHlogLevel(hlog.LevelFatal))
}

func TestMatchZerologLevel(t *testing.T) {
	assert.Equal(t, hlog.LevelTrace, MatchZerologLevel(zerolog.TraceLevel))
	assert.Equal(t, hlog.LevelDebug, MatchZerologLevel(zerolog.DebugLevel))
	assert.Equal(t, hlog.LevelInfo, MatchZerologLevel(zerolog.InfoLevel))
	assert.Equal(t, hlog.LevelWarn, MatchZerologLevel(zerolog.WarnLevel))
	assert.Equal(t, hlog.LevelError, MatchZerologLevel(zerolog.ErrorLevel))
	assert.Equal(t, hlog.LevelFatal, MatchZerologLevel(zerolog.FatalLevel))
}
