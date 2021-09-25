package main

import (
	"os"
	"web-server/queue"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		os.Exit(1)
	}
	queue := queue.NewJobQueue()
	registerTasks(queue)
	worker := queue.NewWorker(hostname)
	if err := worker.Run(); err != nil {
		worker.Stop()
	}
}

func registerTasks(jobQueue queue.JobQueue) {
	handler := queue.NewJobHandler()
	sendEmailJob := queue.Job{
		Name: "send_email",
		Run:  handler.HandleSendEmailJob,
	}
	jobQueue.Register(sendEmailJob)
}
