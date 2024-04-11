package main

import "github.com/zeromicro/go-zero/core/stores/redis"

func main() {

	conf := redis.RedisConf{
		Host: "127.0.0.1:6379,127.0.0.2:6379,127.0.0.3:6379,127.0.0.4:6379",
		Type: redis.ClusterType,
	}

	conn := redis.MustNewRedis(conf)

	conn.Set("", "")

}
