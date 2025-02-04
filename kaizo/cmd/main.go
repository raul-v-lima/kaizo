package main

import (
	"database/sql"
	"fmt"
	"kaizo/kaizo/internal/infrastrucutre/persistence"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=username dbname=kaizo sslmode=disable password=password"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}
	defer db.Close()

	monRepo := persistence.NewMonRepositoryPostgres(db)
	err = monRepo.AddRandomMons(3)
	if err != nil {
		log.Fatalf("Erro ao adicionar os mons %v", err)
	}

	mons, err := monRepo.GetAll()
	if err != nil {
		log.Fatalf("Erros ao buscar os mons %v", err)
	}
	fmt.Printf("Mons %v\n", mons)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem-vindo ao Kaizo!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor na porta %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor %v", err)
	}

}
