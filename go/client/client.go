package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"kjarmicki.github.com/grpc-node-vs-go/tasks"
)

func NewClient(address string) (tasks.TasksClient, func() error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return tasks.NewTasksClient(conn), conn.Close
}
