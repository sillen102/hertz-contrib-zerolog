package hertz_contrib_zerolog

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFrom(t *testing.T) {
	b := &bytes.Buffer{}

	zl := zerolog.New(b)
	l := From(zl.With().Str("key", "test").Logger())

	l.Info("foo")

	assert.Equal(
		t,
		`{"level":"info","key":"test","message":"foo"}
`,
		b.String(),
	)
}

func TestLog(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(b)

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
	l := New(b)

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
	ctx := context.Background()
	b := &bytes.Buffer{}
	l := New(b)

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
	ctx := context.Background()
	b := &bytes.Buffer{}
	l := New(b)

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
	ctx := context.Background()
	b := &bytes.Buffer{}
	l := New(b)

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
	ctx := context.Background()
	b := &bytes.Buffer{}
	l := New(b)

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
	ctx := context.Background()
	b := &bytes.Buffer{}
	l := New(b)

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
	ctx := context.Background()
	b := &bytes.Buffer{}
	l := New(b)

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
	l := New(os.Stdout)

	l.SetLevel(hlog.LevelDebug)
	assert.Equal(t, l.log.GetLevel(), zerolog.DebugLevel)

	l.SetLevel(hlog.LevelDebug)
	assert.Equal(t, l.log.GetLevel(), zerolog.DebugLevel)

	l.SetLevel(hlog.LevelError)
	assert.Equal(t, l.log.GetLevel(), zerolog.ErrorLevel)
}

func TestSetOutput(t *testing.T) {
	b1 := &bytes.Buffer{}
	l := New(os.Stdout)
	l.SetOutput(b1)

	l.Info("foo")
	assert.Equal(
		t,
		`{"level":"info","message":"foo"}
`,
		b1.String(),
	)

	b2 := &bytes.Buffer{}
	l.SetOutput(b2)
	l.Debug("bar")
	assert.Equal(
		t,
		`{"level":"debug","message":"bar"}
`,
		b2.String(),
	)
}
