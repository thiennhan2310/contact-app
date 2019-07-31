package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	c "app/contact"

	"database/sql"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	host     = "satao.db.elephantsql.com"
	port     = 5432
	user     = "erkxmoag"
	password = "fts5UmgTIbP1CUqk1ELAK-tdjK7H6jmp"
	dbname   = "erkxmoag"
)

var db *sql.DB

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) AddContact(ctx context.Context, req *c.AddContactReq) (*c.AddContactRep, error) {
	if req.Name == "" || req.Phone == "" {
		return nil, errors.New("Missing field")
	}

	query := fmt.Sprintf("INSERT INTO contact(name, phone) VALUES('%s', '%s')", req.Name, req.Phone)
	_, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return &c.AddContactRep{Message: "Success"}, nil
}

func (s *server) GetContact(ctx context.Context, in *c.GetContactReq) (*c.GetContactRep, error) {
	query := fmt.Sprintf("Select id,name,phone from contact where id = %v", 1)
	rep := new(c.GetContactRep)
	row := db.QueryRow(query)
	row.Scan(&rep.Id, &rep.Name, &rep.Phone)
	return rep, nil
}

func main() {
	var err error
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := grpc.NewServer()
	c.RegisterContactServer(s, &server{})
	log.Println("Started grpc server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
