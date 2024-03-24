package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	pb "grpctutorial/contracts"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "grpctutorial-mvrs3fl3.b4a.run:8080", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	var c pb.CalculatorClient = pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	Add(&ctx, &c, 1, 2)
	Subtract(&ctx, &c, 1, 2)
	Multiply(&ctx, &c, 2, 2)
	Divide(&ctx, &c, 1, 0)
	Divide(&ctx, &c, 4, 3)
}

func Add(ctx *context.Context, calculator *pb.CalculatorClient, a int64, b int64) {
	r, err := (*calculator).Add(*ctx, &pb.OperationRequest{A: a, B: b})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Result: %d", r.GetA())
}

func Subtract(ctx *context.Context, calculator *pb.CalculatorClient, a int64, b int64) {
	r, err := (*calculator).Subtract(*ctx, &pb.OperationRequest{A: a, B: b})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	log.Printf("Result: %d", r.GetA())
}

func Multiply(ctx *context.Context, calculator *pb.CalculatorClient, a int64, b int64) {
	r, err := (*calculator).Multiply(*ctx, &pb.OperationRequest{A: a, B: b})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}
	log.Printf("Result: %d", r.GetA())
}

func Divide(ctx *context.Context, calculator *pb.CalculatorClient, a int64, b int64) {
	r, err := (*calculator).Divide(*ctx, &pb.OperationRequest{A: a, B: b})
	if err != nil {
		log.Printf("could not divide: %v", err)
		return
	}
	log.Printf("Result: %d", r.GetA())
}
