package config

import (
	"flag"
	"reflect"
	"testing"

	"github.com/linnv/logx"
)

func TestGetDefaultConfigPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefaultConfigPath(); got != tt.want {
				t.Errorf("GetDefaultConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initConfig(t *testing.T) {
	tests := []struct {
		name       string
		wantConfig *Configuration
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfig, err := initConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("initConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("initConfig() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Configuration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig(t *testing.T) {
	defer logx.Flush()
	defaultConfigPath = "./config.yaml"
	InitFlag()
	flag.Parse()
	InitConfig()
	if !addressGcables.Match("3号604房") {
		t.Fail()
	}
	addressGcables.TrimAddrs()
}
