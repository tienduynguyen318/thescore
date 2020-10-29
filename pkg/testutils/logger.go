package testutils

import (
	"encoding/json"
	"io"
	"testing"

	"thescore/pkg/logger"
)

// LogFields to use with ParseAndVerifyLogMessage(...)
type LogField struct {
	Key   string
	Value string
}

// ParseAndVerifyLogMessage() takes in a []byte raw datareader stream and verifies that
// the log message is in the structured JSON datareader format and that the fields match.
// If there is an error parsing the json message, testing.T.Fatal is called.
func ParseAndVerifyLogMessage(t *testing.T, b []byte, fields []LogField) {
	logMsg := make(map[string]string)
	if err := json.Unmarshal(b, &logMsg); err != nil {
		t.Fatalf("Failed to unmarshal log message: %s", err)
	}
	for _, f := range fields {
		if logMsg[f.Key] != f.Value {
			t.Errorf(
				"log[%s] != \"%s\". Got: %s",
				f.Key,
				f.Value,
				logMsg[f.Key],
			)
		}
	}
}

// Sets up a test logger given an io.Writer. If creating the Logger fails, the
// test case causes a Fatal message to be sent.
func SetupTestLogger(t *testing.T, w io.Writer) logger.Logger {
	loggerConfig := logger.Config{
		Writer:     w,
		AppName:    "TestApp",
		AppVersion: "v1.0",
		Hostname:   "localhost",
	}
	log, err := logger.NewZapLogger(loggerConfig)
	if err != nil {
		t.Fatal(err)
	}
	return log
}
