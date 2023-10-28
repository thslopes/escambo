package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/thslopes/escambo/db"
)

func Server() {

	dbClient = &db.MongoDBClient{}
	err := dbClient.Connect()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbClient.Disconnect()
		if err != nil {
			panic(err)
		}
	}()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// RESTy routes for "articles" resource
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", ListArticles)
		r.Post("/", CreateArticle) // POST /articles

		r.Route("/{articleID}", func(r chi.Router) {
			r.Get("/", GetArticle)       // GET /articles/123
			r.Put("/", UpdateArticle)    // PUT /articles/123
			r.Delete("/", DeleteArticle) // DELETE /articles/123
		})
	})

	fmt.Println("Server running on http://localhost:3333")
	err = http.ListenAndServe(":3333", r)
	if err != nil {
		panic(err)
	}
}
