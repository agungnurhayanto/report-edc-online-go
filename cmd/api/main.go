package main

import (
	"monitoring-edc/internal/database"
	"monitoring-edc/internal/importer"
	"monitoring-edc/internal/monitoring"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	repo := monitoring.NewRepository(database.DB)

	service := monitoring.NewService(repo)
	handler := monitoring.NewHandler(service)
	importService := importer.NewService(service)
	importHandler := importer.NewHandler(importService)

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/monitoring", handler.GetAll)
	r.POST("/monitoring/import", importHandler.Import)
	r.GET("/dashboard", handler.GetDashboard)
	r.GET("/dashboard/edp", handler.GetEDP)

	if err := r.Run(":8989"); err != nil {
		panic(err)
	}
}
