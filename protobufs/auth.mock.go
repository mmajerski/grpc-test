// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth.proto

package grpc

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AuthMock struct{}

func (m *AuthMock) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	res :=
		&LoginReply{
			Token: "fuga",
		}
	return res, nil
}
