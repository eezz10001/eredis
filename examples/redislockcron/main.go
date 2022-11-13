package main

import (
	"context"
	"log"

	"github.com/eezz10001/ego"
	"github.com/eezz10001/ego/core/elog"
	"github.com/eezz10001/ego/task/ecron"

	"github.com/eezz10001/eredis"
	"github.com/eezz10001/eredis/ecronlock"
)

var (
	redis *eredis.Component
)

// export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	err := ego.New().Invoker(initRedis).Cron(cronJob()).Run()
	if err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

func initRedis() error {
	redis = eredis.Load("redis.test").Build()
	return nil
}

func cronJob() ecron.Ecron {
	locker := ecronlock.DefaultContainer().Build(ecronlock.WithClient(redis))
	cron := ecron.Load("cron.default").Build(
		ecron.WithLock(locker.NewLock("ego-component:cronjob:syncXxx")),
		ecron.WithJob(helloWorld),
	)
	return cron
}

func helloWorld(ctx context.Context) error {
	log.Println("cron job running")
	return nil
}
