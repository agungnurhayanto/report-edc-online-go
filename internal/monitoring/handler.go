package monitoring

import (
	"monitoring-edc/internal/importer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service       Service
	importService importer.Service
}

func NewHandler(
	service Service,
	importService importer.Service,
) *Handler {
	return &Handler{
		service:       service,
		importService: importService,
	}
}

func (h *Handler) GetAll(c *gin.Context) {

	data, err := h.service.GetAll(
		100,
		0,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
