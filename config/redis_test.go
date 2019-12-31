package config

import (
	"fmt"
	"testing"

	"github.com/linnv/logx"
)

func TestNewRedisGo(t *testing.T) {
	return
	p := NewRedisGo("192.168.1.125:6379", "")

	keys, err := p.Keys("*")
	if err != nil {
		logx.Warnf("err: %+v\n", err)
		return
	}
	for k, v := range keys {
		//@toDelete
		fmt.Printf("%+v: %+v\n", k, v)
	}
	oneKey := "SESSION_KEY_OUTCALL_sess 1560944366944990710-4"
	bs, err := p.Get(oneKey)
	if err != nil {
		logx.Warnf("err: %+v\n", err)
		return
	}
	logx.Warnf("bs: %s\n", bs)

	err = p.Expire(oneKey, 10)
	logx.Warnf("err: %+v\n", err)

	// c := p.Get()
	// defer c.Close()
	// rep, err := c.Do("keys", "*")
	// if err != nil {
	// 	logx.Warnf("err: %+v\n", err)
	// }
	// // var list []string
	// bs, err := redis.ByteSlices(rep, err)
	// if err != nil {
	// 	logx.Warnf("err: %+v\n", err)
	// 	return
	// }
	// // err = json.Unmarshal(bs, &list)
	// logx.Warnf("rep: %s\n", rep)
	// logx.Warnf("err: %+v\n", err)
	// // logx.Warnf("list: %+v\n", list)
	// for k, v := range bs {
	// 	fmt.Printf("%+v: %s\n", k, v)
	// }
}
