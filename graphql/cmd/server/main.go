package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rodrwan/opendata/graphql/earthquake"
	"github.com/rodrwan/opendata/graphql/gmarcone"
	"github.com/rodrwan/opendata/graphql/transapi"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	v1Graph "github.com/rodrwan/opendata/graphql/v1"
	v2Graph "github.com/rodrwan/opendata/graphql/v2"
	cors "github.com/rs/cors/wrapper/gin"
)

// HandlerV1 initializes the graphql middleware.
func HandlerV1(schema graphql.Schema) gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := v1Graph.GraphqlHandler{
		Schema: schema,
	}

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// HandlerV2 initializes the graphql middleware.
func HandlerV2(URL string) gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema

	h := handler.GraphQL(
		v2Graph.NewExecutableSchema(
			v2Graph.Config{Resolvers: &v2Graph.Resolver{
				GMarconeClient: gmarcone.NewClient(URL),
				Transapi:       transapi.NewClient(),
				Earthquake:     earthquake.NewClient(),
			}},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	gmarconeURL := flag.String("gmarcone-url", "http://gmarcone:3002", "gmarcone service url")

	flag.Parse()

	router := gin.Default()
	router.Use(cors.Default())

	schema, err := graphql.NewSchema(v1Graph.SchemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/graphql", HandlerV1(schema))
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/graphql", HandlerV2(*gmarconeURL))
	}

	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Run(":3001")
}
