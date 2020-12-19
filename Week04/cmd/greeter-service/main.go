package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	greeter "work/api/greeter/v1"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	s := grpc.NewServer()

	greeterService := InitializeGreeterService()
	greeter.RegisterGreeterServer(s, &greeterService)

	// grpc server
	g.Go(func() error {
		fmt.Println("grpc")
		go func() {
			<-ctx.Done()
			fmt.Println("grpc ctx done")
			s.GracefulStop()
		}()
		return s.Serve(lis)
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				cancel()
				return nil
			}
		}
	})

	err = g.Wait() // first error return
	fmt.Println(err)
}
