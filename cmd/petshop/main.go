package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	_delivery "github.com/funa1g/microservice-example/pkg/petshop/delivery/http"
	_repository "github.com/funa1g/microservice-example/pkg/petshop/repository/mysql"
	_usecase "github.com/funa1g/microservice-example/pkg/petshop/usecase"
)

func main() {
	fmt.Println("Hello World!")

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
		fmt.Println(err.Error())
		return
	}

	err = dbConn.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	petRepo := _repository.NewPetRepository(dbConn)
	tagRepo := _repository.NewTagRepository(dbConn)
	petUseCase := _usecase.NewPetUsecase(petRepo, tagRepo)
	_delivery.NewPetHandler(e, petUseCase)
	e.Logger.Fatal(e.Start(":8080"))
}
