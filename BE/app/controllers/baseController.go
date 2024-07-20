package controllers

import (
	"fmt"
	"log"
	"login-web/app/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB   *gorm.DB
	App  *fiber.App
	User models.User
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(dbConfig DBConfig) {
	fmt.Println("we are initializing the server")

	server.App = fiber.New()

	// Tambahkan middleware CORS
	server.App.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	server.InitializeDB(dbConfig)
	server.Setup(server.App)
}

func (server *Server) InitializeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected successfully")
}

func (server *Server) dbMigrate() {
	var err error
	for _, model := range models.RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database migrated successfully")
}

func (server *Server) InitCommands(dbConfig DBConfig) {
	server.InitializeDB(dbConfig)

	cmdApp := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "db:migrate",
				Action: func(c *cli.Context) error {
					server.dbMigrate()
					return nil
				},
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	server.App.Listen(addr)
}
