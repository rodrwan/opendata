package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

// GraphRequest ...
type GraphRequest struct {
	Query string `json:"query"`
}

// RootQuery ...
var RootQuery = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	},
}

// SchemaConfig ...
var SchemaConfig = graphql.SchemaConfig{
	Query: graphql.NewObject(RootQuery),
}

// GraphqlHandler ...
type GraphqlHandler struct {
	Schema graphql.Schema
}

func (gh *GraphqlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var query GraphRequest

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
	if err := json.Unmarshal(bb, &query); err != nil {
		http.Error(w, "could not unmarshal data", http.StatusInternalServerError)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        gh.Schema,
		RequestString: query.Query,
	})
	if len(result.Errors) > 0 {
		http.Error(w, fmt.Sprintf("could not marshal data: %v", err), http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not marshal data: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
