package main

import (
	"flag"
	"reflect"
	"testing"
)

func Test_collectCliArgs(t *testing.T) {
	type args struct {
		flags *flag.FlagSet
		args  []string
	}
	var flags flag.FlagSet

	tests := []struct {
		name             string
		args             args
		wantCliFlags     []string
		wantTemplateOpts map[string]interface{}
		wantErr          bool
	}{
		{
			name: "Collect servers cli arguments",
			args: args{
				flags: &flags,
				args:  []string{"-servers=http://localhost:8080;description,http://localhost:8081;more description"},
			},
			wantCliFlags: nil,
			wantTemplateOpts: map[string]interface{}{
				"servers": [][]string{
					{
						"http://localhost:8080",
						"description",
					},
					{
						"http://localhost:8081",
						"more description",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCliFlags, gotTemplateOpts, err := collectCliArgs(tt.args.flags, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("collectCliArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCliFlags, tt.wantCliFlags) {
				t.Errorf("collectCliArgs() gotCliFlags = %v, want %v", gotCliFlags, tt.wantCliFlags)
			}
			if !reflect.DeepEqual(gotTemplateOpts, tt.wantTemplateOpts) {
				t.Errorf("collectCliArgs() gotTemplateOpts = %v, want %v", gotTemplateOpts, tt.wantTemplateOpts)
			}
		})
	}
}
