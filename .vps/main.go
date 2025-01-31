package main

import (
	// "github.com/Nelson2017-8/webdata/app"
	"fmt"
	"os"

	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/nelsonp17/webdata/app/database"

	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/nelsonp17/webdata/app/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	var (
		host = os.Getenv("HOST")
		port = os.Getenv("PORT")

		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")
	)

	pgxPrueba, err := database.NewPGXDB(dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Create a new engine
	engine := html.New("./frontend", ".html")

	//const HeaderName = "X-Csrf-Token"
	fiberApp := fiber.New(fiber.Config{
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		Views:             engine,
		PassLocalsToViews: false,
	})

	fiberApp.Static("/assets", "./frontend/assets")

	fiberApp.Use(cors.New())
	fiberApp.Use(logger.New())

	// routes
	rv2 := router.Router{
		Fwk: fiberApp,
	}

	// rutas web
	routers := router.Handler{
		Database: pgxPrueba,
	}
	routers.Api(rv2)
	routers.Web(rv2)

	// Imprimir todas las rutas registradas
	for _, route := range rv2.Fwk.GetRoutes() {
		fmt.Printf("Ruta registrada: %s %s\n", route.Method, route.Path)
	}

	err = fiberApp.Listen(host + ":" + port)
	if err != nil {
		return
	}
}
