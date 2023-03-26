package main

import (
	"context"
	"io"
	"log"
	"time"

	"kjarmicki.github.com/grpc-node-vs-go/client"
	server "kjarmicki.github.com/grpc-node-vs-go/server"
	"kjarmicki.github.com/grpc-node-vs-go/tasks"
)

const SERVER_ADDRESS = "localhost:50051"

func main() {
	go func() {
		server.StartNewTasksServer(SERVER_ADDRESS)
	}()
	<-time.After(time.Second)
	cl, close := client.NewClient(SERVER_ADDRESS)
	defer close()

	_, err := cl.AddTask(context.Background(), &tasks.Task{
		Title:       "first task",
		Description: "hello world",
		Status:      tasks.TaskStatus_TODO,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	list, err := cl.ListTasks(context.Background(), &tasks.Empty{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	for {
		task, err := list.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Println(task)
	}
}
