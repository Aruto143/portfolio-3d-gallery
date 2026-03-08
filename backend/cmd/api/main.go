package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	infra "github.com/Aruto143/portfolio-backend/internal/infrastructure/mysql"
	ifhttp "github.com/Aruto143/portfolio-backend/internal/interface/http"
	usecase "github.com/Aruto143/portfolio-backend/internal/usecase/work"
)

func main() {
	// なぜ環境変数か？
	// → 本番（Railway）とローカルで接続先を切り替えるため
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// ローカル開発用のデフォルト値
		dsn = "appuser:apppass@tcp(localhost:3306)/portfolio?parseTime=true&charset=utf8mb4"
	}

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	// DI（依存性注入）
	repo := infra.NewWorkRepository(db)
	listUC := usecase.NewListWorksUsecase(repo)
	handler := ifhttp.NewWorkHandler(listUC)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/works", handler.GetWorks)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
