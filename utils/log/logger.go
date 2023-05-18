package log

import (
	"fmt"

	"github.com/mboufous/k-island/utils/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TOOD: add more config, log rotation, file logging, ...

var (
	base   = zap.NewNop()
	logger = base.Sugar()
)

type ZapLogger struct {
	Base *zap.Logger
	Log  *zap.SugaredLogger
}

// New initializes the logger with the specified log level and output file.
func NewZapLogger(config *config.Log) error {

	// Create new zap config
	zapConfig := zap.NewProductionConfig()
	zapConfig.Sampling = nil

	// Set Log Level
	var logLevel zapcore.Level
	if err := logLevel.Set(config.Level); err != nil {
		return fmt.Errorf("could not determine log level: %w", err)
	}
	zapConfig.Level.SetLevel(logLevel)

	zapConfig.Encoding = "console"

	// Enable colors
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zapConfig.DisableStacktrace = config.DisableStacktrace

	// Use sane timestamp when logging to console
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// JSON Fields
	zapConfig.EncoderConfig.MessageKey = "msg"
	zapConfig.EncoderConfig.LevelKey = "level"
	zapConfig.EncoderConfig.CallerKey = "caller"

	//Dev mode
	zapConfig.Development = config.DevMode

	// Build the Logger
	globalLogger, err := zapConfig.Build()
	if err != nil {
		return fmt.Errorf("could not build log config: %w", err)
	}
	zap.ReplaceGlobals(globalLogger)

	base = zap.L()
	logger = base.WithOptions(zap.AddCallerSkip(1)).Sugar()

	return nil
}
