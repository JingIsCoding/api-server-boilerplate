package queue

import "context"

type Worker interface {
	Run() error
	Stop()
}

type JobQueue interface {
	Register(Job) error
	Publish(context.Context, Job) error
	NewWorker(name string) Worker
}
