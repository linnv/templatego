package util

import (
	"testing"
)

func TestVerifyContact(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"18290015121"}, true},
		{"normal", args{"17290015121"}, true},
	}
	for _, tt := range tests {
		if got := VerifyContact(tt.args.number); got != tt.want {
			t.Errorf("VerifyContact() = %v, want %v", got, tt.want)
		}
	}
}

func TestVerifyContactFilterFirst(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name            string
		args            args
		wantFilterQuery string
		wantOk          bool
	}{
		{"normal", args{"一零一"}, "101", false},
		{"normal", args{"一圾一"}, "11", false},
		{"normal", args{"零叁佰贰拾玖亿壹仟柒佰伍拾陆万肆仟捌佰叁拾陆圆"}, "032917564836", false},
		{"normal", args{"贰佰叁拾贰万壹仟叁佰壹拾贰圆"}, "2321312", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFilterQuery, _, gotOk := VerifyContactFilterFirst(tt.args.number)
			if gotFilterQuery != tt.wantFilterQuery {
				t.Errorf("VerifyContactFilterFirst() gotFilterQuery = %v, want %v", gotFilterQuery, tt.wantFilterQuery)
			}
			if gotOk != tt.wantOk {
				t.Errorf("VerifyContactFilterFirst() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestNumber2Cn(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		// {"normal", args{"1234"}, "eee"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := Number2Cn(tt.args.number); gotRet != tt.wantRet {
				t.Errorf("Number2Cn() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestVerifyCardID(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name            string
		args            args
		wantFilterQuery string
		wantOk          bool
	}{
		{"normal", args{"闽D12345"}, "闽D12345", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFilterQuery, gotOk := VerifyCardID(tt.args.number)
			if gotFilterQuery != tt.wantFilterQuery {
				t.Errorf("VerifyCardID() gotFilterQuery = %v, want %v", gotFilterQuery, tt.wantFilterQuery)
			}
			if gotOk != tt.wantOk {
				t.Errorf("VerifyCardID() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
