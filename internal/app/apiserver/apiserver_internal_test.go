package apiserver

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func ConfigureTestLogger(t *testing.T, s *APIServer, writer io.Writer) {
	t.Helper()
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		t.Fatal(err)
	}
	s.logger.SetLevel(level)
	s.logger.SetOutput(writer)
}

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	buffer := bytes.Buffer{}
	ConfigureTestLogger(t, s, &buffer)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	t.Run("check response body", func(t *testing.T) {
		assert.Equal(t, rec.Body.String(), "Hello, world!")
	})

	t.Run("check logger msg is not empty", func(t *testing.T) {
		assert.NotEmpty(t, buffer.String())
	})

}
