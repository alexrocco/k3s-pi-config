package log

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func Test_Writer_Write(t *testing.T) {
	type fields struct {
		level logrus.Level
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantN      int
		wantErr    bool
		wantOutput string
	}{
		{
			name: "It should log as a info message",
			fields: fields{
				level: logrus.InfoLevel,
			},
			args: args{
				p: []byte("test info log"),
			},
			wantN:      13,
			wantErr:    false,
			wantOutput: "{\"Level\":\"info\",\"message\":\"test info log\",\"time\":\"%d\"}\n",
		},
		{
			name: "It should log as a error message",
			fields: fields{
				level: logrus.ErrorLevel,
			},
			args: args{
				p: []byte("test info error"),
			},
			wantN:      15,
			wantErr:    false,
			wantOutput: "{\"Level\":\"error\",\"message\":\"test info error\",\"time\":\"%d\"}\n",
		},
		{
			name: "It should log with unknown Level",
			fields: fields{
				level: 1000,
			},
			args: args{
				p: []byte("test info unknown"),
			},
			wantN:      17,
			wantErr:    false,
			wantOutput: "{\"Level\":\"unknown\",\"message\":\"test info unknown\",\"time\":\"%d\"}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := logrus.New()

			var output bytes.Buffer
			logger.Out = &output

			logger.Level = tt.fields.level

			logger.Formatter = &logrus.JSONFormatter{
				// To simplify the asserts, format data with the year only
				TimestampFormat: "2006",
				FieldMap: logrus.FieldMap{
					logrus.FieldKeyLevel: "Level",
					logrus.FieldKeyMsg:   "message",
					logrus.FieldKeyTime:  "time",
				},
				CallerPrettyfier: nil,
				PrettyPrint:      false,
			}

			k := &Writer{
				Logger: logger,
				Level:  tt.fields.level,
			}

			gotN, err := k.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotN != tt.wantN {
				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
			}

			if output.String() != fmt.Sprintf(tt.wantOutput, time.Now().Year()) {
				t.Errorf("Write() output = %v, want %v", output.String(), tt.wantOutput)
			}
		})
	}
}
