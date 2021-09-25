package queue

import "errors"

type JobHandler interface {
	HandleSendEmailJob(email string, content string) error
}

type jobHanderImpl struct {
}

func (handler *jobHanderImpl) HandleSendEmailJob(email string, content string) error {
	return errors.New("can not send email")
}

func NewJobHandler() JobHandler {
	return &jobHanderImpl{}
}
