package zapex

import (
	"io"
	"strings"
	"testing"

	"log/slog"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

// Learn from https://betterstack.com/community/guides/logging/go/zap/

func newExampleLogger(w io.Writer) *zap.Logger {
	// copy from zap.NewExample()
	// Example write debug level and omit timestamp
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.AddSync(w), zapcore.DebugLevel)
	return zap.New(core)

}

func TestInitZap(t *testing.T) {
	buf := strings.Builder{}
	logger := newExampleLogger(&buf)
	defer logger.Sync()

	logger.Info("Hello from Zap logger!")

	want := `{"level":"info","msg":"Hello from Zap logger!"}
`
	assert.Equal(t, want, buf.String())
}

func TestLoggerLog(t *testing.T) {
	buf := strings.Builder{}
	logger := newExampleLogger(&buf)
	defer logger.Sync()

	logger.Debug("Amikai info", zap.Int("age", 18), zap.String("id", "123"), zap.Bool("is married", false))
	logger.Info("Amikai info", zap.Int("age", 18), zap.String("id", "123"), zap.Bool("is married", false))
	logger.Warn("Amikai info", zap.Int("age", 18), zap.String("id", "123"), zap.Bool("is married", false))
	logger.Error("Amikai info", zap.Int("age", 18), zap.String("id", "123"), zap.Bool("is married", false))
	want := `{"level":"debug","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"info","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"warn","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"error","msg":"Amikai info","age":18,"id":"123","is married":false}
`
	assert.Equal(t, want, buf.String())
}

func TestMust(t *testing.T) {
	assert.NotPanics(t, func() {
		_ = zap.Must(zap.NewProduction())
	})
}

func TestZapGlobalLogger(t *testing.T) {
	buf := strings.Builder{}
	exampleLogger := newExampleLogger(&buf)
	// replace global logger
	zap.ReplaceGlobals(exampleLogger)
	// use global logger
	logger := zap.L()
	logger.Info("Hello from Zap logger!")

	want := `{"level":"info","msg":"Hello from Zap logger!"}
`

	assert.Equal(t, want, buf.String())
}

func TestSuggar(t *testing.T) {
	buf := strings.Builder{}
	sl := newExampleLogger(&buf).Sugar()
	defer sl.Sync()

	sl.Info("Hello", "World!")
	sl.Infof("%s,%s!", "Hello", "World")
	sl.Infow("Amikai info",
		"age", 18,
		"id", "123",
		"is married", false,
	)
	want := `{"level":"info","msg":"HelloWorld!"}
{"level":"info","msg":"Hello,World!"}
{"level":"info","msg":"Amikai info","age":18,"id":"123","is married":false}
`
	assert.Equal(t, want, buf.String())
}

func TestAddingContext(t *testing.T) {
	buf := strings.Builder{}
	logger := newExampleLogger(&buf)
	childLogger := logger.With(zap.String("service", "myService"), zap.String("requestID", "123"))

	childLogger.Info("user registration successful",
		zap.String("username", "john.doe"),
		zap.String("email", "john@example.com"),
	)

	childLogger.Info("redirecting user to admin dashboard")
	want := `{"level":"info","msg":"user registration successful","service":"myService","requestID":"123","username":"john.doe","email":"john@example.com"}
{"level":"info","msg":"redirecting user to admin dashboard","service":"myService","requestID":"123"}
`
	assert.Equal(t, want, buf.String())
}

func TestZapSlogBackend(t *testing.T) {
	buf := strings.Builder{}
	logger := newExampleLogger(&buf)
	defer logger.Sync()

	sl := slog.New(zapslog.NewHandler(logger.Core(), nil))

	sl.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.String("path", "/api/user"),
		slog.Int("status", 200),
	)
	want := `{"level":"info","msg":"incoming request","method":"GET","path":"/api/user","status":200}
`
	assert.Equal(t, want, buf.String())
}

func TestDesuggar(t *testing.T) {
	logger := zap.NewExample()
	sugar := logger.Sugar()
	logger2 := sugar.Desugar()

	assert.IsType(t, &zap.Logger{}, logger)
	assert.IsType(t, &zap.SugaredLogger{}, sugar)
	assert.IsType(t, &zap.Logger{}, logger2)
}
