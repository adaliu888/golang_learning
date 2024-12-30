package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("userID" + userID))

	})

	//func(w http.ResponseWriter, r *http.Request) {
	//userID := r.PathValue("userID")
	//w.Write([]byte("userID" + userID))

	MiddlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequestAuthMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: MiddlewareChain(router),
	}
	log.Printf("Server has start %s", s.addr)

	return server.ListenAndServe()

}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, Path %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func RequestAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authoriztion")
		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}
