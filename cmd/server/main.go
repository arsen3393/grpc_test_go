package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"team00_01/internal/config"
	logger2 "team00_01/pkg/logger"
	"team00_01/pkg/transmitter"
)

func main() {
	stop := make(chan os.Signal, 1)
	serverChan := make(chan *grpc.Server, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	logger, _ := logger2.MustInitLogger()
	cfg := config.MustLoadConfig()
	logger.Info("Конфлиг подключен")
	go func() {
		logger.Info("gRPC сервер запущен", "port", cfg.GRPC.Port)
		s := transmitter.MustStartGrpcServer()
		serverChan <- s
		l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
		if err != nil {
			panic("Ошибка старта listener: " + err.Error())
		}
		if err := s.Serve(l); err != nil {
			panic(err)
		}
	}()

	func() {
		s := <-serverChan
		<-stop
		fmt.Println("Получен сигнал остановки")
		s.GracefulStop()
		fmt.Println("Сервер остановлен")
	}()
}
