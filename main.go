package main

import (
	"github.com/ABHIJEET-MUNESHWAR/Todo/internal/db"
	"github.com/ABHIJEET-MUNESHWAR/Todo/internal/todo"
	"github.com/ABHIJEET-MUNESHWAR/Todo/internal/transport"
	"log"
)

func main() {
	d, err := db.New("postgres", "Timpoo@1", "postgres", "localhost", 5432)
	if err != nil {
		log.Fatal(err)
	}
	svc := todo.NewService(d)
	server := transport.NewServer(svc)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}