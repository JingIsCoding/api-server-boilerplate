package queue

import (
	"web-server/model"

	"github.com/RichardKnop/machinery/v2/tasks"
)

type Job struct {
	Name string
	Args []Arg
	// should be a function
	Run       interface{}
	Retries   int
	OnSuccess *Job
	OnError   *Job
}

type Arg struct {
	Name  string
	Type  string
	Value interface{}
}

func (arg Arg) ToTaskArg() tasks.Arg {
	return tasks.Arg{
		Name:  arg.Name,
		Type:  arg.Type,
		Value: arg.Value,
	}
}

func NewSendEmailJob(user model.User, content string) Job {
	return Job{
		Name: "send_email",
		Args: []Arg{
			Arg{
				Name:  "email",
				Type:  "string",
				Value: user.Email,
			},
			Arg{
				Name:  "content",
				Type:  "string",
				Value: content,
			},
		},
		Retries: 3,
	}
}
