package service

import (
	"context"
	"fmt"
	v1 "work/api/greeter/v1"
	"work/internal/biz"
)

type Greeter struct {
	v1.UnimplementedGreeterServer

	biz GreeterBiz
}

type GreeterBiz interface {
	SayHello(ctx context.Context, in biz.GreeterReq) *biz.GreeterRes
	SayHelloAgain(ctx context.Context, in biz.GreeterReq) *biz.GreeterRes
}

func NewGreeter(biz biz.Greeter) Greeter {
	return Greeter{
		biz: biz,
	}
}

func (s *Greeter) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {

	req := biz.GreeterReq{
		Name: in.GetName(),
	}

	res := s.biz.SayHello(ctx, req)

	return &v1.HelloReply{Message: fmt.Sprintf("Hello! name=%s, age=%d, size=%s", in.Name, res.Age, res.Size)}, nil
}

func (s *Greeter) SayHelloAgain(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {

	req := biz.GreeterReq{
		Name: in.GetName(),
	}

	res := s.biz.SayHelloAgain(ctx, req)

	return &v1.HelloReply{Message: fmt.Sprintf("Hello again! name=%s, age=%d, size=%s", in.Name, res.Age, res.Size)}, nil
}
