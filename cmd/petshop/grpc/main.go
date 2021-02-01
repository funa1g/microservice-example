package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	pb "github.com/funa1g/microservice-example/api/petshop"
	_delivery "github.com/funa1g/microservice-example/pkg/petshop/delivery/grpc"
	_repository "github.com/funa1g/microservice-example/pkg/petshop/repository/mysql"
	_usecase "github.com/funa1g/microservice-example/pkg/petshop/usecase"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Start!")

	dbConn := initDbConn()
	if dbConn == nil {
		return
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()

	s := grpc.NewServer()
	petRepo := _repository.NewPetRepository(dbConn)
	tagRepo := _repository.NewTagRepository(dbConn)
	petUseCase := _usecase.NewPetUsecase(petRepo, tagRepo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("failed to listen: %v", err))
	}
	pb.RegisterPetshopServer(s, &_delivery.PetHandler{PUsecase: petUseCase})
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msg(fmt.Sprintf("failed to serve: %v", err))
	}
}

func initDbConn() *sql.DB {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbName := "petshop"
	connection := fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}

	err = dbConn.Ping()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}

	return dbConn
}
