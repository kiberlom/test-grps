package tests

import (
	"context"
	"log"
	"net"
	"testing"

	"proto-test/internal"
	"proto-test/proto/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var service *bufconn.Listener

// инициализация grpc сервера
func init() {

	service = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	protobuf.RegisterAuthCheckerServer(s, &internal.SerWork{})

	go func() {
		if err := s.Serve(service); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return service.Dial()
}

func TestLogin(t *testing.T) {

	t.Run("GetUser", func(t *testing.T) {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
		if err != nil {
			t.Fatalf("Failed to dial bufnet: %v", err)
		}
		defer conn.Close()
		client := protobuf.NewAuthCheckerClient(conn)

		resp, err := client.GetUser(ctx, &protobuf.UUIDUser{
			UUID: "qwerty",
		})

		if err != nil {
			t.Fatalf("Login failed: %v", err)
		}

		log.Printf("Response: %+v", resp)
	})

}
