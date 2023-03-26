package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	tasks_grpc "kjarmicki.github.com/grpc-node-vs-go/tasks"
)

func NewTasksServer() *TasksServer {
	return &TasksServer{
		list: make([]*tasks_grpc.Task, 0),
	}
}

func StartNewTasksServer(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	tasks_grpc.RegisterTasksServer(grpcServer, NewTasksServer())
	fmt.Printf("bringing up server at %s\n", address)
	grpcServer.Serve(lis)
}

type TasksServer struct {
	list []*tasks_grpc.Task
	tasks_grpc.UnimplementedTasksServer
}

func (ts *TasksServer) AddTask(ctx context.Context, task *tasks_grpc.Task) (*tasks_grpc.Empty, error) {
	ts.list = append(ts.list, task)
	fmt.Println("added 1 task")
	return &tasks_grpc.Empty{}, nil
}

func (ts *TasksServer) ListTasks(empty *tasks_grpc.Empty, listTasks tasks_grpc.Tasks_ListTasksServer) error {
	for _, task := range ts.list {
		err := listTasks.Send(task)
		if err != nil {
			return err
		}
	}
	fmt.Printf("listed %d tasks\n", len(ts.list))
	return nil
}
