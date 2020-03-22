package configpi

import (
	"reflect"
	"testing"
)

func Test_factory_Configuration(t *testing.T) {
	type args struct {
		nodeType string
	}
	tests := []struct {
		name string
		args args
		want Configuration
	}{
		{
			name: "Server configuration should return when nodeType is server",
			args: args{
				nodeType: "server",
			},
			want: &Server{},
		},
		{
			name: "Agent configuration should return when nodeType is agent",
			args: args{
				nodeType: "agent",
			},
			want: &Agent{},
		},
		{
			name: "No configuration should return when nodeType is invalid",
			args: args{
				nodeType: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFactory()
			if got := f.Configuration(tt.args.nodeType); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("Configuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
