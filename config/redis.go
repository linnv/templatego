package config

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/linnv/logx"
)

type RedisGo struct {
	Pool *redis.Pool
}

func (r *RedisGo) Get(key string) (bs []byte, err error) {
	if r == nil || key == "" {
		return
	}
	c := r.Pool.Get()
	defer c.Close()
	rep, err := c.Do("Get", key)
	if err != nil {
		logx.Errorf("err: %+v\n", err)
		return
	}
	bs, err = redis.Bytes(rep, err)
	if err != nil {
		logx.Errorf("err: %+v\n", err)
	}
	return
}

func (r *RedisGo) Expire(key string, second int) (err error) {
	if r == nil || key == "" {
		return
	}
	c := r.Pool.Get()
	defer c.Close()
	var args = []interface{}{key, second}
	rep, err := c.Do("EXPIRE", args...)
	if err != nil {
		logx.Errorf("err: %+v rep: %+v\n", err, rep)
		return
	}
	return
}

func (r *RedisGo) Keys(keyRegular string) (bs []string, err error) {
	if r == nil || keyRegular == "" {
		return
	}
	c := r.Pool.Get()
	defer c.Close()
	rep, err := c.Do("keys", keyRegular)
	if err != nil {
		logx.Errorf("err: %+v\n", err)
		return
	}

	bsbs, err := redis.ByteSlices(rep, err)
	if err != nil {
		logx.Errorf("err: %+v\n", err)
		return
	}
	bs = make([]string, len(bsbs))
	for i := 0; i < len(bsbs); i++ {
		bs[i] = string(bsbs[i])
	}
	return
}

var RedisPool *RedisGo

func NewRedisGo(addr, password string) *RedisGo {
	pool := &redis.Pool{
		MaxIdle:     200,
		IdleTimeout: 240 * time.Second,
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				logx.Errorf("err: %+v\n", err)
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					logx.Errorf("err: %+v\n", err)
					return nil, err
				}
			}
			//@TODO
			// if _, err := c.Do("SELECT", db); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			//@TODO
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			logx.Errorf("err: %+v\n", err)
			return err
		},
	}
	RedisPool = new(RedisGo)
	RedisPool.Pool = pool
	return RedisPool
}
