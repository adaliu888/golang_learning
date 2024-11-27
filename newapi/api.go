package main

import (
	"log"
	"net/http"
)

// 定义一个结构体，表示一个任务
type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

// 任务的执行函数

func (s *APIServer) Run() error {
	// 假设任务需要执行的任务
	router := http.NewServeMux()
	router.HandleFunc("/userS/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte(userID))
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	log.Panicln("Server has started", s.addr)
	return server.ListenAndServe()

}
