package api

import (
	"flag"
	"reflect"
	"testing"

	conf "github.com/linnv/templatego/config"
)

func Test_hello(t *testing.T) {
	conf.InitFlag()
	flag.Parse()
	*conf.ConfigFile = "/home/jialin/go/src/github.com/linnv/templatego/config/config.yaml"
	conf.InitConfig()

	Init()
	type args struct {
		ctx *Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantBs  []byte
		wantErr bool
	}{
		{"normal", args{nil}, []byte("hello" + config.AppName), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBs, err := hello(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBs, tt.wantBs) {
				t.Errorf("hello() = %v, want %v", gotBs, tt.wantBs)
			}
		})
	}
}
