package zerolog

import (
	"bytes"
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	b := &bytes.Buffer{}

	zl := zerolog.New(b).With().Str("key", "test").Logger()
	l := From(&zl)

	l.Info("foo")

	assert.Equal(
		t,
		`{"level":"info","key":"test","message":"foo"}
`,
		b.String(),
	)
}

func TestUnwrap(t *testing.T) {
	l := New()

	logger := l.Unwrap()

	assert.NotNil(t, logger)
	assert.IsType(t, &zerolog.Logger{}, logger)
}

func TestLog(t *testing.T) {
	b := &bytes.Buffer{}
	l := New()
	l.SetOutput(b)

	l.Trace("foo")
	assert.Equal(
		t,
		`{"level":"debug","message":"foo"}
`,
		b.String(),
	)

	b.Reset()
	l.Debug("foo")
	assert.Equal(
		t,
		`{"level":"debug","message":"foo"}
`,
		b.String(),
	)

	b.Reset()
	l.Info("foo")
	assert.Equal(
		t,
		`{"level":"info","message":"foo"}
`,
		b.String(),
	)

	b.Reset()
	l.Notice("foo")
	assert.Equal(
		t,
		`{"level":"warn","message":"foo"}
`,
		b.String(),
	)

	b.Reset()
	l.Warn("foo")
	assert.Equal(
		t,
		`{"level":"warn","message":"foo"}
`,
		b.String(),
	)

	b.Reset()
	l.Error("foo")
	assert.Equal(
		t,
		`{"level":"error","message":"foo"}
`,
		b.String(),
	)
}

func TestLogf(t *testing.T) {
	b := &bytes.Buffer{}
	l := New()
	l.SetOutput(b)

	l.Tracef("foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"debug","message":"foobar"}
`,
		b.String(),
	)

	b.Reset()
	l.Debugf("foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"debug","message":"foobar"}
`,
		b.String(),
	)

	b.Reset()
	l.Infof("foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"info","message":"foobar"}
`,
		b.String(),
	)

	b.Reset()
	l.Noticef("foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"warn","message":"foobar"}
`,
		b.String(),
	)

	b.Reset()
	l.Warnf("foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"warn","message":"foobar"}
`,
		b.String(),
	)

	b.Reset()
	l.Errorf("foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"error","message":"foobar"}
`,
		b.String(),
	)
}

func TestCtxTracef(t *testing.T) {
	b := &bytes.Buffer{}
	l := New()
	l.SetOutput(b)
	ctx := l.log.WithContext(context.Background())

	l.CtxTracef(ctx, "foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"debug","message":"foobar"}
`,
		b.String(),
	)
	assert.NotNil(t, log.Ctx(ctx))
}

func TestCtxDebugf(t *testing.T) {
	b := &bytes.Buffer{}
	l := New()
	l.SetOutput(b)
	ctx := l.log.WithContext(context.Background())

	l.CtxDebugf(ctx, "foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"debug","message":"foobar"}
`,
		b.String(),
	)
	assert.NotNil(t, log.Ctx(ctx))
}

func TestCtxInfof(t *testing.T) {
	b := &bytes.Buffer{}
	l := New()
	l.SetOutput(b)
	ctx := l.log.WithContext(context.Background())

	l.CtxInfof(ctx, "foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"info","message":"foobar"}
`,
		b.String(),
	)
	assert.NotNil(t, log.Ctx(ctx))
}

func TestCtxNoticef(t *testing.T) {
	b := &bytes.Buffer{}
	l := New()
	l.SetOutput(b)
	ctx := l.log.WithContext(context.Background())

	l.CtxNoticef(ctx, "foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"warn","message":"foobar"}
`,
		b.String(),
	)
	assert.NotNil(t, log.Ctx(ctx))
}

func TestCtxWarnf(t *testing.T) {
	b := &bytes.Buffer{}
	l := New()
	l.SetOutput(b)
	ctx := l.log.WithContext(context.Background())

	l.CtxWarnf(ctx, "foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"warn","message":"foobar"}
`,
		b.String(),
	)
	assert.NotNil(t, log.Ctx(ctx))
}

func TestCtxErrorf(t *testing.T) {
	b := &bytes.Buffer{}
	l := New()
	l.SetOutput(b)
	ctx := l.log.WithContext(context.Background())

	l.CtxErrorf(ctx, "foo%s", "bar")
	assert.Equal(
		t,
		`{"level":"error","message":"foobar"}
`,
		b.String(),
	)
	assert.NotNil(t, log.Ctx(ctx))
}

func TestSetLevel(t *testing.T) {
	l := New()

	l.SetLevel(hlog.LevelDebug)
	assert.Equal(t, l.log.GetLevel(), zerolog.DebugLevel)

	l.SetLevel(hlog.LevelDebug)
	assert.Equal(t, l.log.GetLevel(), zerolog.DebugLevel)

	l.SetLevel(hlog.LevelError)
	assert.Equal(t, l.log.GetLevel(), zerolog.ErrorLevel)
}
