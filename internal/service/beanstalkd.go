package service

import (
	"fmt"
	"github.com/beanstalkd/go-beanstalk"
)

func PutQueueMessageToBeanstalkd(queueName string) error {
	conn, err := beanstalk.Dial("tcp", "beanstalkd:11300")
	if err != nil {
		return fmt.Errorf("failed to connect to Beanstalkd: %v", err)
	}
	defer conn.Close()

	_, err = conn.Put([]byte(String(1000)), 1, 0, 0)
	if err != nil {
		return fmt.Errorf("failed to put message to Beanstalkd: %v", err)
	}

	return nil
}

func ReadQueueMessageFromBeanstalkd(queueName string) (string, error) {
	conn, err := beanstalk.Dial("tcp", "beanstalkd:11300")
	if err != nil {
		return "", fmt.Errorf("failed to connect to Beanstalkd: %v", err)
	}
	defer conn.Close()

	job, body, err := conn.Reserve(0)
	if err != nil {
		return "", fmt.Errorf("failed to read message from Beanstalkd: %v", err)
	}

	message := string(body)
	err = conn.Delete(job)
	if err != nil {
		return "", err
	}

	return message, nil
}
