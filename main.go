package main

import (
	"fmt"
	"github.com/gocraft/dbr/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"

	"github.com/artengin/rest-api-go/internal/http"
	"github.com/artengin/rest-api-go/internal/logic"
	"github.com/artengin/rest-api-go/internal/middleware"
	"github.com/artengin/rest-api-go/internal/repository/postgres"
)

const (
	defaultTimeout = 30
	defaultAddress = ":8080"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func runMigrations(session *dbr.Session) {
	sqlBytes, err := os.ReadFile("person.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = session.Exec(string(sqlBytes))
	if err != nil {
		log.Fatalf("Error executing migration: %v", err)
	}
}

func main() {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	conn, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		log.Fatal(err)
	}

	session := conn.NewSession(nil)

	runMigrations(session)

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("got error when closing the DB connection: %v", err)
		}
	}()

	repo := postgres.NewPersonRepository(session)
	logicLayer := logic.NewPersonLogic(repo)
	handler := http.NewPersonHandler(logicLayer)

	e := echo.New()
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Warn("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	e.GET("/person", handler.GetAllPerson)
	e.GET("/person/:id", handler.GetPerson)
	e.POST("/person", handler.CreatePerson)
	e.PUT("/person/:id", handler.UpdatePerson)
	e.DELETE("/person/:id", handler.DeletePerson)

	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	e.Logger.Fatal(e.Start(address))
}
