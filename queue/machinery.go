package queue

import (
	"context"
	"fmt"
	"web-server/config"

	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"

	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	machineryConfig "github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
)

type machineryQueue struct {
	server *machinery.Server
}

type machineryWorker struct {
	worker *machinery.Worker
}

func NewJobQueue() JobQueue {
	// Create server instance
	redisConfig := config.Get().RedisConfig
	redisPath := fmt.Sprintf("redis://%s,%s,@%s:%d", redisConfig.Name, redisConfig.Password, redisConfig.Host, redisConfig.Port)
	conf := &machineryConfig.Config{
		ResultBackend:   redisPath,
		Broker:          redisPath,
		DefaultQueue:    "default_job",
		ResultsExpireIn: 3600,
		Redis: &machineryConfig.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	broker := redisbroker.NewGR(conf, []string{"localhost:6379"}, 0)
	backend := redisbackend.NewGR(conf, []string{"localhost:6379"}, 0)
	lock := eagerlock.New()

	server := machinery.NewServer(conf, broker, backend, lock)

	return &machineryQueue{
		server: server,
	}
}

func (queue *machineryQueue) Register(job Job) error {
	return queue.server.RegisterTask(job.Name, job.Run)
}

func (queue *machineryQueue) Publish(ctx context.Context, job Job) error {
	args := make([]tasks.Arg, len(job.Args))
	for i := range job.Args {
		args[i] = job.Args[i].ToTaskArg()
	}
	sig := &tasks.Signature{
		Name:       job.Name,
		Args:       args,
		RetryCount: job.Retries,
	}
	_, err := queue.server.SendTaskWithContext(ctx, sig)
	return err
}

func (queue *machineryQueue) NewWorker(name string) Worker {
	const concurrency = 10
	worker := queue.server.NewWorker(name, concurrency)
	return &machineryWorker{
		worker: worker,
	}
}

func (worker *machineryWorker) Run() error {
	return worker.worker.Launch()
}

func (worker *machineryWorker) Stop() {
	worker.worker.Quit()
}
