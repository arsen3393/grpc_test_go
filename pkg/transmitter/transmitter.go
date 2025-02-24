package transmitter

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"math/rand"
	"team00_01/pkg/api"
	logger2 "team00_01/pkg/logger"
	"time"
)

type GrpcServer struct {
	api.UnimplementedTransmitterServer
}

func (s *GrpcServer) StreamData(req *api.RequestMessage, stream grpc.ServerStreamingServer[api.TransmitData]) error {
	uuid := uuid.New().String()
	_ = uuid
	_, fileLogger := logger2.MustInitLogger()
	for {
		mean := rand.Float64()*20 - 10
		standartDeviation := rand.Float64()*1.2 + 0.3

		request := &api.TransmitData{
			Frequency: rand.NormFloat64()*standartDeviation - mean,
			SessionId: uuid,
			Timestamp: time.Now().UTC().Unix(),
		}
		if err := stream.Send(request); err != nil {
			return err
		}
		fileLogger.Info(fmt.Sprintf("SessionID: %s, Mean: %f, standartDeviation: %f", request.SessionId, mean, standartDeviation))
		time.Sleep(time.Second*1 - time.Millisecond*250)
	}

	return nil
}

func MustStartGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	srv := &GrpcServer{}
	api.RegisterTransmitterServer(s, srv)

	return s
}
