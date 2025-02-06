package main

import (
	"fmt"
	"kaizo/kaizo/internal/infrastrucutre/persistence"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Erro ao carregar o .env: %v\n", err)
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	appPort := os.Getenv("APP_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Erro ao obter *sql.DB: %v", err)
	}
	monRepo := persistence.NewMonRepositoryPostgres(sqlDB)

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

	db.AutoMigrate(&mons)

	if appPort == "" {
		appPort = "8080"
	}
	log.Printf("Servidor na porta %s", appPort)
	if err := http.ListenAndServe(":"+appPort, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor %v", err)
	}

}
