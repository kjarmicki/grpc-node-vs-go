package main

import (
	"context"
	"flag"
	"io"
	"log"

	"kjarmicki.github.com/grpc-node-vs-go/client"
	server "kjarmicki.github.com/grpc-node-vs-go/server"
	"kjarmicki.github.com/grpc-node-vs-go/tasks"
)

const NODE_SERVER_ADDRESS = "localhost:50050"
const GO_SERVER_ADDRESS = "localhost:50051"

func startServer() {
	server.StartNewTasksServer(GO_SERVER_ADDRESS)
}

func startClient() {
	cl, close := client.NewClient(NODE_SERVER_ADDRESS)
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

func main() {
	mode := flag.String("mode", "server", "")
	flag.Parse()

	switch *mode {
	case "server":
		startServer()
	case "client":
		startClient()
	default:
		log.Fatalf("unrecognized run mode: %s", *mode)
	}
}
