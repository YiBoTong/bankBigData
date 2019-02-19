package redis

import "gitee.com/johng/gf/g"

const prefix = "bbd_"
const cache = "cache"

func Set(key string, value []byte) error {
	_, e := g.Redis(cache).Do("Set", prefix+key, value)
	return e
}

func Clear() error {
	_, e := g.Redis(cache).Do("FLUSHDB")
	return e
}

func Get(key string) (interface{}, error) {
	res, err := g.Redis(cache).Do("GET", prefix+key)
	return res, err
}
