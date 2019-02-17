package main

import (
	"log"

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
func HandlerV2() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.GraphQL(
		v2Graph.NewExecutableSchema(
			v2Graph.Config{Resolvers: &v2Graph.Resolver{}},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
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
		v2.POST("/graphql", HandlerV2())
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Run(":3001")
}
