package main

import (
	"fmt"
	"github.com/NazarBiloys/nosql-database-redis/internal/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func handlerRedisRDB(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		queueName := r.URL.Path[len("/redis-rdb/"):]
		err := service.PutQueueMessageToRedis(queueName, "RDB")
		if err != nil {
			log.Error(err)
		}
	}

	if r.Method == "GET" {
		queueName := r.URL.Query().Get("queueName")
		message, err := service.ReadQueueMessageFromRedis(queueName, "RDB")
		if err != nil {
			log.Error(err)
		}
		_, err = fmt.Fprintf(w, "message : %s", message)
		if err != nil {
			log.Error(err)
		}
	}
}

func handlerRedisAOF(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		queueName := r.URL.Path[len("/redis-aof/"):]
		err := service.PutQueueMessageToRedis(queueName, "AOF")
		if err != nil {
			log.Error(err)
		}
	}

	if r.Method == "GET" {
		queueName := r.URL.Query().Get("queueName")
		message, err := service.ReadQueueMessageFromRedis(queueName, "AOF")
		if err != nil {
			log.Error(err)
		}
		_, err = fmt.Fprintf(w, "message : %s", message)
		if err != nil {
			log.Error(err)
		}
	}
}

func handlerBeanstalkd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		queueName := r.URL.Path[len("/beanstalkd/"):]
		err := service.PutQueueMessageToBeanstalkd(queueName)
		if err != nil {
			log.Error(err)
		}
	}

	if r.Method == "GET" {
		queueName := r.URL.Query().Get("queueName")
		message, err := service.ReadQueueMessageFromBeanstalkd(queueName)
		if err != nil {
			log.Error(err)
		}
		_, err = fmt.Fprintf(w, "message : %s", message)
		if err != nil {
			log.Error(err)
		}
	}
}

func main() {
	http.HandleFunc("/redis-rdb/", handlerRedisRDB)
	http.HandleFunc("/redis-aof/", handlerRedisAOF)
	http.HandleFunc("/beanstalkd/", handlerBeanstalkd)
	log.Fatal(http.ListenAndServe(":90", nil))
}
