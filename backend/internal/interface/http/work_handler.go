package http

import (
	"net/http"

	usecase "github.com/Aruto143/portfolio-backend/internal/usecase/work"
	"github.com/gin-gonic/gin"
)

type WorkHandler struct {
	listWorks *usecase.ListWorksUsecase
}

func NewWorkHandler(list *usecase.ListWorksUsecase) *WorkHandler {
	return &WorkHandler{listWorks: list}
}

// GetWorks は GET /api/works を処理するHTTPハンドラー。
func (h *WorkHandler) GetWorks(c *gin.Context) {
	works, err := h.listWorks.Run(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch works"})
		return
	}
	c.JSON(http.StatusOK, works)
}
