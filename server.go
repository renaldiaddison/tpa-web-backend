package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/renaldiaddison/tpa-web-backend/database"
	"github.com/renaldiaddison/tpa-web-backend/directives"
	"github.com/renaldiaddison/tpa-web-backend/graph"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/middlewares"
)

const defaultPort = "8080"

func MyCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.GetDatabase()
	database.MigrateDatabase()

	c := generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router := mux.NewRouter()
	router.Use(MyCors)
	router.Use(middlewares.AuthMiddleware)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
