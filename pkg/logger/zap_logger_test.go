package logger

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"
)

const rfc3339Pattern = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(Z|[-+]\d{2}:\d{2})$`

var expectations = []struct {
	k string
	v string
}{
	{"ecs.version", "1.5.0"},
	{"service.name", "TestApp"},
	{"service.type", "TestApp"},
	{"service.version", "v1.0"},
	{"host.hostname", "localhost"},
}

func TestLogger(t *testing.T) {
	var logger Logger
	var buf bytes.Buffer

	config := Config{
		Writer:     &buf,
		AppName:    "TestApp",
		AppVersion: "v1.0",
		Hostname:   "localhost",
	}
	logger, err := NewZapLogger(config)
	if err != nil {
		t.Fatalf("Did not expect an error initializing the logger. Got: %s\n", err)
	}

	runs := []struct {
		name  string
		level string
		f     func(string)
	}{
		{
			name:  "Info Level",
			level: "info",
			f: func(n string) {
				logger.Infof(n)
			},
		},
		{
			name:  "Warn Level",
			level: "warn",
			f:     func(n string) { logger.Warnf(n) },
		},
		// {
		// 	name:  "Debug Level",
		// 	level: "debug",
		// 	f:     func(n string) { logger.Debugf(n) },
		// },
		{
			name:  "Error Level",
			level: "error",
			f:     func(n string) { logger.Errorf(n) },
		},
		{
			name:  "Panic Level",
			level: "panic",
			f: func(n string) {
				defer func() {
					e := recover()
					if e != n {
						t.Errorf("Expected: %s\nGot: %s", n, e)
					}
				}()
				logger.Panicf(n)
			},
		},
	}

	for _, run := range runs {
		t.Run(run.name, func(t *testing.T) {
			buf.Reset()
			run.f(run.name)
			x := make(map[string]string)
			if err := json.Unmarshal(buf.Bytes(), &x); err != nil {
				t.Fatalf("Failed to unmarshal log json: %s", err)
			}

			if x["log.level"] != run.level {
				t.Errorf(`log.level != "%s". Got: "%s"`, run.level, x["log.level"])
			}
			if x["message"] != run.name {
				t.Errorf(`message != "%s". Got: "%s"`, run.name, x["message"])
			}
			for _, e := range expectations {
				if x[e.k] != e.v {
					t.Errorf(`%s != "%s".\nGot: "%s"`, e.k, e.v, x[e.k])
				}
			}
			if match, err := regexp.MatchString(rfc3339Pattern, x["@timestamp"]); err != nil || !match {
				t.Errorf("Timestamp does not follow the RFC3339 patterns.\nGot: %s", x["@timestamp"])
			}

		})
	}
}

func TestLoggerFatal(t *testing.T) {
	fpath := filepath.Join(os.TempDir(), t.Name())

	if os.Getenv("RUN_FATAL") == "1" {
		var logger Logger
		f, err := os.Create(fpath)
		if err != nil {
			t.Fatalf("Failed to create assertion file: %s", err)
			return
		}
		logger, err = NewZapLogger(Config{
			Writer:     f,
			AppName:    "TestApp",
			AppVersion: "v1.0",
			Hostname:   "localhost",
		})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		logger.Fatalf("Fatal Message")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestLoggerFatal")
	cmd.Env = append(os.Environ(), "RUN_FATAL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); !ok || e.Success() {
		t.Fatalf("Logger.Fatalf() failed to exit")
	}
	f, err := os.Open(fpath)
	if err != nil {
		t.Fatalf("Failed to open assertion file: %s", err)
	}
	x := make(map[string]string)
	if err := json.NewDecoder(f).Decode(&x); err != nil {
		t.Fatalf("Failed to unmarshal log message: %s", err)
	}

	if x["log.level"] != "fatal" {
		t.Errorf(`log.level != "fatal". Got: "%s"`, x["log.level"])
	}
	if x["message"] != "Fatal Message" {
		t.Errorf(`message != "Fatal Message". Got: "%s"`, x["message"])
	}
	for _, e := range expectations {
		if x[e.k] != e.v {
			t.Errorf(`%s != "%s".\nGot: "%s"`, e.k, e.v, x[e.k])
		}
	}
	if match, err := regexp.MatchString(rfc3339Pattern, x["@timestamp"]); err != nil || !match {
		t.Errorf("Timestamp does not follow the RFC3339 patterns.\nGot: %s", x["@timestamp"])
	}

}

func TestLoggerWithFields(t *testing.T) {
	var logger Logger
	var buf bytes.Buffer

	logger, err := NewZapLogger(Config{
		Writer:     &buf,
		AppName:    "TestApp",
		AppVersion: "v1.0",
		Hostname:   "localhost",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	logger.WithFields("user", "bob").Infof("info message")

	x := make(map[string]string)
	if err := json.Unmarshal(buf.Bytes(), &x); err != nil {
		t.Fatalf("Failed to unmarshal log message: %s", err)
	}

	for _, e := range expectations {
		if x[e.k] != e.v {
			t.Errorf(`%s != "%s".\nGot: "%s"`, e.k, e.v, x[e.k])
		}
	}
	if x["log.level"] != "info" {
		t.Errorf(`log.level != "info". Got: "%s"`, x["log.level"])
	}
	if x["message"] != "info message" {
		t.Errorf(`message != "info message". Got: "%s"`, x["message"])
	}
	if match, err := regexp.MatchString(rfc3339Pattern, x["@timestamp"]); err != nil || !match {
		t.Errorf("Timestamp does not follow the RFC3339 patterns.\nGot: %s", x["@timestamp"])
	}
	if match, err := regexp.MatchString(rfc3339Pattern, x["@timestamp"]); err != nil || !match {
		t.Errorf("Timestamp does not follow the RFC3339 patterns.\nGot: %s", x["@timestamp"])
	}
	if x["user"] != "bob" {
		t.Errorf("Settings field \"user\" failed.\nExpected:\nbob\n\nGot:\n%s", x["user"])
	}
}
