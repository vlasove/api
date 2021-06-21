package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/vlasove/api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)

	go func() {
		if err := s.Start(); err != nil {
			s.Logger().Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	s.Logger().Info("Shutting down API Server....")
	if err := s.Shutdown(context.Background()); err != nil {
		s.Logger().Errorf("error occured on server shutting down: %s", err.Error())
	}
	fmt.Println("I can do more stuff...")
	time.Sleep(time.Second * 2)
	fmt.Println("Bye")
}