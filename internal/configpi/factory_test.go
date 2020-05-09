package configpi

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_factory_Configuration(t *testing.T) {
	type args struct {
		nodeType string
		input    Input
		log      *logrus.Logger
	}
	tests := []struct {
		name string
		args args
		want Configuration
	}{
		{
			name: "server configuration should return when nodeType is server",
			args: args{
				nodeType: "server",
				input:    Input{},
				log:      logrus.New(),
			},
			want: &server{},
		},
		{
			name: "Agent configuration should return when nodeType is agent",
			args: args{
				nodeType: "agent",
				input:    Input{},

				log: logrus.New(),
			},
			want: &Agent{},
		},
		{
			name: "No configuration should return when nodeType is invalid",
			args: args{
				nodeType: "",
				input:    Input{},
				log:      logrus.New(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFactory()
			if got := f.Configuration(tt.args.nodeType, tt.args.input, tt.args.log); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("Configuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
