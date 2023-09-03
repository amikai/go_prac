package zapex

import (
	"testing"

	"log/slog"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

// Learn from https://betterstack.com/community/guides/logging/go/zap/

func ExampleInitZap() {
	// Example write debug level and omit timestamp
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("Hello from Zap logger!")

	// Output:
	// {"level":"info","msg":"Hello from Zap logger!"}
}

func ExampleLoggerLog() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Debug("Amikai info", zap.Int("age", 18), zap.String("id", "123"), zap.Bool("is married", false))
	logger.Info("Amikai info", zap.Int("age", 18), zap.String("id", "123"), zap.Bool("is married", false))
	logger.Warn("Amikai info", zap.Int("age", 18), zap.String("id", "123"), zap.Bool("is married", false))
	logger.Error("Amikai info", zap.Int("age", 18), zap.String("id", "123"), zap.Bool("is married", false))
	// Output:
	// {"level":"debug","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"info","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"warn","msg":"Amikai info","age":18,"id":"123","is married":false}
	// {"level":"error","msg":"Amikai info","age":18,"id":"123","is married":false}
}

func ExampleMust() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
}

func ExampleZapGlobalLogger() {
	// replace global logger
	zap.ReplaceGlobals(zap.NewExample())
	// use global logger
	logger := zap.L()
	logger.Info("Hello from Zap logger!")
	// Output:
	// {"level":"info","msg":"Hello from Zap logger!"}
}

func ExampleSuggar() {
	sl := zap.NewExample().Sugar()
	defer sl.Sync()

	sl.Info("Hello", "World!")
	sl.Infof("%s,%s!", "Hello", "World")
	sl.Infow("Amikai info",
		"age", 18,
		"id", "123",
		"is married", false,
	)
	// Output:
	// {"level":"info","msg":"HelloWorld!"}
	// {"level":"info","msg":"Hello,World!"}
	// {"level":"info","msg":"Amikai info","age":18,"id":"123","is married":false}
}

func ExampleAddingContext() {
	logger := zap.NewExample()
	childLogger := logger.With(zap.String("service", "myService"), zap.String("requestID", "123"))

	childLogger.Info("user registration successful",
		zap.String("username", "john.doe"),
		zap.String("email", "john@example.com"),
	)

	childLogger.Info("redirecting user to admin dashboard")
	// Output:
	//{"level":"info","msg":"user registration successful","service":"myService","requestID":"123","username":"john.doe","email":"john@example.com"}
	//{"level":"info","msg":"redirecting user to admin dashboard","service":"myService","requestID":"123"}
}

func ExampleZapSlogBackend() {
	logger := zap.NewExample()

	defer logger.Sync()

	sl := slog.New(zapslog.NewHandler(logger.Core(), nil))

	sl.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.String("path", "/api/user"),
		slog.Int("status", 200),
	)
	// Output:
	// {"level":"info","msg":"incoming request","method":"GET","path":"/api/user","status":200}
}

func TestExampleDesuggar(t *testing.T) {
	logger := zap.NewExample()
	sugar := logger.Sugar()
	logger2 := sugar.Desugar()

	assert.IsType(t, &zap.Logger{}, logger)
	assert.IsType(t, &zap.SugaredLogger{}, sugar)
	assert.IsType(t, &zap.Logger{}, logger2)
}
