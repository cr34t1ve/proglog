package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	api "github.com/cr34t1ve/proglog/api/v1"
	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("addr", ":8400", "service address")
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewLogClient(conn)
	ctx := context.Background()
	res, err := client.GetServers(ctx, &api.GetServersRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("servers:")
	for _, server := range res.Servers {
		fmt.Printf("\t- %v\n", server)
	}

}
