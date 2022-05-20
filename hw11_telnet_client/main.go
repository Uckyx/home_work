package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var timeout time.Duration

const (
	minArgsCount   = 2
	defaultTimeout = 10
)

func init() {
	flag.DurationVar(&timeout, "timeout", defaultTimeout*time.Second, "connection timeout")
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != minArgsCount {
		log.Fatal("host or port are missed")
	}

	cli := NewTelnetClient(net.JoinHostPort(args[0], args[1]), timeout, os.Stdin, os.Stdout)
	if err := cli.Connect(); err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		if err := cli.Send(); err != nil {
			log.Printf("error during send: %v", err)
		} else {
			log.Printf("EOF")
		}
		cancel()
	}()

	go func() {
		if err := cli.Receive(); err != nil {
			log.Printf("error during receive: %v", err)
		} else {
			log.Printf("connection was closed by peer")
		}
		cancel()
	}()

	<-ctx.Done()
}
