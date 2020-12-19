package biz

import (
	"context"
	"work/internal/data"
)

type GreeterReq struct {
	Name string
}

type GreeterRes struct {
	Age  int
	Size string
}

type Greeter struct {
	data GreeterData
}

type GreeterData interface {
	GetGreeter(name string) (age int, size string)
}

func NewGreeter(data data.Greeter) Greeter {
	return Greeter{
		data: data,
	}
}

func (g Greeter) SayHello(ctx context.Context, in GreeterReq) *GreeterRes {
	res := GreeterRes{}
	res.Age, res.Size = g.data.GetGreeter(in.Name)
	return &res
}

func (g Greeter) SayHelloAgain(ctx context.Context, in GreeterReq) *GreeterRes {
	res := GreeterRes{}
	res.Age, res.Size = g.data.GetGreeter(in.Name)
	return &res
}
