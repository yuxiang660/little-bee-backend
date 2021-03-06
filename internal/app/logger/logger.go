// Wrapper of logrus. Do not use logrus directly in the project.

package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

// SetLevel sets the logger level.
// 1:fatal,2:error,3:warn,4:info,5:debug,6:trace.
func SetLevel(level int) {
	logrus.SetLevel(logrus.Level(level))
}

// SetFormatter sets the format of logger message.
// Supported format: JSON or Text.
func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		logrus.SetFormatter(new(logrus.TextFormatter))
	}
}

// SetOutput sets the output of the logger. Supported output: stdout, stderr or file.
// If the output is file, returns a function to release the resource.
func SetOutput(out string, filename string) (func(), error) {
	if out == "" {
		return nil, nil
	}

	switch out {
	case "stdout":
		logrus.SetOutput(os.Stdout)
	case "stderr":
		logrus.SetOutput(os.Stderr)
	case "file":
		if name := filename; name != "" {
			_ = os.MkdirAll(filepath.Dir(name), 0777)
			f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				return nil, err
			}
			logrus.SetOutput(f)

			return func() {
				if f != nil {
					f.Close()
				}
			}, nil
		}
	}

	return nil, nil
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	entry := logrus.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Debug(args...)
}

// DebugWithFields logs a message with fields at level Debug on the standard logger.
func DebugWithFields(l interface{}, f Fields) {
	entry := logrus.WithFields(logrus.Fields(f))
	entry.Data["file"] = fileInfo(2)
	entry.Debug(l)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	entry := logrus.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Info(args...)
}

// InfoWithFields logs a message with fields at level Info on the standard logger.
func InfoWithFields(l interface{}, f Fields) {
	entry := logrus.WithFields(logrus.Fields(f))
	entry.Data["file"] = fileInfo(2)
	entry.Info(l)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	entry := logrus.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Warn(args...)
}

// WarnWithFields logs a message with fields at level Warn on the standard logger.
func WarnWithFields(l interface{}, f Fields) {
	entry := logrus.WithFields(logrus.Fields(f))
	entry.Data["file"] = fileInfo(2)
	entry.Warn(l)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	entry := logrus.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Error(args...)
}

// ErrorWithFields logs a message with fields at level Error on the standard logger.
func ErrorWithFields(l interface{}, f Fields) {
	entry := logrus.WithFields(logrus.Fields(f))
	entry.Data["file"] = fileInfo(2)
	entry.Error(l)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	entry := logrus.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Fatal(args...)
}

// FatalWithFields logs a message with fields at level Fatal on the standard logger.
func FatalWithFields(l interface{}, f Fields) {
	entry := logrus.WithFields(logrus.Fields(f))
	entry.Data["file"] = fileInfo(2)
	entry.Fatal(l)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
