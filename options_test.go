package zerolog

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

func TestWithOutput(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithOutput(b))

	l.Info("foobar")

	assert.Equal(
		t,
		`{"level":"info","message":"foobar"}
`,
		b.String(),
	)
}

func TestWithCaller(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithCaller())
	l.SetOutput(b)

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

	assert.Equal(t, "logger.go", filePath)
}

func TestWithField(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithField("service", "logging"))
	l.SetOutput(b)

	l.Info("foobar")

	type Log struct {
		Level   string `json:"level"`
		Service string `json:"service"`
		Message string `json:"message"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.Equal(t, "logging", log.Service)
}

func TestWithFields(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithFields(map[string]interface{}{
		"host": "localhost",
		"port": 8080,
	}))
	l.SetOutput(b)

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
	assert.Equal(t, "localhost", log.Host)
	assert.Equal(t, 8080, log.Port)
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
	l := New(WithHook(h))
	l.SetOutput(b)

	l.Info("Foo")
	l.Warn("Bar")

	assert.Len(t, h.logs, 2)
	assert.Equal(t, zerolog.InfoLevel, h.logs[0].level)
	assert.Equal(t, "Foo", h.logs[0].message)
	assert.Equal(t, zerolog.WarnLevel, h.logs[1].level)
	assert.Equal(t, "Bar", h.logs[1].message)
}

func TestWithHookFunc(t *testing.T) {
	b := &bytes.Buffer{}
	logs := make([]HookLog, 0, 2)
	l := New(WithHookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
		logs = append(logs, HookLog{
			level:   level,
			message: message,
		})
	}))
	l.SetOutput(b)

	l.Info("Foo")
	l.Warn("Bar")

	assert.Len(t, logs, 2)
	assert.Equal(t, zerolog.InfoLevel, logs[0].level)
	assert.Equal(t, "Foo", logs[0].message)
	assert.Equal(t, zerolog.WarnLevel, logs[1].level)
	assert.Equal(t, "Bar", logs[1].message)
}

func TestWithLevel(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithLevel(hlog.LevelInfo))
	l.SetOutput(b)

	l.Debug("Test")

	assert.Equal(t, b.String(), "")

	l.Info("foobar")

	assert.Equal(t, `{"level":"info","message":"foobar"}
`, b.String())
}

type Log struct {
	Level   string    `json:"level"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func TestWithTimestamp(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithTimestamp())
	l.SetOutput(b)

	l.Info("foobar")

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.NotEmpty(t, log.Time)
}

func TestWithFormattedTimestamp(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithFormattedTimestamp(time.RFC3339Nano))
	l.SetOutput(b)

	l.Info("foobar")

	log := &Log{}
	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.NotEmpty(t, log.Time)
}
