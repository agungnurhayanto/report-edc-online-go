package importer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Import(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file wajib diupload",
		})
		return
	}

	tempFile := "temp/" + file.Filename

	if err := c.SaveUploadedFile(file, tempFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.service.ImportFile(
		c.Request.Context(),
		tempFile,
	)
	// data, err := ParseMonitoring(tempFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		//"total_data": len(data),
		"message": "import berhasil",
	})
}
