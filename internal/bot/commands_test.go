package bot

import (
	"reflect"
	"testing"
)

func Test_parseCommand(t *testing.T) {
	type args struct {
		raw    string
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		want    *parsedCommand
		wantErr bool
	}{
		{"Simple 1", args{"-play", "-"}, &parsedCommand{Command: "play", Args: ""}, false},
		{"Simple 2", args{"^play", "^"}, &parsedCommand{Command: "play", Args: ""}, false},
		{"With args 1", args{"^play test", "^"}, &parsedCommand{Command: "play", Args: "test"}, false},
		{"With args 2", args{"^play args with space", "^"}, &parsedCommand{Command: "play", Args: "args with space"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCommand(tt.args.raw, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
