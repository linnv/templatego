package config

import (
	"fmt"
	"strings"

	"github.com/linnv/logx"
)

type AddressGcable struct {
	ID           string `json:"ID"`
	Address      string `json:"Address"`
	AddressShort string `json:"AddressShort"`
	IsNormal     string `json:"IsNormal"`
	IsNormalBool bool   `json:"IsNormalBool"`
}

type AddressGcables struct {
	Addrs []AddressGcable `json:"Addrs"`
}

func (asg *AddressGcables) Match(keyword string) bool {
	for _, v := range asg.Addrs {
		if strings.Contains(v.Address, keyword) {
			return true
		}
	}
	return false
}

func (asg *AddressGcables) MatchShort(keyword string) bool {
	for _, v := range asg.Addrs {
		if strings.Contains(v.Address, keyword) {
			return true
		}
	}
	return false
}

func (asg *AddressGcables) String() {
	for k, v := range asg.Addrs {
		fmt.Printf("%+v: %+v\n", k, v)
	}
}

func (asg *AddressGcables) TrimAddrs() {
	trimPrefix := "市"
	y := "是"
	// n := "否"
	for k, v := range asg.Addrs {
		trimIndex := strings.Index(v.Address, trimPrefix)
		if trimIndex > 0 {
			logx.Debugf("%+v: %+v trim:%s\n", k, v, v.Address[trimIndex+1:])
			v.AddressShort = v.Address[trimIndex+1:]
		} else {
			logx.Debugf("%+v: %+v\n", k, v)
		}
		if v.IsNormal == y {
			v.IsNormalBool = true
		} else {
			v.IsNormalBool = false
		}
	}
}
