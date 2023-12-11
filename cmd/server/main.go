package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/felipefbs/grpc/internal/databases"
	"github.com/felipefbs/grpc/internal/pb"
	"github.com/felipefbs/grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS categories (id string, name string, description string); CREATE UNIQUE INDEX category_name_index ON categories (name)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS  courses (id string, name string, description string, category_id); CREATE UNIQUE INDEX course_name_index ON courses (name)")
	if err != nil {
		log.Fatal(err)
	}

	categoryRepo := databases.NewCategory(db)
	_, err = categoryRepo.Create("Tecnologia", "Cursos de Tecnologia")
	if err != nil {
		log.Fatal(err)
	}

	categorySVC := service.NewCategoryService(categoryRepo)

	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categorySVC)
	reflection.Register(server)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
