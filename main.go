package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
	"github.com/vinceferro/summarist-graphql/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var categoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book, ok := params.Source.(*models.Book)
					if !ok {
						fmt.Println("Source was not a book")
						return nil, nil
					}
					return book.Author.String, nil
				},
			},
			"isbn": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book, ok := params.Source.(*models.Book)
					if !ok {
						fmt.Println("Source was not a book")
						return nil, nil
					}
					return book.Isbn.Int64, nil
				},
			},
			"published_at": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book, ok := params.Source.(*models.Book)
					if !ok {
						fmt.Println("Source was not a book")
						return nil, nil
					}
					return book.PublishedAt.Time, nil
				},
			},
			"publisher": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book, ok := params.Source.(*models.Book)
					if !ok {
						fmt.Println("Source was not a book")
						return nil, nil
					}
					return book.Publisher.String, nil
				},
			},
			"cover_url": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book, ok := params.Source.(*models.Book)
					if !ok {
						fmt.Println("Source was not a book")
						return nil, nil
					}
					return book.CoverURL.String, nil
				},
			},
			"overview": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book, ok := params.Source.(*models.Book)
					if !ok {
						fmt.Println("Source was not a book")
						return nil, nil
					}
					return book.Overview.String, nil
				},
			},
			"key_insights": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
			"category": &graphql.Field{
				Type: categoryType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book, ok := params.Source.(*models.Book)
					if !ok {
						fmt.Println("Source was not a book")
						return nil, nil
					}
					return book.R.Category, nil
				},
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"books": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get book list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					books, err := models.Books(
						qm.Load(models.BookRels.Category),
					).AllG()
					if err != nil {
						fmt.Println(err)
					}
					return books, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	db, err := sql.Open("postgres", "dbname=summarist user=user password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	boil.SetDB(db)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
