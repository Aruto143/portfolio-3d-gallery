package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	infraDB "github.com/Aruto143/portfolio-backend/internal/infra/db"
	infraRepo "github.com/Aruto143/portfolio-backend/internal/infrastructure/postgres"
	httpif "github.com/Aruto143/portfolio-backend/internal/interface/http"
	usecase "github.com/Aruto143/portfolio-backend/internal/usecase/work"
)

func main() {
	db := infraDB.NewPostgresDB()
	defer db.Close()

	workRepo := infraRepo.NewWorkRepository(db)

	listWorksUC := usecase.NewListWorksUsecase(workRepo)
	findWorkBySlugUC := usecase.NewFindWorkBySlugUsecase(workRepo)

	workHandler := httpif.NewWorkHandler(listWorksUC, findWorkBySlugUC)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		api.GET("/works", workHandler.GetWorks)
		api.GET("/works/:slug", workHandler.GetWorkBySlug)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "18080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
