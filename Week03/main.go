package main

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {

	ctx := context.Background()

	srv := http.Server{
		Addr:    ":8080",
		Handler: &appHander{ctx: ctx},
	}

	idleConnsClosed := make(chan struct{})
	// 注销
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt,syscall.SIGTERM, syscall.SIGUSR1)
		<-sigint

		if err := srv.Shutdown(ctx); err != nil {

			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

type appHander struct {
	ctx context.Context
}

func (h *appHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	Info := func(ctx context.Context) (*sync.Map, error) {
		g, ctx := errgroup.WithContext(ctx)
		var info sync.Map
		// 获取cpu信息
		g.Go(func() error {
			result, err := cpu.Info()

			if err == nil {
				res := "Cpu info"
				for i := 0; i < len(result); i++ {
					res = strings.Join([]string{res, result[i].String()}, ",")
				}
				fmt.Println(res)
				info.Store("cpu", res)
			}
			return err
		})
		// 获取内存信息
		g.Go(func() error {
			result, err := mem.VirtualMemory()
			if err == nil {
				res := "Men info"
				res = strings.Join([]string{res, result.String()}, ":")
				info.Store("mem", res)
			}
			return err
		})

		if err := g.Wait(); err != nil {
			return &info, err
		}
		return &info, nil
	}
	result, err := Info(h.ctx)
	if err != nil {
		log.Fatal(err)
	}

	cpuinfo, _ := result.Load("cpu")

	if _, ok := vars["cpu"]; ok {

		io.WriteString(w, cpuinfo.(string))
	}
	meminfo, _ := result.Load("men")
	if _, ok := vars["mem"]; ok {
		io.WriteString(w, meminfo.(string))
	}

}
