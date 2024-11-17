package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var ctx context.Context = context.Background()

func main() {
	Conn()
}

func Conn() {
	// psql-U postgres -h localhost -d postgres
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@0.0.0.0:5432/postgres")
	if err != nil {
		log.Println("error on connecting to database", err)
		return
	}
	fmt.Println(conn)
}
