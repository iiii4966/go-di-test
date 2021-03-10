package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-di-fiber/domain/model"
	"go-di-fiber/repository"
	"go-di-fiber/service"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var isPrefork = flag.Bool("prefork", false, "whether server multiprocessor")
var host = flag.String("host", "0.0.0.0", "server host")
var port = flag.String("port", "8080", "server port")
var dbUserName = flag.String("dbUserName", "tester", "db user")
var dbPassword = flag.String("dbPassword", "1234", "db password")
var dbHost = flag.String("dbHost", "localhost", "db host")
var dbPort = flag.String("dbPort", "5432", "db port")
var dbName = flag.String("dbName", "test", "db name")

func NewServer() *fiber.App {
	serverConfig := fiber.Config{Prefork: *isPrefork}
	return fiber.New(serverConfig)
}

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		*dbHost, *dbUserName, *dbPassword, *dbName, *dbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func SetupMiddleware(server *fiber.App) {
	server.Use(recover.New())
	server.Use(cors.New(corsConfig))
	server.Use(logger.New(loggerConfig))
	server.Use(pprof.New())
}

func SetupRoutes(server *fiber.App, userService *service.UserService, validator *validator.Validate) {
	server.Get("/healthcheck", heathCheck)
	userRoutes := server.Group("/user")
	userRoutes.Post("/register", signup(userService, validator))
}

func migration(db *gorm.DB) {
	model.UserMigrate(db)
}

func Register(
	lc fx.Lifecycle,
	server *fiber.App,
	db *gorm.DB,
	userService *service.UserService,
	validator *validator.Validate,
) *fiber.App {
	// Don't change order.
	SetupMiddleware(server)
	SetupRoutes(server, userService, validator)

	migration(db)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.Listen(*host + ":" + *port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})

	return server
}

func main() {
	flag.Parse()
	fx.New(
		fx.Options(
			fx.Provide(
				NewValidator,
				NewServer,
				NewDB,
			),
			repository.UserRepoModule,
			service.UserServiceModule,
		),
		fx.Invoke(
			Register,
		),
	).Run()
}
