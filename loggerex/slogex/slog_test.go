package slogex

import (
	"context"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// write debug level and omit timestamp for testing output easily
func newExampleTextHandler(buf *strings.Builder) *slog.TextHandler {
	return slog.NewTextHandler(buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
}

// write debug level and omit timestamp for testing output easily
func newExampleJSONHandler(buf *strings.Builder) *slog.JSONHandler {
	return slog.NewJSONHandler(buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
}

func TestInitSlogWithTextHandler(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(newExampleTextHandler(&buf))
	logger.Info("Hello from slog logger!")
	want := `level=INFO msg="Hello from slog logger!"
`
	assert.Equal(t, want, buf.String())
}

func TestInitSlogWithJsonHandler(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(newExampleJSONHandler(&buf))
	logger.Info("Hello from slog logger!")
	want := `{"level":"INFO","msg":"Hello from slog logger!"}
`
	assert.Equal(t, want, buf.String())
}

func TestSlogLoggerLog(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(newExampleJSONHandler(&buf))

	logger.Debug("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.Info("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.Warn("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.Error("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	want := `{"level":"DEBUG","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"INFO","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"WARN","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"ERROR","msg":"Amikai info","age":18,"id":"123","is married":false}
`
	assert.Equal(t, want, buf.String())
}

func TestSlogLoggerLogWithoutAttr(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(newExampleJSONHandler(&buf))

	logger.Info("Amikai info",
		"age", 18,
		"id", "123",
		"is married", false,
	)
	want := `{"level":"INFO","msg":"Amikai info","age":18,"id":"123","is married":false}
`
	assert.Equal(t, want, buf.String())
}

func TestSlogLoggerLogAttr(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(newExampleJSONHandler(&buf))

	logger.LogAttrs(context.Background(), slog.LevelDebug, "Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.LogAttrs(context.Background(), slog.LevelInfo, "Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.LogAttrs(context.Background(), slog.LevelWarn, "Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.LogAttrs(context.Background(), slog.LevelError, "Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	want := `{"level":"DEBUG","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"INFO","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"WARN","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"ERROR","msg":"Amikai info","age":18,"id":"123","is married":false}
`
	assert.Equal(t, want, buf.String())
}

func TestGroupAttr(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(newExampleJSONHandler(&buf))
	logger.Info("Usage Statistics",
		slog.Group("memory",
			slog.Int("current", 50),
			slog.Int("min", 20),
			slog.Int("max", 80)),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)
	want := `{"level":"INFO","msg":"Usage Statistics","memory":{"current":50,"min":20,"max":80},"cpu":10,"app-version":"v0.0.1-beta"}
`
	assert.Equal(t, want, buf.String())
}

func TestAddingContextByLogger(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(newExampleJSONHandler(&buf))
	childLogger := logger.With(slog.String("service", "myService"), slog.String("requestID", "123"))

	childLogger.Info("user registration successful",
		slog.String("username", "john.doe"),
		slog.String("email", "john@example.com"),
	)
	want := `{"level":"INFO","msg":"user registration successful","service":"myService","requestID":"123","username":"john.doe","email":"john@example.com"}
`
	assert.Equal(t, want, buf.String())

	buf.Reset()
	childLogger.Info("redirecting user to admin dashboard")
	want = `{"level":"INFO","msg":"redirecting user to admin dashboard","service":"myService","requestID":"123"}
`
	assert.Equal(t, want, buf.String())
}

func TestUseAndReplaceGlobalLogger(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(newExampleJSONHandler(&buf))

	// Replace global logger
	slog.SetDefault(logger)

	// get the default logger
	_ = slog.Default()

	// Call global log method
	slog.Debug("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	slog.Info("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	slog.Warn("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	slog.Error("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))

	want := `{"level":"DEBUG","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"INFO","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"WARN","msg":"Amikai info","age":18,"id":"123","is married":false}
{"level":"ERROR","msg":"Amikai info","age":18,"id":"123","is married":false}
`
	assert.Equal(t, want, buf.String())
}
