package support

import (
	"reflect"
	"testing"

	"github.com/appist/appy/test"
)

type loggerSuite struct {
	test.Suite
}

func (s *loggerSuite) TestNewLogger() {
	logger := NewLogger()
	_, ok := reflect.TypeOf(logger).MethodByName("Desugar")

	s.Equal(true, ok)
}

func (s *loggerSuite) TestNewFakeLoggerInDebugBuild() {
	logger, buf, writer := NewTestLogger()
	logger.Info("test")
	writer.Flush()

	s.NotNil(logger)
	s.Contains(buf.String(), "\x1b[34mINFO\x1b[0m\ttest")
}

func (s *loggerSuite) TestNewFakeLoggerInReleaseBuild() {
	Build = ReleaseBuild
	defer func() { Build = DebugBuild }()

	logger, buf, writer := NewTestLogger()
	logger.Info("test")
	writer.Flush()

	s.NotNil(logger)
	s.Contains(buf.String(), "info\ttest")
}

func TestLoggerSuite(t *testing.T) {
	test.Run(t, new(loggerSuite))
}
