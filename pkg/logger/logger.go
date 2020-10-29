package logger

import (
	"errors"
	"io"
)

var Log Logger

type Config struct {
	Writer     io.Writer
	AppName    string
	AppVersion string
	Hostname   string
}

func (c Config) valid() (bool, error) {
	if c.Writer == nil {
		return false, errors.New("Writer is required")
	}
	if c.AppName == "" {
		return false, errors.New("AppName is required")
	}
	if c.AppVersion == "" {
		return false, errors.New("AppVersion is required")
	}
	if c.Hostname == "" {
		return false, errors.New("Hostname is required")
	}
	return true, nil
}

type Logger interface {
	Debugf(tmp string, args ...interface{})
	Infof(tmp string, args ...interface{})
	Errorf(tmp string, args ...interface{})
	Warnf(tmp string, args ...interface{})
	Fatalf(tmp string, args ...interface{})
	Panicf(tmp string, args ...interface{})
	WithFields(args ...interface{}) Logger
}

func SetLogger(newLogger Logger) {
	Log = newLogger
}
