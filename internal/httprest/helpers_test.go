package httprest

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"thescore/pkg/testutils"
)

func runServerTest(t *testing.T, r *http.Request) *http.Response {
	w := httptest.NewRecorder()
	logger := testutils.SetupTestLogger(t, ioutil.Discard)
	var config Config
	config.Logger = logger
	config.Service = newServiceDouble()
	server, err := New(config)
	if err != nil {
		t.Fatalf("Failed setting up HTTP Server: %v", err)
	}
	server.ServeHTTP(w, r)
	return w.Result()
}

func TestMain(m *testing.M) {
	loadHandlerFixtures()
	os.Exit(m.Run())
}

var (
	handlerFixtures = make(map[string]io.ReadCloser)
)

func loadHandlerFixtures() {
	fixturePaths, err := filepath.Glob("testdata/*.json")
	if err != nil {
		panic(fmt.Sprintf("Cannot glob fixture files: %s", err))
	}
	for _, fpath := range fixturePaths {
		f, err := ioutil.ReadFile(fpath)
		if err != nil {
			panic(fmt.Sprintf("Cannot read fixture file: %s", err))
		}
		handlerFixtures[filepath.Base(fpath)] = ioutil.NopCloser(bytes.NewReader(f))
	}
}

func handlerFixture(fname string) io.ReadCloser {
	var err error
	out := new(bytes.Buffer)
	b1 := bytes.NewBuffer([]byte{})
	b2 := bytes.NewBuffer([]byte{})
	tr := io.TeeReader(handlerFixtures[fname], b1)

	defer func() { handlerFixtures[fname] = ioutil.NopCloser(b1) }()
	if _, err = io.Copy(b2, tr); err != nil {
		log.Println(err.Error())
	}
	if _, err = out.ReadFrom(b2); err != nil {
		log.Println(err.Error())
	}

	return ioutil.NopCloser(out)
}
