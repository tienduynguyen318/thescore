package logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"testing"
)

func TestSetLogger(t *testing.T) {
	var buf bytes.Buffer

	log, err := NewZapLogger(Config{
		Writer:     &buf,
		AppName:    "TestApp",
		AppVersion: "v1.0",
		Hostname:   "localhost",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	SetLogger(log)
	Log.Infof("global logger info message")

	x := make(map[string]string)
	if err := json.Unmarshal(buf.Bytes(), &x); err != nil {
		t.Fatalf("Failed to unmarshal log json: %s", err)
	}
	if x["log.level"] != "info" {
		t.Errorf(`log.level != "info". Got: "%s"`, x["log.level"])
	}
	if x["message"] != "global logger info message" {
		t.Errorf(`message != "global logger info message". Got: "%s"`, x["message"])
	}
	if match, err := regexp.MatchString(rfc3339Pattern, x["@timestamp"]); err != nil || !match {
		t.Errorf("Timestamp does not follow the RFC3339 patterns.\nGot: %s", x["@timestamp"])
	}
	if match, err := regexp.MatchString(rfc3339Pattern, x["@timestamp"]); err != nil || !match {
		t.Errorf("Timestamp does not follow the RFC3339 patterns.\nGot: %s", x["@timestamp"])
	}

}

func withWriter(c Config, w io.Writer) Config {
	config := c
	config.Writer = w
	return config
}

func withAppName(c Config, name string) Config {
	config := c
	config.AppName = name
	return config
}

func withAppVersion(orig Config, version string) Config {
	c := orig
	c.AppVersion = version
	return c
}

func withHostname(orig Config, hostname string) Config {
	c := orig
	c.Hostname = hostname
	return c
}

func TestLoggerConfig(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer

	config := Config{
		Writer:     &buf,
		AppName:    "MyApp",
		AppVersion: "v0.1.0",
		Hostname:   "localhost",
	}
	testcases := []struct {
		Name          string
		Config        Config
		Valid         bool
		ExpectedError error
	}{
		{"Valid config", config, true, nil},
		{"Without a Writer", withWriter(config, nil), false, errors.New("Writer is required")},
		{"Invalid without an AppName", withAppName(config, ""), false, errors.New("AppName is required")},
		{"Without an AppVersion", withAppVersion(config, ""), false, errors.New("AppVersion is required")},
		{"Without a Hostnam", withHostname(config, ""), false, errors.New("Hostname is required")},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			v, err := tc.Config.valid()
			if v != tc.Valid {
				if tc.Valid {
					t.Logf("%v | %v", v, err)
					t.Errorf("Expected a valid configuration")
				} else {
					t.Errorf("Expected an invalid configuration")
				}
			}
			if tc.ExpectedError == nil {
				return
			}
			if err == nil || err.Error() != tc.ExpectedError.Error() {
				t.Errorf(
					"Wrong error was returned.\n\nExpected:\n%+v\n\nGot:\n%+v\n",
					tc.ExpectedError,
					err,
				)

			}

		})
	}
}
