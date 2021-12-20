package service

import (
	"log"
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/husanmusa/to-do-service/genproto"
)

var client pb.TaskServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}

	client = pb.NewTaskServiceClient(conn)
	os.Exit(m.Run())
}
