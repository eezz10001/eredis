package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/eezz10001/eredis"
	"github.com/eezz10001/ego"
	"github.com/eezz10001/ego/core/elog"
	"github.com/eezz10001/ego/server/egovernor"
)

// export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	err := ego.New().Invoker(
		invokerRedis,
		testRedis,
	).Serve(
		egovernor.Load("server.governor").Build(),
	).Run()
	if err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

var eredisClient *eredis.Component

func invokerRedis() error {
	eredisClient = eredis.Load("redis.test").Build()
	return nil
}

func testRedis() error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Ego-Uid", 9527)
	err := eredisClient.Set(ctx, "hello", "world", 0)
	fmt.Println("set hello", err)

	str, err := eredisClient.Get(ctx, "hello")
	fmt.Println("get hello", str, err)

	str, err = eredisClient.Get(ctx, "lee")
	fmt.Println("Get lee", errors.Is(err, eredis.Nil), "err="+err.Error())

	return nil
}
