package http

import (
	"database/sql"
	"errors"
	"net/http"

	usecase "github.com/Aruto143/portfolio-backend/internal/usecase/work"
	"github.com/gin-gonic/gin"
)

type WorkHandler struct {
	listWorks      *usecase.ListWorksUsecase
	findWorkBySlug *usecase.FindWorkBySlugUsecase
}

func NewWorkHandler(
	list *usecase.ListWorksUsecase,
	find *usecase.FindWorkBySlugUsecase,
) *WorkHandler {
	return &WorkHandler{
		listWorks:      list,
		findWorkBySlug: find,
	}
}

func (h *WorkHandler) GetWorks(c *gin.Context) {
	works, err := h.listWorks.Run(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch works"})
		return
	}
	c.JSON(http.StatusOK, works)
}

func (h *WorkHandler) GetWorkBySlug(c *gin.Context) {
	slug := c.Param("slug")

	work, err := h.findWorkBySlug.Run(c.Request.Context(), slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "work not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch work"})
		return
	}

	c.JSON(http.StatusOK, work)
}