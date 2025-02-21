package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	JWT     JWTConfig
	MyRedis MyRedis
}

type JWTConfig struct {
	Secret string
}

type MyRedis struct {
	Addr string
}
