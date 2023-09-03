package slogex

import (
	"context"
	"os"

	"log/slog"
)

// write debug level and omit timestamp for testing output easily
func newExampleTextHandler() *slog.TextHandler {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
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
func newExampleJSONHandler() *slog.JSONHandler {
	return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
}

func ExampleInitSlogWithTextHandler() {
	logger := slog.New(newExampleTextHandler())
	logger.Info("Hello from slog logger!")
	// Output:
	// level=INFO msg="Hello from slog logger!"
}

func ExampleInitSlogWithJsonHandler() {
	logger := slog.New(newExampleJSONHandler())
	logger.Info("Hello from slog logger!")
	// Output:
	// {"level":"INFO","msg":"Hello from slog logger!"}
}

func ExampleSlogLoggerLog() {
	logger := slog.New(newExampleJSONHandler())

	logger.Debug("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.Info("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.Warn("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.Error("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	// Output:
	// {"level":"DEBUG","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"INFO","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"WARN","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"ERROR","msg":"Amikai info","age":18,"id":"123","is married":false}
}

func ExampleSlogLoggerLogWithoutAttr() {
	logger := slog.New(newExampleJSONHandler())

	logger.Info("Amikai info",
		"age", 18,
		"id", "123",
		"is married", false,
	)
	// Output:
	// {"level":"INFO","msg":"Amikai info","age":18,"id":"123","is married":false}
}

func ExampleSlogLoggerLogAttr() {
	logger := slog.New(newExampleJSONHandler())

	logger.LogAttrs(context.Background(), slog.LevelDebug, "Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.LogAttrs(context.Background(), slog.LevelInfo, "Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.LogAttrs(context.Background(), slog.LevelWarn, "Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	logger.LogAttrs(context.Background(), slog.LevelError, "Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	// Output:
	// {"level":"DEBUG","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"INFO","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"WARN","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"ERROR","msg":"Amikai info","age":18,"id":"123","is married":false}
}

func ExampleGroupAttr() {
	logger := slog.New(newExampleJSONHandler())
	logger.Info("Usage Statistics",
		slog.Group("memory",
			slog.Int("current", 50),
			slog.Int("min", 20),
			slog.Int("max", 80)),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)
	// Output:
	// {"level":"INFO","msg":"Usage Statistics","memory":{"current":50,"min":20,"max":80},"cpu":10,"app-version":"v0.0.1-beta"}
}

func ExampleAddingContextByLogger() {
	logger := slog.New(newExampleJSONHandler())
	childLogger := logger.With(slog.String("service", "myService"), slog.String("requestID", "123"))

	childLogger.Info("user registration successful",
		slog.String("username", "john.doe"),
		slog.String("email", "john@example.com"),
	)

	childLogger.Info("redirecting user to admin dashboard")
	// Output:
	//{"level":"INFO","msg":"user registration successful","service":"myService","requestID":"123","username":"john.doe","email":"john@example.com"}
	//{"level":"INFO","msg":"redirecting user to admin dashboard","service":"myService","requestID":"123"}
}

func ExampleUseAndReplaceGlobalLogger() {
	logger := slog.New(newExampleJSONHandler())

	// Replace global logger
	slog.SetDefault(logger)

	// get the default logger
	_ = slog.Default()

	// Call global log method
	slog.Debug("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	slog.Info("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	slog.Warn("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	slog.Error("Amikai info", slog.Int("age", 18), slog.String("id", "123"), slog.Bool("is married", false))
	// Output:
	// {"level":"DEBUG","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"INFO","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"WARN","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"ERROR","msg":"Amikai info","age":18,"id":"123","is married":false}
}
