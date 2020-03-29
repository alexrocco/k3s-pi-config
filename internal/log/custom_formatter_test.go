package log

import (
	"bytes"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestCustomFormatter_Format(t *testing.T) {
	type arg struct {
		message  string
		command  string
		logLevel logrus.Level
		logTime  string
	}

	type test struct {
		name      string
		args      arg
		wantedLog string
	}

	tests := []test{
		{
			name: "Test TRACE",
			args: arg{
				message:  "test message",
				command:  "test",
				logLevel: logrus.TraceLevel,
				logTime:  "2020-03-15T12:59:01",
			},
			wantedLog: "20-03-15 12:59:01 \x1b[37mTRACE\x1b[0m test test message\n",
		},
		{
			name: "Test DEBUG",
			args: arg{
				message:  "test message",
				command:  "test",
				logLevel: logrus.DebugLevel,
				logTime:  "2020-03-15T12:59:01",
			},
			wantedLog: "20-03-15 12:59:01 \x1b[37mDEBUG\x1b[0m test test message\n",
		},
		{
			name: "Test INFO",
			args: arg{
				message:  "test message",
				command:  "test",
				logLevel: logrus.InfoLevel,
				logTime:  "2020-03-15T12:59:01",
			},
			wantedLog: "20-03-15 12:59:01 \x1b[36m INFO\x1b[0m test test message\n",
		},
		{
			name: "Test WARN",
			args: arg{
				message:  "test message",
				command:  "test",
				logLevel: logrus.WarnLevel,
				logTime:  "2020-03-15T12:59:01",
			},
			wantedLog: "20-03-15 12:59:01 \x1b[33m WARN\x1b[0m test test message\n",
		},
		{
			name: "Test ERROR",
			args: arg{
				message:  "test message",
				command:  "test",
				logLevel: logrus.ErrorLevel,
				logTime:  "2020-03-15T12:59:01",
			},
			wantedLog: "20-03-15 12:59:01 \x1b[31mERROR\x1b[0m test test message\n",
		},
		{
			name: "Test FATAL",
			args: arg{
				message:  "test message",
				command:  "test",
				logLevel: logrus.FatalLevel,
				logTime:  "2020-03-15T12:59:01",
			},
			wantedLog: "20-03-15 12:59:01 \x1b[31mFATAL\x1b[0m test test message\n",
		},
		{
			name: "Test PANIC",
			args: arg{
				message:  "test message",
				command:  "test",
				logLevel: logrus.PanicLevel,
				logTime:  "2020-03-15T12:59:01",
			},
			wantedLog: "20-03-15 12:59:01 \x1b[31mPANIC\x1b[0m test test message\n",
		},

		{
			name: "Test invalid log Level",
			args: arg{
				message:  "test message",
				command:  "test",
				logLevel: 99,
				logTime:  "2020-03-15T12:59:01",
			},
			wantedLog: "20-03-15 12:59:01 \x1b[37m NONE\x1b[0m test test message\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testLog := logrus.New()

			testLog.Level = tt.args.logLevel
			testLog.Formatter = &CustomFormatter{Command: tt.args.command}

			// output to a buffer to compare logs
			var out bytes.Buffer
			testLog.Out = &out

			logTime, err := time.Parse("2006-01-02T15:04:05", tt.args.logTime)
			if err != nil {
				t.Error(err)
				return
			}

			validLogMsg := func() {
				if out.String() != tt.wantedLog {
					t.Errorf("want %s, got %s", tt.wantedLog, out.String())
				}
			}

			defer func() {
				if r := recover(); r != nil {
					if logrus.PanicLevel == tt.args.logLevel {
						validLogMsg()
					}
				}
			}()

			// Set the time and Level to the log
			testLog.WithTime(logTime).Log(tt.args.logLevel, tt.args.message)

			validLogMsg()
		})
	}
}
