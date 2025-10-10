package main

import (
    "context"
    "log"
    "net"

    pb "github.com/example/analytics/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedAnalyzerServer
}

func (s *server) GetUserScore(ctx context.Context, req *pb.UserRequest) (*pb.UserScore, error) {
    score := float64(req.UserId) * 1.42
    log.Printf("Request user_id=%d, score=%.2f", req.UserId, score)
    return &pb.UserScore{Score: score}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterAnalyzerServer(grpcServer, &server{})

    log.Println("gRPC server listening on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
