package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"sync"
	"team00_01/internal/config"
	"team00_01/pkg/api"
	logger2 "team00_01/pkg/logger"
	"team00_01/pkg/model"
	"team00_01/pkg/stat"
	"team00_01/storage"
	"time"
)

var (
	k float64
)

func main() {

	flag.Float64Var(&k, "k", 0.5, "anomaly coefficient")
	flag.Parse()
	logger, _ := logger2.MustInitLogger()
	cfg := config.MustLoadConfig()
	db := storage.MustConnectDB(&cfg.DB)

	address := fmt.Sprintf("%s:%d", cfg.GRPC.Address, cfg.GRPC.Port)
	logger.Info("gRPC address:" + address)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("Failed to create gRPC client: ", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			logger.Error("Error to close connect server: %s", err)
			os.Exit(1)
		}
	}()
	client := api.NewTransmitterClient(conn)
	ctx := context.Background()
	req := api.RequestMessage{
		ClientId: "1",
	}
	stream, err := client.StreamData(ctx, &req)
	if err != nil {
		log.Panic("Error to send data: ", err)
	}

	buffer := sync.Pool{
		New: func() interface{} {
			return new(api.TransmitData)
		},
	}
	statistic := &stat.Stat{
		Count: 0,
	}
	var msg *api.TransmitData
	for {
		msg = buffer.Get().(*api.TransmitData)
		msg.Reset()
		msg, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				logger.Info("Stream ended")
				break
			} else {
				logger.Error("Error to receive data: %s", err)
				os.Exit(1)
			}
		}
		statistic.InsertNewValue(msg.Frequency)
		logger.Info(fmt.Sprintf("count : %d Mean: %f, StandartDevation: %f", statistic.Count, statistic.Mean, statistic.StdDev))
		if statistic.Count > 10 {
			if statistic.CheckAnomaly(msg.GetFrequency(), k) {
				logger.Info(fmt.Sprintf("Anomaly detected: count: %d , mean: %f, stdandartDev: %f ", int(statistic.Count), statistic.Mean, statistic.StdDev))
				db.Create(&model.Anomalies{
					SessionID: msg.SessionId,
					Frequency: msg.Frequency,
					Timestamp: time.Unix(msg.Timestamp, 0).UTC(),
				})
			}
		}
		buffer.Put(msg)
	}
}
