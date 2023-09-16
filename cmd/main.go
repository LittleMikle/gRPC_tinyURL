package main

import (
	"github.com/LittleMikle/gRPC_tinyurl/configs"
	"github.com/LittleMikle/gRPC_tinyurl/pkg/handlers"
	"github.com/LittleMikle/gRPC_tinyurl/pkg/repository"
	"github.com/LittleMikle/gRPC_tinyurl/pkg/service"
	"github.com/LittleMikle/gRPC_tinyurl/proto"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const port = ":8081"

func main() {
	err := configs.InitConfig()
	if err != nil {
		log.Fatal().Msg("failed with viper config")
	} else {
		log.Info().Msg("Config initialization successful")
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("error with .env file %s", err)
	} else {
		log.Info().Msg("Config load successful")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed with Postgres connection:")
	} else {
		log.Info().Msg("Connection to Postgres successful")
	}

	s := grpc.NewServer()

	repo := repository.NewRepository(db)
	serviceURL := service.NewService(repo)
	handlersURL := handlers.NewURL(serviceURL)

	proto.RegisterURLserviseServer(s, handlersURL)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen port")
	}

	go func() {
		log.Info().Msg("LISTENING SERVER")
		err = s.Serve(lis)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to serve")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
