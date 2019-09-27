package util

import (
	"testing"
)

func TestRemoveNewLine(t *testing.T) {
	// v := "201801"
	// v := "20181"
	// logx.Warnf("v: %+v\n", v)
	// cTime, err := time.Parse("200601", strings.TrimSpace(v))
	// if err != nil {
	// 	logx.Warnf("err: %+v\n", err)
	// 	time.Sleep(time.Second)
	// 	cTime, err = time.Parse("20061", strings.TrimSpace(v))
	// 	if err != nil {
	// 		logx.Warnf("err: %+v\n", err)
	// 		time.Sleep(time.Second)
	// 	}
	// 	return
	// }
	// logx.Debugf("cTime: %+v\n", cTime)
	// time.Sleep(time.Second)

	// logx.Warnf("strings.Trimspace(): %+v\n", "fjeif\n")
	// logx.Warnf("strings.Trimspace(): %+v\n", strings.TrimSpace("fjeif\r\n"))
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{""}, ""},
		{"normal", args{"\n"}, ""},
		{"normal", args{"\r\n"}, ""},
		{"normal", args{"ajfe"}, "ajfe"},
		{"normal", args{"ajfe\n"}, "ajfe"},
		{"normal", args{"ajfe\r\n"}, "ajfe"},
		{"normal", args{"a\r\nb"}, "ab"},
		{"normal", args{"a\r\nb\n"}, "ab"},
		{"normal", args{"a\r\n    b\n"}, "ab"},
		{"normal", args{"ab\r\n"}, "ab"},
		{"normal", args{"ab\n"}, "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveNewLine(tt.args.str); got != tt.want {
				t.Errorf("RemoveNewLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
