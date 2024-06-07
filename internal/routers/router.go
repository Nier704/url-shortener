package routes

import (
	"log"
	"net/http"

	"github.com/Nier704/url-shortener/internal/handlers"
	"github.com/Nier704/url-shortener/internal/middlewares"
)

type Router struct {
	Handler *handlers.UrlHandler
	mux     *http.ServeMux
	Port    string
}

func (r *Router) Init() {
	r.mux.Handle("POST /shortener/newUrl", middlewares.Log(http.HandlerFunc(r.Handler.CreateUrl)))
	r.mux.Handle("GET /shortener", middlewares.Log(http.HandlerFunc(r.Handler.GetUrl)))
}

func (r *Router) Start() {
	println("Server is running at port: " + r.Port)
	if err := http.ListenAndServe(":"+r.Port, r.mux); err != nil {
		log.Printf("Error starting the server: %v", err)
	}
}

func NewRouter(h *handlers.UrlHandler) *Router {
	return &Router{
		mux:     http.NewServeMux(),
		Handler: h,
		Port:    "8080",
	}
}
