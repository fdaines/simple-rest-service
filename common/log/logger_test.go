package log

import (
	"bytes"
	"go.uber.org/zap"
	"net/url"
	"os"
	"strings"
	"testing"
)

type MemorySink struct {
	*bytes.Buffer
}
func (s *MemorySink) Close() error { return nil }
func (s *MemorySink) Sync() error  { return nil }

var sink *MemorySink

// TestMain will exec each test, one by one
func TestMain(m *testing.M) {
	sink = &MemorySink{new(bytes.Buffer)}
	zap.RegisterSink("memory", func(*url.URL) (zap.Sink, error) {
		return sink, nil
	})
	conf := zap.NewProductionConfig() // TODO: replace with real config
	conf.OutputPaths = []string{"memory://"}
	logger, err := conf.Build()
	if err != nil {
		os.Exit(-1)
	}
	instance.log = logger
	retCode := m.Run()
	os.Exit(retCode)
}

func TestLogger(t *testing.T) {
	t.Run("should logs info message", func(t *testing.T) {
		instance.Info("foobar")
		output := sink.String()

		if !strings.Contains(output, "info") {
			t.Errorf("Expects info message")
		}
		if !strings.Contains(output, "foobar") {
			t.Errorf("Expects 'foobar' message")
		}
	})
	/*
	t.Run("should logs error message", func(t *testing.T) {
		instance.Error("foobar")
		output := sink.String()

		if !strings.Contains(output, "error") {
			t.Errorf("Expects info message")
		}
		if !strings.Contains(output, "foobar") {
			t.Errorf("Expects 'foobar' message")
		}
	})

	t.Run("should logs fatal message", func(t *testing.T) {
		instance.Fatal("foobar")
		output := sink.String()

		if !strings.Contains(output, "fatal") {
			t.Errorf("Expects info message")
		}
		if !strings.Contains(output, "foobar") {
			t.Errorf("Expects 'foobar' message")
		}
	})
	*/
}
