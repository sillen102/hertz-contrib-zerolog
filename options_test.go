package hertzZerolog

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestWithCaller(t *testing.T) {
	b := &bytes.Buffer{}

	l := New(b, WithCaller())

	l.Info("foobar")

	type Log struct {
		Level   string `json:"level"`
		Caller  string `json:"caller"`
		Message string `json:"message"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)

	segments := strings.Split(log.Caller, ":")
	filePath := filepath.Base(segments[0])

	assert.Equal(t, filePath, "logger.go")
}

func TestWithField(t *testing.T) {
	b := &bytes.Buffer{}

	l := New(b, WithField("service", "logging"))

	l.Info("foobar")

	type Log struct {
		Level   string `json:"level"`
		Service string `json:"service"`
		Message string `json:"message"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.Equal(t, log.Service, "logging")
}

func TestWithFields(t *testing.T) {
	b := &bytes.Buffer{}

	l := New(b, WithFields(map[string]interface{}{
		"host": "localhost",
		"port": 8080,
	}))

	l.Info("foobar")

	type Log struct {
		Level   string `json:"level"`
		Host    string `json:"host"`
		Port    int    `json:"port"`
		Message string `json:"message"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.Equal(t, log.Host, "localhost")
	assert.Equal(t, log.Port, 8080)
}

type (
	Hook struct {
		logs []HookLog
	}

	HookLog struct {
		level   zerolog.Level
		message string
	}
)

func (h *Hook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	h.logs = append(h.logs, HookLog{
		level:   level,
		message: message,
	})
}

func TestWithHook(t *testing.T) {
	b := &bytes.Buffer{}
	h := &Hook{}
	l := New(b, WithHook(h))

	l.Info("Foo")
	l.Warn("Bar")

	assert.Len(t, h.logs, 2)
	assert.Equal(t, h.logs[0].level, zerolog.InfoLevel)
	assert.Equal(t, h.logs[0].message, "Foo")
	assert.Equal(t, h.logs[1].level, zerolog.WarnLevel)
	assert.Equal(t, h.logs[1].message, "Bar")
}

func TestWithHookFunc(t *testing.T) {
	b := &bytes.Buffer{}
	logs := make([]HookLog, 0, 2)
	l := New(b, WithHookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
		logs = append(logs, HookLog{
			level:   level,
			message: message,
		})
	}))

	l.Info("Foo")
	l.Warn("Bar")

	assert.Len(t, logs, 2)
	assert.Equal(t, logs[0].level, zerolog.InfoLevel)
	assert.Equal(t, logs[0].message, "Foo")
	assert.Equal(t, logs[1].level, zerolog.WarnLevel)
	assert.Equal(t, logs[1].message, "Bar")
}

func TestWithLevel(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(b, WithLevel(hlog.LevelInfo))

	l.Debug("Test")

	assert.Equal(t, b.String(), "")

	l.Info("foobar")

	assert.Equal(t, b.String(), `{"level":"info","message":"foobar"}
`)
}

func TestWithTimestamp(t *testing.T) {
	b := &bytes.Buffer{}

	l := New(b, WithTimestamp())

	l.Info("foobar")

	type Log struct {
		Level   string    `json:"level"`
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.NotEmpty(t, log.Time)
}

func TestWithFormattedTimestamp(t *testing.T) {
	b := &bytes.Buffer{}

	l := New(b, WithFormattedTimestamp(time.RFC3339Nano))

	l.Info("foobar")

	type Log struct {
		Level   string    `json:"level"`
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.NotEmpty(t, log.Time)
}
