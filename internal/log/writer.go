package log

import (
	"github.com/sirupsen/logrus"
)

type Writer struct {
	Logger *logrus.Logger
	Level  logrus.Level
}

// Write implements io.Writer interface and writes it as a log, with logrus
func (w *Writer) Write(p []byte) (n int, err error) {
	entry := logrus.NewEntry(w.Logger)
	entry.Log(w.Level, string(p))

	return len(p), nil
}
