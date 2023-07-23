//Implementation of the calculator client
//By default listens to the port ":8081"
//WARNING: runs with GRPC method WithInsecure, don't use while on production

package main

import (
	"context"
	"flag"
	"github.com/olhavasylievadev/calc-grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	//Checking command line arguments. If there are less than 3, user missed
	//some arguments. If there are more than 3, user added some extra input
	//which is not supported
	flag.Parse()
	if flag.NArg() < 3 {
		log.Fatal("not enough arguments: number 'operator' number")
	} else if flag.NArg() > 3 {
		log.Fatal("please, don't use any unrelated symbols: number 'operator' number")
	}

	first, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal("cannot read the left side of expression")
	}

	operation := flag.Arg(1)
	second, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		log.Fatal("cannot read the right side of expression")
	}

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot connect to the server")
	}

	c := api.NewCalculatingServiceClient(conn)

	res, err := c.Calculate(context.Background(), &api.CalculateRequest{
		First:     float32(first),
		Operation: operation,
		Second:    float32(second),
	})
	if err != nil {
		log.Fatal("cannot perform the operation")
	}

	log.Println(res.GetRes())
}
