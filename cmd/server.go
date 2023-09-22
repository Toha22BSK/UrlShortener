package cmd

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Toha22BSK/UrlShortener/gen/database"
	"github.com/Toha22BSK/UrlShortener/gen/restapi"
	"github.com/Toha22BSK/UrlShortener/gen/restapi/operations"
	"github.com/Toha22BSK/UrlShortener/internal/urls"
	"github.com/go-openapi/loads"

	// "github.com/go-swagger/go-swagger/examples/oauth2/restapi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func RunServer() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Panic("failed to load .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		logrus.WithError(err).Panic("PORT is not found in environment variables")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		logrus.WithError(err).Panic("Can't connect to database: ", err)
	}

	host, err := os.Hostname()
	if err != nil {
		logrus.WithError(err).Fatal("failed to get hostname")
	}

	logrus.SetLevel(logrus.InfoLevel)
	logrus.WithFields(logrus.Fields{
		"Host": host,
	}).Info("Service Startup")

	swaggerSpec, err := loads.Analyzed(restapi.FlatSwaggerJSON, "")
	if err != nil {
		panic(err)
	}

	api := operations.NewBackendCoreAPI(swaggerSpec)
	api.Logger = logrus.Printf

	db := database.New(conn)

	urlService := urls.NewApiConfig(db)
	urls.Configure(api, urlService)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080
	server.EnabledListeners = []string{"http"}

	c := cors.AllowAll()
	mux := http.NewServeMux()
	mux.Handle("/", api.Serve(nil))

	server.SetHandler(c.Handler(mux))

	if err := server.Serve(); err != nil {
		logrus.Panicln(err)
	}

}
