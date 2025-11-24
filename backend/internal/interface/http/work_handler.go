package http

import (
	"log"
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

func (h *WorkHandler) GetWorks(c *gin.Context) {
    works, err := h.listWorks.Run(c.Request.Context())
    if err != nil {
        // ★ ここ追加：エラー内容をログに出す
        log.Printf("GetWorks error: %v", err)

        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch works"})
        return
    }
    c.JSON(http.StatusOK, works)
}
