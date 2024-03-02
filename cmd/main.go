package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gdegiorgio/grpc-bookshelf-client/internal/proto/book"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

var (
	host   = flag.String("host", "localhost", "Set the GRPC Server host")
	port   = flag.String("port", "4500", "Set the GRPC Server port")
	pretty = flag.Bool("pretty", false, "Prettify logs")
)

func main() {
	flag.Parse()
	if *pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", *host, *port), opts...)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to connect to grpc server %s:%s - %+v", *host, *port, err))
	}
	defer conn.Close()
	client := book.NewBookClient(conn)
	book, err := client.GetBook(context.Background(), &book.GetBookRequest{Id: "1"})
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to retrieve book info - %+v", err))
	}
	log.Info().Msg(fmt.Sprintf("Got a new book %+v", book))
}
