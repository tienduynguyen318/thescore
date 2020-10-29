package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var String = zap.String

type zLogger struct {
	l *zap.SugaredLogger
}

func NewZapLogger(config Config) (Logger, error) {
	if _, err := config.valid(); err != nil {
		return nil, err
	}

	writer := zapcore.Lock(zapcore.AddSync(config.Writer))
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "@timestamp"
	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "log.level"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	logger := zap.New(
		core,
		zap.Fields(
			zap.String("ecs.version", "1.5.0"),
			zap.String("service.name", config.AppName),
			zap.String("service.type", config.AppName),
			zap.String("service.version", config.AppVersion),
			zap.String("host.hostname", config.Hostname),
		),
	)

	return &zLogger{
		l: logger.Sugar(),
	}, nil
}

func (zl *zLogger) Debugf(tmp string, args ...interface{}) {
	zl.l.Debugf(tmp, args...)
}

func (zl *zLogger) Infof(tmp string, args ...interface{}) {
	zl.l.Infof(tmp, args...)
}

func (zl *zLogger) Errorf(tmp string, args ...interface{}) {
	zl.l.Errorf(tmp, args...)
}

func (zl *zLogger) Warnf(tmp string, args ...interface{}) {
	zl.l.Warnf(tmp, args...)
}

func (zl *zLogger) Fatalf(tmp string, args ...interface{}) {
	zl.l.Fatalf(tmp, args...)
}

func (zl *zLogger) Panicf(tmp string, args ...interface{}) {
	zl.l.Panicf(tmp, args...)
}

func (zl *zLogger) WithFields(args ...interface{}) Logger {
	return &zLogger{l: zl.l.With(args...)}
}
