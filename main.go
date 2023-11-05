package main

import (
	"context"
	"fmt"
	"github.com/xvbnm48/grpc-basic/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

type myInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoiceServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.ResponseRequest, error) {
	return &invoicer.ResponseRequest{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot created listener: %s", port)
	}
	fmt.Println("server is started at", port)
	serverRegistrar := grpc.NewServer()
	service := &myInvoiceServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	go func() {
		err = serverRegistrar.Serve(lis)
		if err != nil {
			log.Fatalf("Gagal melayani: %s", err)
		}
	}()
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("inposibble to serve: %s", err)
	}
}
