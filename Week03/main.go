package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func myHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 5) // 测试 Shutdown
		fmt.Fprintln(w, "Hello, world!")
	})
}

func main() {

	log.SetFlags(log.Lshortfile)

	s := &http.Server{
		// Addr: ":3306", // 测试启动错误
		Addr:    ":8080",
		Handler: myHandler(),
	}

	sigs := make(chan os.Signal, 0)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	g := new(errgroup.Group)

	g.Go(func() error {
		err := s.ListenAndServe()
		close(sigs)
		return err
	})

	g.Go(func() error {
		<-sigs
		// log.Println("sig:", sig) // 打印退出信号
		log.Println("http: Server closing")
		s.Shutdown(context.Background())
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}
