package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

var corsConfig = cors.Config{
	Next:         nil,
	AllowOrigins: "*",
	AllowMethods: strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodDelete,
		fiber.MethodPatch,
		fiber.MethodOptions,
	}, ","),
	AllowHeaders:     "*",
	AllowCredentials: false,
	ExposeHeaders:    "",
	MaxAge:           60,
}

var loggerConfig = logger.Config{
	Next:         nil,
	Format:       "${time} | ${status} | ${latency} | ${method} | ${path} | ${ua} \n",
	TimeFormat:   "2006-01-02T15:04:05",
	TimeZone:     "Local",
	TimeInterval: 500 * time.Millisecond,
	Output:       os.Stderr,
}

var gormConfig = gorm.Config{
	PrepareStmt: true,
}
