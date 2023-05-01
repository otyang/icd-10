package config

import "testing"

func TestLoad(t *testing.T) {
	type args struct {
		configFile string
		config     any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Load(tt.args.configFile, tt.args.config)
		})
	}
}
