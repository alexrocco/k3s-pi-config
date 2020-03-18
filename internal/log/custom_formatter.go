package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	timeFormat = "06-01-02 15:04:05"

	// terminal colors
	red       = 31
	yellow    = 33
	cyan      = 36
	lightGray = 37
)

// CustomFormatter creates a custom formatter for logrus inheriting the default behavior logrus.TextFormatter
type CustomFormatter struct {
	Command string
	logrus.TextFormatter
}

func (cf *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Define log level colors and level text
	var levelColor int
	var levelText string
	switch entry.Level {
	case logrus.TraceLevel:
		levelText = "TRACE"
		levelColor = lightGray
	case logrus.DebugLevel:
		levelText = "DEBUG"
		levelColor = lightGray
	case logrus.InfoLevel:
		levelText = "INFO"
		levelColor = cyan
	case logrus.WarnLevel:
		levelText = "WARN"
		levelColor = yellow
	case logrus.ErrorLevel:
		levelText = "ERROR"
		levelColor = red
	case logrus.FatalLevel:
		levelText = "FATAL"
		levelColor = red
	case logrus.PanicLevel:
		levelText = "PANIC"
		levelColor = red
	default:
		levelText = "NONE"
		levelColor = lightGray
	}

	// Format the message
	msg := fmt.Sprintf("%s \x1b[%dm%5s\x1b[0m %s %s\n",
		entry.Time.Format(timeFormat),
		levelColor,
		levelText,
		cf.Command,
		entry.Message,
	)

	return []byte(msg), nil
}
