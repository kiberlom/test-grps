package internal

import (
	"context"
	"fmt"
	"proto-test/proto/protobuf"
)

type SerWork struct {
	protobuf.UnimplementedAuthCheckerServer
}

func (s *SerWork) GetUser(ctx context.Context, uuid *protobuf.UUIDUser) (*protobuf.User, error) {

	if uuid.UUID == "qwerty" {
		return &protobuf.User{
			UUID: uuid.UUID,
			Name: "denis",
		}, nil
	}

	return nil, fmt.Errorf("Не верный UUID")
}
